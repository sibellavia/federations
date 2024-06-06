package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type NewFederation struct {
	MemberID              *int   `json:"member_id"`       // Mandatory
	FederationName        string `json:"fed_name"`        // Mandatory
	FederationDescription string `json:"fed_description"` // Optional
	Enabled               *bool  `json:"enabled"`         // Mandatory
}

type Federation struct {
	FederationID          int    `json:"fed_id"`
	FederationName        string `json:"fed_name"`        // Mandatory
	FederationDescription string `json:"fed_description"` // Optional
	MemberID              *int   `json:"member_id"`       // Mandatory
	Enabled               *bool  `json:"enabled"`         // Mandatory
}

type Service struct {
	ServiceID          int    `json:"service_id"`
	FederationID       int    `json:"fed_id"`
	ServiceName        string `json:"service_name"`
	ServiceDescription string `json:"service_description"`
}

// A global variable to hold the database connection
var db *sql.DB

// Main opens the SQLite database, creates federations and services tables if they don't exist, sets up HTTP routes, and starts the HTTP server.
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
        fed_id INTEGER PRIMARY KEY AUTOINCREMENT,
        fed_name TEXT NOT NULL,
		fed_description TEXT,
		member_id INTEGER,
		enabled BOOLEAN NOT NULL,
		FOREIGN KEY (member_id) REFERENCES fed_admins(member_id)
    );
    `
	_, err = db.Exec(createFederationsTable)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createFederationsTable)
		return
	}

	createServicesTable := `
    CREATE TABLE IF NOT EXISTS services (
        service_id INTEGER PRIMARY KEY AUTOINCREMENT,
        fed_id INTEGER,
        service_name TEXT NOT NULL,
        service_description TEXT,
        FOREIGN KEY (fed_id) REFERENCES federations(fed_id)
    );
    `
	_, err = db.Exec(createServicesTable)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createServicesTable)
		return
	}

	// Create a mux router for handling HTTP requests
	router := mux.NewRouter()

	// IEEE-2302-2021 :: Member-FHS API

	// 4. Federation
	router.HandleFunc("/federations", handleFederations).Methods(http.MethodPost, http.MethodGet)

	// Service running
	log.Println("Federation Management Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

// handleFederations handles requests to /federations/
// handles POST requests to create a new federation and GET requests to retrieve all federations.
func handleFederations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	// Decodes the request body to a Federation struct, inserts it into the database, and returns the created federation.
	case http.MethodPost:
		var newFederation NewFederation
		err := json.NewDecoder(r.Body).Decode(&newFederation)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if newFederation.MemberID == nil {
			http.Error(w, "Mandatory field 'member_id' is required", http.StatusBadRequest)
			return
		}

		if newFederation.FederationName == "" {
			http.Error(w, "Mandatory field 'fed_name' is required", http.StatusBadRequest)
			return
		}

		if newFederation.Enabled == nil {
			http.Error(w, "Mandatory field 'enabled' is required", http.StatusBadRequest)
			return
		}

		// executes SQL insert statement
		result, err := db.Exec("INSERT INTO federations (fed_name, fed_description, enabled, member_id) VALUES (?, ?, ?, ?)",
			newFederation.FederationName, newFederation.FederationDescription, *newFederation.Enabled, *newFederation.MemberID)
		if err != nil {
			http.Error(w, "Failed to create federation", http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
			return
		}

		// Create the response struct
		federation := Federation{
			FederationID:          int(id),
			FederationName:        newFederation.FederationName,
			FederationDescription: newFederation.FederationDescription,
			MemberID:              newFederation.MemberID,
			Enabled:               newFederation.Enabled,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(federation)

	// Queries the database for all federations and returns them as a JSON response.
	case http.MethodGet:

		rows, err := db.Query("SELECT * FROM federations")
		if err != nil {
			http.Error(w, "Failed to retrieve federations", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var federations []Federation
		for rows.Next() {
			var federation Federation
			err := rows.Scan(&federation.FederationID, &federation.FederationName, &federation.FederationDescription, &federation.MemberID, &federation.Enabled)
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
