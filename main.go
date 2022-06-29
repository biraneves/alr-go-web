package main

import (
	"html/template"
	"net/http"
	"github.com/biraneves/alr-go-web/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaProdutos()
	templ.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}