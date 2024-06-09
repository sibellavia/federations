package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type InstanceRegistration struct {
	InstanceID       *int   `json:"instance_id"`
	InstanceIP       string `json:"instance_ip"`
	InstancePort     *int   `json:"instance_port"`
	InstanceMetadata string `json:"instance_metadata"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "../federation-management/federations.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	createInstancesTable := `
    CREATE TABLE IF NOT EXISTS instances (
        instance_id INTEGER PRIMARY KEY,
        instance_ip TEXT NOT NULL,
		instance_port INTEGER NOT NULL,
		instance_metadata TEXT
    );
    `

	_, err = db.Exec(createInstancesTable)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createInstancesTable)
		return
	}

	router := mux.NewRouter()

	// FCP 0.1 APIs
	router.HandleFunc("/register", handleRegister).Methods(http.MethodPost)
	// router.HandleFunc("/heartbeat", handleHeartbeat).Methods(http.MethodPost)

	// Service running
	log.Println("Federation Communication Protocol servers are running on port 8087")
	log.Fatal(http.ListenAndServe(":8087", router))
}

// handleRegister handles requests to /register/
func handleRegister(w http.ResponseWriter, r *http.Request) {
	var newInstance InstanceRegistration
	err := json.NewDecoder(r.Body).Decode(&newInstance)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if newInstance.InstanceID == nil {
		http.Error(w, "Mandatory field 'instance_id' is required", http.StatusBadRequest)
		return
	}

	if newInstance.InstanceIP == "" {
		http.Error(w, "Mandatory field 'instance_ip' is required", http.StatusBadRequest)
		return
	}

	if newInstance.InstancePort == nil {
		http.Error(w, "Mandatory field 'instance_port' is required", http.StatusBadRequest)
		return
	}

	// executes SQL insert statement
	result, err := db.Exec("INSERT INTO instances (instance_id, instance_ip, instance_port, instance_metadata) VALUES (?, ?, ?, ?)",
		newInstance.InstanceID, newInstance.InstanceIP, newInstance.InstancePort, newInstance.InstanceMetadata)

	if err != nil {
		http.Error(w, "Failed to add new instance to DB", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last inserted ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	log.Println(id)
}
