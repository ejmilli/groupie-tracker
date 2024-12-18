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

<<<<<<< HEAD
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	models.Tmpl = template.Must(template.ParseGlob("templates/*.html"))

=======
	// Handle API route
>>>>>>> b6da0aaa11ee93b0e6c5543c3ec2626e963ab5a9
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists", handlers.HandleArtists)

	// Start the server
	port := ":10000"
	fmt.Println("Server running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
