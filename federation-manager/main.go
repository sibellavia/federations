package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Federation struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ServiceCatalogue []Service `json:"service_catalogue"`
}

type Service struct {
	ID           int    `json:"id"`
	FederationID int    `json:"federation_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./federations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables if not exists
	createFederationsTable := `
    CREATE TABLE IF NOT EXISTS federations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );
    `
	_, err = db.Exec(createFederationsTable)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createFederationsTable)
		return
	}

	createServicesTable := `
    CREATE TABLE IF NOT EXISTS services (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        federation_id INTEGER NOT NULL,
        name TEXT NOT NULL,
        description TEXT,
        FOREIGN KEY(federation_id) REFERENCES federations(id)
    );
    `
	_, err = db.Exec(createServicesTable)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createServicesTable)
		return
	}

	router := mux.NewRouter() // Create a mux router

	// Federation management endpoints:
	router.HandleFunc("/federations", handleFederations).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/federations/{fed_id}", handleFederationByID).Methods(http.MethodGet, http.MethodDelete)

	log.Println("Federation Manager Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router)) // Pass the router to ListenAndServe
}

// handleFederations handles requests to /federations
func handleFederations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost: // Create a new federation
		var federation Federation
		err := json.NewDecoder(r.Body).Decode(&federation)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO federations (name) VALUES (?)", federation.Name)
		if err != nil {
			http.Error(w, "Failed to create federation", http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
			return
		}

		federation.ID = int(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(federation)

	case http.MethodGet: // Get all federations
		rows, err := db.Query("SELECT id, name FROM federations")
		if err != nil {
			http.Error(w, "Failed to retrieve federations", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var federations []Federation
		for rows.Next() {
			var federation Federation
			err := rows.Scan(&federation.ID, &federation.Name)
			if err != nil {
				http.Error(w, "Failed to scan federation", http.StatusInternalServerError)
				return
			}
			federations = append(federations, federation)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(federations)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleFederationByID handles requests to /federations/{fed_id}
func handleFederationByID(w http.ResponseWriter, r *http.Request) {
	// 1. Get the federation ID from the URL parameters
	vars := mux.Vars(r)
	fedID, err := strconv.Atoi(vars["fed_id"])
	if err != nil {
		http.Error(w, "Invalid federation ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet: // Get a specific federation
		// 1. Query the database for the federation
		var federation Federation
		err = db.QueryRow("SELECT id, name FROM federations WHERE id = ?", fedID).Scan(&federation.ID, &federation.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Federation not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to get federation", http.StatusInternalServerError)
			}
			return
		}

		// 2. Send the response (federation details)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(federation)

	case http.MethodDelete: // Delete a federation
		// 1. Delete the federation from the database
		result, err := db.Exec("DELETE FROM federations WHERE id = ?", fedID)
		if err != nil {
			http.Error(w, "Failed to delete federation", http.StatusInternalServerError)
			return
		}

		// 2. Check if any rows were affected (federation existed)
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, "Failed to check deleted rows", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Federation not found", http.StatusNotFound)
			return
		}

		// 3. Send a success response
		w.WriteHeader(http.StatusNoContent) // 204 No Content

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
