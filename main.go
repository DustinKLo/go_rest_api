package main

import (
	"fmt"
	"log"
	"net/http"

	L "github.com/dustinklo/go_rest_api/functions"
	"github.com/gorilla/mux"
)

func main() {
	// init mux router
	r := mux.NewRouter()

	// route handler / endpoints
	r.HandleFunc("/api/books", L.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", L.GetBook).Methods("GET")
	r.HandleFunc("/api/books", L.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", L.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", L.DeleteBook).Methods("DELETE")

	fmt.Printf("Server running on port %d!\n", 8000)
	server := http.ListenAndServe(":8000", r)
	log.Fatal(server)
}
