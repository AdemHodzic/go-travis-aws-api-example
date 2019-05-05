package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type Message struct {
	Msg string `json:"msg"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	response := &Message{Msg: "Hello, World!"}
	result := json.NewEncoder(w).Encode(response)
	fmt.Printf("%v", result)
}

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/", Index);
	log.Fatal(http.ListenAndServe(":8080", router))
}