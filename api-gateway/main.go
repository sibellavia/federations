package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/createFederation", createFederationHandler).Methods("POST")
	r.HandleFunc("/federations", listFederationsHandler).Methods("GET")
	r.HandleFunc("/federations/addService", addServiceHandler).Methods("POST")
	r.HandleFunc("/federations/getServices", getServicesHandler).Methods("GET")
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/sendMessage", sendMessageHandler).Methods("POST")
	r.HandleFunc("/getMessages", getMessagesHandler).Methods("GET")

	log.Println("API Gateway running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

func createFederationHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8081/createFederation", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to create federation", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func listFederationsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/federations")
	if err != nil {
		http.Error(w, "Failed to list federations", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func addServiceHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8081/federations/addService", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to add service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func getServicesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/federations/getServices?" + r.URL.RawQuery)
	if err != nil {
		http.Error(w, "Failed to get services", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8083/register", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8083/login", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to login user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8084/sendMessage", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8084/getMessages")
	if err != nil {
		http.Error(w, "Failed to get messages", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
