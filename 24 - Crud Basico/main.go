package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", servidor.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", servidor.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", servidor.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", servidor.UpdateUser).Methods(http.MethodPut)

	fmt.Println("Listen on port 5000")
	log.Fatal(http.ListenAndServe("127.0.0.1:5000", router))
}
