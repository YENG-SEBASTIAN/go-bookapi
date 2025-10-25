package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"booksapi/models"
)

// GetBooks godoc
// @Summary Get all books
// @Description Retrieve a list of all available books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Books)
}

// AddBook godoc
// @Summary Add a new book
// @Description Create a new book entry
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book data"
// @Success 201 {object} models.Book
// @Router /books [post]
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

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Updated book"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
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

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Param id path string true "Book ID"
// @Success 204
// @Router /books/{id} [delete]
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
