package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home", nil)
}
