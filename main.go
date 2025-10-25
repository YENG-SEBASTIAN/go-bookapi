// @title Books API
// @version 1.0
// @description A simple RESTful API to manage books built in Go.
// @host localhost:8080
// @BasePath /


package main

import (
	"fmt"
	"log"
	"net/http"

	"booksapi/models"
	"booksapi/routes"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "booksapi/docs" // generated docs package
)

func main() {
	models.LoadBooks()
	routes.RegisterRoutes()

	// Serve Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("📚 Books API running at http://localhost:8080")
	fmt.Println("📘 Swagger docs available at http://localhost:8080/swagger/index.html")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
