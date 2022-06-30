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

func Delete(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	templ.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do Id:", err)
		}

		fltPreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		intQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.AtualizaProduto(intId, nome, descricao, fltPreco, intQuantidade)

	}

}