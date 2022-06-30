package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/biraneves/alr-go-web/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaProdutos()
	templ.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {

	templ.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		fltPreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {

			log.Println("Erro na conversão do preço:", err)

		}

		intQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {

			log.Println("Erro na conversão da quantidade:", err)

		}

		models.NovoProduto(nome, descricao, fltPreco, intQuantidade)

	}

	http.Redirect(w, r, "/", 301)

}