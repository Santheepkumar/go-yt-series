package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"Text`
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hello: Go"}
	json.NewEncoder(w).Encode(msg)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "User Created"}
	json.NewEncoder(w).Encode(msg)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/api/user", userHandler)
	http.HandleFunc("/api/user2", userHandler)
	http.ListenAndServe(":8000", nil)

}
