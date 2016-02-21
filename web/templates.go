package web

import (
	"html/template"
)

// templates for frontend
var templates map[string]*template.Template

func setUpTemplates() {
	templates = map[string]*template.Template{
		"signup": template.Must(template.ParseFiles("frontend/templates/signup.html", "frontend/templates/base.html")),
	}
}
