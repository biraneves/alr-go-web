package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {

	templ.ExecuteTemplate(w, "Index", nil)

}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}