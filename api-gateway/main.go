package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Federation Management Endpoints
	router.HandleFunc("/federations", handleFederations).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/federations/{fed_id}", handleFederationByID).Methods(http.MethodGet, http.MethodDelete)

	log.Println("API Gateway running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func handleFederations(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		resp, err := http.Post("http://localhost:8081/federations", "application/json", r.Body)
		if err != nil {
			http.Error(w, "Failed to create federation", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Forward the response from federation_management to the client
		forwardResponse(w, resp)
	} else if r.Method == http.MethodGet {
		resp, err := http.Get("http://localhost:8081/federations")
		if err != nil {
			http.Error(w, "Failed to list federations", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Forward the response from federation_management to the client
		forwardResponse(w, resp)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleFederationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fedID, err := strconv.Atoi(vars["fed_id"])
	if err != nil {
		http.Error(w, "Invalid federation ID", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		resp, err := http.Get("http://localhost:8081/federations/" + strconv.Itoa(fedID))
		if err != nil {
			http.Error(w, "Failed to get federation", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Forward the response from federation_management to the client
		forwardResponse(w, resp)
	} else if r.Method == http.MethodDelete {
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8081/federations/"+strconv.Itoa(fedID), nil)
		if err != nil {
			http.Error(w, "Failed to create delete request", http.StatusInternalServerError)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to delete federation", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Forward the response from federation_management to the client
		forwardResponse(w, resp)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Helper function to forward responses from federation_management to the client
func forwardResponse(w http.ResponseWriter, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
