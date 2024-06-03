package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Federation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./federations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS federations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
		return
	}

	http.HandleFunc("/createFederation", createFederation)
	http.HandleFunc("/federations", listFederations)

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
		federations = append(federations, fed)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(federations)
}
