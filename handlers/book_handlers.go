package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"booksapi/models"
)

// GetBooks - returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Books)
}

// AddBook - adds a new book
func AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	models.Books = append(models.Books, newBook)
	models.SaveBooks()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// UpdateBook - updates a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	var updated models.Book
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, b := range models.Books {
		if b.ID == id {
			models.Books[i] = updated
			models.SaveBooks()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// DeleteBook - removes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	for i, b := range models.Books {
		if b.ID == id {
			models.Books = append(models.Books[:i], models.Books[i+1:]...)
			models.SaveBooks()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
