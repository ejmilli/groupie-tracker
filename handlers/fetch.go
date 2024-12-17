package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
}

// Global variable to store artists data
var artists []Artist

func FetchArtistsFromAPI() error {
	// API URL to fetch data
	apiartist := "https://groupietrackers.herokuapp.com/api/artists"

	// Perform the GET request
	resp, err := http.Get(apiartist)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	// Unmarshal JSON data into the `artists` slice
	if err := json.Unmarshal(body, &artists); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	return nil
}
