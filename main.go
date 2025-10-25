package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Book represents a book object
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
const dataFile = "books.json"

// loadBooks reads the JSON file into memory
func loadBooks() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		// If file doesn't exist, create it with sample data
		if os.IsNotExist(err) {
			books = []Book{
				{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan"},
				{ID: "2", Title: "Learning Go", Author: "Jon Bodner"},
			}
			saveBooks()
			return
		}
		log.Fatalf("Error reading file: %v", err)
	}
	json.Unmarshal(file, &books)
}

// saveBooks writes the current books to the JSON file
func saveBooks() {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		log.Printf("Error marshaling books: %v", err)
		return
	}
	ioutil.WriteFile(dataFile, data, 0644)
}

// getBooks returns all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// addBook adds a new book
func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	books = append(books, newBook)
	saveBooks()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// updateBook updates an existing book by ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	var updated Book
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, b := range books {
		if b.ID == id {
			books[i] = updated
			saveBooks()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// deleteBook removes a book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			saveBooks()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
	loadBooks()

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBooks(w, r)
		case http.MethodPost:
			addBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateBook(w, r)
		case http.MethodDelete:
			deleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("📚 Books API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
