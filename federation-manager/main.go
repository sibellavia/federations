package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Federation struct represents a federation entity
type Federation struct {
	ID   string `json:"id"`   // ID
	Name string `json:"name"` // specifies JSON keys when encoding and decoding
}

// Map to store federations with federation IDs as keys and Federation structs as values
var federations = make(map[string]Federation)

func main() {
	http.HandleFunc("/createFederation", createFederation)
	http.HandleFunc("/federations", listFederations)

	log.Println("Federation Manager Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// createFederation handles federation creation requests
func createFederation(w http.ResponseWriter, r *http.Request) {
	var fed Federation
	_ = json.NewDecoder(r.Body).Decode(&fed)
	fed.ID = fmt.Sprintf("fed-%d", len(federations)+1)
	federations[fed.ID] = fed

	json.NewEncoder(w).Encode(fed)
}

// listFederations handles requests to list all federations
func listFederations(w http.ResponseWriter, r *http.Request) {
	feds := make([]Federation, 0, len(federations))
	for _, fed := range federations {
		feds = append(feds, fed)
	}

	json.NewEncoder(w).Encode(feds)
}
