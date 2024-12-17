package models

import (
	"html/template"
)

var Tmpl *template.Template

func init() {
    
    
    Tmpl = template.Must(template.ParseGlob("templates/*.html"))
   
}
