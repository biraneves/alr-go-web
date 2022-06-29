package models

import (
	database "github.com/biraneves/alr-go-web/db"
)

type Produto struct {
	Id			int
	Nome		string
	Descricao	string
	Preco		string
	Quantidade	int
}

func BuscaProdutos() []Produto {

	db := database.DbConnect()

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

	defer db.Close()
	return produtos

}