package routes

import (
	"net/http"

	"booksapi/handlers"
)

// RegisterRoutes maps endpoints to handler functions
func RegisterRoutes() {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetBooks(w, r)
		case http.MethodPost:
			handlers.AddBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			handlers.UpdateBook(w, r)
		case http.MethodDelete:
			handlers.DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
