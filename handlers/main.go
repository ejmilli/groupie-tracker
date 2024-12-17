package main

import (
	"fmt" // Replace with the correct import path for handlers
	"net/http"
)

func main() {
	if err := fetchArtistsFromAPI(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Artists data fetched successfully!")

	// Handle API route
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/api/artists", HandleArtists)

	// Start the server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
