package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	tpl = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", handlers.HomeHandler)

	fmt.Printf("Server is running on http://localhost:10000")

	http.ListenAndServe(":10000", nil)

}
