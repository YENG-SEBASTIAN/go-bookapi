package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books []Book

const dataFile = "data/books.json"

// LoadBooks reads books from the JSON file into memory
func LoadBooks() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			Books = []Book{
				{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan"},
				{ID: "2", Title: "Learning Go", Author: "Jon Bodner"},
			}
			SaveBooks()
			return
		}
		log.Fatalf("Error reading file: %v", err)
	}
	json.Unmarshal(file, &Books)
}

// SaveBooks writes the in-memory data back to JSON file
func SaveBooks() {
	data, err := json.MarshalIndent(Books, "", "  ")
	if err != nil {
		log.Printf("Error marshaling books: %v", err)
		return
	}
	ioutil.WriteFile(dataFile, data, 0644)
}
