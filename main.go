package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := "templates"
	tmpl := "index.html"
	t := filepath.Join(tmplPath, tmpl)
	renderTemplate(w, t, tmpl)
}

func renderTemplate(w http.ResponseWriter, t string, tmpl string) {
	var templates = template.Must(template.ParseFiles(t))
	err := templates.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
