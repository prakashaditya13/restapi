// Adding 3rd party package to add into go mod file commands - go mod init & go mod tidy then go build

package main

import (
	// "encoding/json"
	"encoding/json"
	"log"
	"net/http"

	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Books Model Structure

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Structure
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// init Books variable

var books []Book

// Get /getBooks - Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET /getBook/{id} - Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// POST /createBooks  - Create a Book
func createBooks(w http.ResponseWriter, r *http.Request) {

}

// PUT /updateBooks/{id} - Update a Book
func updateBooks(w http.ResponseWriter, r *http.Request) {

}

// DELETE /deleteBooks/{id}  - Delete a Book
func deleteBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//init Router
	r := mux.NewRouter()

	// Mock Data For Implement database
	books = append(books, Book{ID: "1", Isbn: "234534", Title: "Book One", Author: &Author{FirstName: "Aditya", LastName: "Prakash"}})
	books = append(books, Book{ID: "2", Isbn: "234544", Title: "Book Two", Author: &Author{FirstName: "John", LastName: "Doe"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	// Run a Server
	log.Fatal(http.ListenAndServe(":8000", r))
}
