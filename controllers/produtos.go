package controllers

import (
	"html/template"
	"net/http"
	"github.com/biraneves/alr-go-web/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaProdutos()
	templ.ExecuteTemplate(w, "Index", todosOsProdutos)

}