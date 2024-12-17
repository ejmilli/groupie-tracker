package main

import (
	"fmt" // Replace with the correct import path for handlers
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	if err := handlers.FetchArtistsFromAPI(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Artists data fetched successfully!")

	// Handle API route
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists", handlers.HandleArtists)

	// Start the server
	port := ":10000"
	fmt.Println("Server running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
