package frontendserver

import (
	"html/template"
)

// templates for frontend
var templates map[string]*template.Template

const templatesDir = "frontend/templates/"

func setUpTemplates() {
	templates = map[string]*template.Template{
		"signup": template.Must(template.ParseFiles(
			templatesDir+"signup.html",
			templatesDir+"base.html",
			templatesDir+"page-status.html",
		)),
		"index_public": template.Must(template.ParseFiles(
			templatesDir+"index-public.html",
			templatesDir+"base.html",
			templatesDir+"page-status.html",
		)),
	}
}
