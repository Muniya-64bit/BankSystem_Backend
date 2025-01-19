package main

import (
	"fmt"
	"net/http"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Showing books")
	case "POST":
		fmt.Fprintf(w, "Adding a new book")
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to BackEnd!")
	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
