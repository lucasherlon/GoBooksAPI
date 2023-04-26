package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID          string `json: "id"`
	Title       string `json: "Title"`
	Author      string `json: "Author"`
	Publisher   string `json: "Publisher"`
	Year        string `json: "Year"`
}

// Repository of books
var library []Book

var id int = 0

// Get all the movies in the library.
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(library)
}

// Get one book from the library according to its id.
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range library {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

// Delete a book in the library according to its id.
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, book := range library {
		if book.ID == params["id"] {
			library = append(library[:index], library[index+1:]...)
			break
		}
	}
}

// Create a new book and append it in the library.
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}
	id++
	book.ID = strconv.Itoa(id)
	library = append(library, book)
	json.NewEncoder(w).Encode(book)
}

// Update a book in the library.
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range library {
		if book.ID == params["id"] {
			library = append(library[:index], library[index+1:]...)
			var updatedBook Book
			err := json.NewDecoder(r.Body).Decode(&updatedBook)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				json.NewEncoder(w).Encode("Something went wrong")
				return
			}
			updatedBook.ID = params["id"]
			library = append(library, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	library = append(library, Book{ID: "0", Title: "Example", Author: "Lucas Herlon", Publisher: "Example", Year: "2023"})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Access: localhost:8080/api/books\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
