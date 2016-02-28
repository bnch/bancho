package frontendserver

import (
	"errors"
	"github.com/bnch/bancho/bindata"
	"html/template"
)

// templates for frontend
var templates map[string]*template.Template

const templatesDir = "frontend/templates/"

func setUpTemplates() {
	templates = map[string]*template.Template{
		"signup":       makeTemplate("signup.html", "base.html", "page-status.html"),
		"index_public": makeTemplate("index-public.html", "base.html", "page-status.html"),
	}
}
func makeTemplate(files ...string) *template.Template {
	if len(files) == 0 {
		panic(errors.New("bnch/frontendserver:makeTemplate: no files passed to function"))
	}

	for index, val := range files {
		files[index] = templatesDir + val
	}

	if frontendFolderExists {
		return template.Must(template.ParseFiles(files...))
	}
	// (else...)
	tmpl := template.New(files[0])
	var st string
	for _, file := range files {
		st += string(bindata.MustAsset(file))
	}
	return template.Must(tmpl.Parse(st))
}
