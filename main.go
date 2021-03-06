package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Message struct {
	Msg string `json:"msg"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	response := &Message{Msg: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func Another(w http.ResponseWriter, r *http.Request) {
	response := &Message{Msg: "Another Hello, World!"}
	json.NewEncoder(w).Encode(response)
}


func main() {
	router := mux.NewRouter();
	router.HandleFunc("/", Index);
	router.HandleFunc("/another", Another);
	log.Println("RUNNING ON PORT 8080");
	log.Fatal(http.ListenAndServe(":8080", router));
}