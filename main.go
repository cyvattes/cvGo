package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", indexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":"+port, nil)
	// http.ListenAndServe(":8080", nil)
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
