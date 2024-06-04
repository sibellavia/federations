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
	router.HandleFunc("/FHSOperator/FedAdmins", listFedAdmins).Methods(http.MethodGet)

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

	result, err := db.Exec("INSERT INTO fed_admins (member_name, email, description, enabled) VALUES (?, ?, ?, ?)", nil, newFedAdmin.Name, *newFedAdmin.Email, *newFedAdmin.Description, newFedAdmin.Enabled)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var FedAdminInfo FedAdminInfo

	FedAdminInfo.MemberID = strconv.FormatInt(id, 10)

	json.NewEncoder(w).Encode(newFedAdmin)
}

func listFedAdmins(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT member_id, member_name, email, description, enabled, feds_owned FROM fed_admins")
	if err != nil {
		http.Error(w, "Failed to retrieve fed admins", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var fedAdmins []FedAdminInfo

	for rows.Next() {
		var fedAdmin FedAdminInfo
		err := rows.Scan(&fedAdmin.MemberID, &fedAdmin.MemberName, &fedAdmin.Email, &fedAdmin.Description, &fedAdmin.Enabled, &fedAdmin.FedsOwned)
		if err != nil {
			http.Error(w, "Failed to scan Federation Admins", http.StatusInternalServerError)
			return
		}
		fedAdmins = append(fedAdmins, fedAdmin)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fedAdmins)

	rows.Close()
}
