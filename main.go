package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id			int
	Nome		string
	Descricao	string
	Preco		string
	Quantidade	int
}

var templ = template.Must(template.ParseGlob("templates/*.html"))

func dbConnect() *sql.DB {

	conn := "user=postgres dbname=alura_loja password=senhalocal host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {

		panic(err.Error())

	}

	return db

}

func index(w http.ResponseWriter, r *http.Request) {

	db := dbConnect()

	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {

		panic(err.Error())

	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {

		var id, quantidade 	int
		var nome, descricao	string
		var preco			string

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {

			panic(err.Error())

		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	templ.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()

}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}