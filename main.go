package main

import (
	"fmt"
	"log"
	"net/http"

	"booksapi/models"
	"booksapi/routes"
)

func main() {
	// Load data from file
	models.LoadBooks()

	// Register all routes
	routes.RegisterRoutes()

	// Start server
	fmt.Println("📚 Books API is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
