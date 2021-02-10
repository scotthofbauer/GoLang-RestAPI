package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struck (model)
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

//Author Struck
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

func main() {
	// Init Router
	r := mux.NewRouter()

	//MockData - @todo -  implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "847634", Title: "Book Two", Author: Author{Firstname: "Jane", Lastname: "Doe"}})
	books = append(books, Book{ID: "3", Isbn: "901734", Title: "Book Three", Author: Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
