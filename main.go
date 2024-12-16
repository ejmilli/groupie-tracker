package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Groupie Tracker!")
}

func main() {

	http.HandleFunc("/", homeHandler)

	// Define the port to listen on

	fmt.Printf("Server is running on http://localhost:8080")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: 8080", err)
	}
}
