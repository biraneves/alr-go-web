package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome		string
	Descricao	string
	Preco		float64
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

	produtos := []Produto {
		{
			Nome: "Camiseta",
			Descricao: "Azul, bem bonita",
			Preco: 39,
			Quantidade: 5,
		},
		{
			Nome: "Tênis",
			Descricao: "Confortável",
			Preco: 89,
			Quantidade: 3,
		},
		{
			Nome: "Fone",
			Descricao: "Muito bom",
			Preco: 59,
			Quantidade: 2,
		},
		{
			Nome: "Produto novo",
			Descricao: "Muito legal",
			Preco: 1.99,
			Quantidade: 2,
		},
	}

	templ.ExecuteTemplate(w, "Index", produtos)

}

func main() {

	db := dbConnect()
	defer db.Close()
	
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}