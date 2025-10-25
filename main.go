package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)


type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

// books is our in-memory data store
var books = []Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan"},
	{ID: "2", Title: "Learning Go", Author: "Jon Bodner"},
}

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Route handlers
	http.HandleFunc("/books", getBooks)

	// Start server
	fmt.Println("📚 Books API is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}