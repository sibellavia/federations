package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// NewFedAdmin struct for NewFedAdmin
type NewFedAdmin struct {
	Name        string  `json:"name"`
	Email       *string `json:"email,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled     bool    `json:"enabled"`
}

type FedAdminInfo struct {
	MemberID    string         `json:"member_id" yaml:"member_id"`
	MemberName  string         `json:"member_name" yaml:"member_name"`
	Email       string         `json:"email" yaml:"email"`
	Description string         `json:"description" yaml:"description"`
	Enabled     bool           `json:"enabled" yaml:"enabled"`
	FedsOwned   []FederationID `json:"feds_owned" yaml:"feds_owned"`
}

type FederationID struct {
	FedID string `json:"fed_id"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "../federation-management/federations.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables if not exists
	createFedAdminsTable := `
	CREATE TABLE IF NOT EXISTS fed_admins (
		member_id INTEGER PRIMARY KEY AUTOINCREMENT,
        member_name TEXT NOT NULL,
		email TEXT,
		description TEXT,
		enabled BOOLEAN NOT NULL,
		feds_owned 
	);
	`

	_, err = db.Exec(createFedAdminsTable)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createFedAdminsTable)
		return
	}

	createFedsOwnedTable := `
    CREATE TABLE IF NOT EXISTS fed_admins_fed_owned (
        fed_id INTEGER,
        member_id INTEGER,
     PRIMARY KEY (fed_id, member_id)
    );
    `

	_, err = db.Exec(createFedsOwnedTable)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createFedsOwnedTable)
		return
	}

	// Create a mux router for handling HTTP requests
	router := mux.NewRouter()

	// IEEE-2302-2021 :: FHS-Operator API

	// 1. FHSOperator Core
	router.HandleFunc("/FHSOperator/NewFedAdmin", handleNewFedAdmin).Methods(http.MethodPost)

	// Service running
	log.Println("Federation Member Management Service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}

func handleNewFedAdmin(w http.ResponseWriter, r *http.Request) {
	var newFedAdmin NewFedAdmin

	// Parse JSON request body into a NewFedAdmin struct
	err := json.NewDecoder(r.Body).Decode(&newFedAdmin)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate mandatory fields
	if newFedAdmin.Name == "" || !newFedAdmin.Enabled {
		http.Error(w, "Mandatory fields 'name' and 'enabled' are required", http.StatusBadRequest)
		return
	}

	// Insert the new fedAdmin into the db
	db, err := sql.Open("sqlite3", "../federation-management/federations.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	insertQuery := `
    INSERT INTO fed_admins (member_name, enabled) 
    VALUES (?, ?);
    `

	result, err := db.Exec(insertQuery, newFedAdmin.Name, newFedAdmin.Enabled)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the newly inserted member_id
	// memberId, err := result.LastInsertId()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }

	// Return the newly inserted Fed Admin as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
