package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

	http.HandleFunc("/createFederation", createFederation)
	http.HandleFunc("/federations", listFederations)
	http.HandleFunc("/federations/addService", addService)
	http.HandleFunc("/federations/getServices", getServices)

	log.Println("Federation Manager Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func createFederation(w http.ResponseWriter, r *http.Request) {
	var fed Federation
	_ = json.NewDecoder(r.Body).Decode(&fed)

	stmt, err := db.Prepare("INSERT INTO federations(name) VALUES(?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	res, err := stmt.Exec(fed.Name)
	if err != nil {
		http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		return
	}

	fed.ID = int(id)
	json.NewEncoder(w).Encode(fed)
}

func listFederations(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM federations")
	if err != nil {
		http.Error(w, "Failed to query federations", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var federations []Federation
	for rows.Next() {
		var fed Federation
		err = rows.Scan(&fed.ID, &fed.Name)
		if err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}

		// Retrieve services for the federation
		services, err := getServicesByFederationID(fed.ID)
		if err != nil {
			http.Error(w, "Failed to retrieve services", http.StatusInternalServerError)
			return
		}
		fed.ServiceCatalogue = services

		federations = append(federations, fed)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(federations)
}

func addService(w http.ResponseWriter, r *http.Request) {
	var svc Service
	_ = json.NewDecoder(r.Body).Decode(&svc)

	stmt, err := db.Prepare("INSERT INTO services(federation_id, name, description) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	res, err := stmt.Exec(svc.FederationID, svc.Name, svc.Description)
	if err != nil {
		http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		return
	}

	svc.ID = int(id)
	json.NewEncoder(w).Encode(svc)
}

func getServices(w http.ResponseWriter, r *http.Request) {
	federationIDStr := r.URL.Query().Get("federation_id")
	if federationIDStr == "" {
		http.Error(w, "Missing federation_id parameter", http.StatusBadRequest)
		return
	}

	federationID, err := strconv.Atoi(federationIDStr)
	if err != nil {
		http.Error(w, "Invalid federation_id parameter", http.StatusBadRequest)
		return
	}

	services, err := getServicesByFederationID(federationID)
	if err != nil {
		http.Error(w, "Failed to retrieve services", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(services)
}

func getServicesByFederationID(federationID int) ([]Service, error) {
	rows, err := db.Query("SELECT id, name, description FROM services WHERE federation_id = ?", federationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []Service
	for rows.Next() {
		var svc Service
		err = rows.Scan(&svc.ID, &svc.Name, &svc.Description)
		if err != nil {
			return nil, err
		}
		services = append(services, svc)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return services, nil
}
