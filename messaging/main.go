package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

var messages []Message

func main() {
	http.HandleFunc("/sendMessage", sendMessage)
	http.HandleFunc("/getMessages", getMessages)

	log.Println("Messaging Service running on port 8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	_ = json.NewDecoder(r.Body).Decode(&msg)
	messages = append(messages, msg)

	json.NewEncoder(w).Encode(map[string]string{"status": "message sent"})
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(messages)
}
