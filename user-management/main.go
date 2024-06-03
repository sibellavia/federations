package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]User)

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	log.Println("User Management Service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func register(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users[user.Username] = user

	json.NewEncoder(w).Encode(map[string]string{"status": "registered"})
}

func login(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if storedUser, exists := users[user.Username]; exists && storedUser.Password == user.Password {
		json.NewEncoder(w).Encode(map[string]string{"status": "logged in"})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": "login failed"})
	}
}
