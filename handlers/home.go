package handlers

import (
	"groupie-tracker/models"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	models.Tmpl.ExecuteTemplate(w, "home", nil)
}
