package main

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Return the artists data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	models.Tmpl.ExecuteTemplate(w, "home", nil)
}
