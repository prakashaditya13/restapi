package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Books Model Structure

type Book struct{
	ID 			string		`json:"id"`
	Isbn 		string		`json:"isbn"`
	Title 		string		`json:"title"`
	Author 		*Author		`json:"author"`

}

// Author Structure
type Author struct{
	FirstName		string		`json:"firstname"`
	LastName		string		`json:"lastname"`

}

// Get /getBooks
func getBooks(w http.ResponseWriter, r *http.Request){

}

// GET /getBook/{id}
func getBook(w http.ResponseWriter, r *http.Request){

}

// POST /createBooks
func createBooks(w http.ResponseWriter, r *http.Request){

}

// PUT /updateBooks/{id}
func updateBooks(w http.ResponseWriter, r *http.Request){

}

// DELETE /deleteBooks/{id}
func deleteBooks(w http.ResponseWriter, r *http.Request){

}


func main() {
	//init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	// Run a Server
	log.Fatal(http.ListenAndServe(":8000", r))
}
