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

func NovoProduto(nome, descricao string, preco float64, quantidade int) {

	db := database.DbConnect()

	novoProduto, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {

		panic(err.Error())

	}

	novoProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()

}

func DeletaProduto(id string) {

	db := database.DbConnect()

	deletarProduto, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		
		panic(err.Error())

	}

	deletarProduto.Exec(id)
	defer db.Close()

}

func EditaProduto(id string) Produto {

	db := database.DbConnect()

	editarProduto, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {

		panic(err.Error())

	}
	
	produto := Produto{}

	for editarProduto.Next() {

		var id, quantidade int
		var nome, descricao, preco string

		err = editarProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {

			panic(err.Error())

		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

	}

	return produto

}

func AtualizaProduto(id, quantidade int, nome, descricao string, preco float64) {

	db := database.DbConnect()

	atualizaProduto, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")

	if err != nil {

		panic(err.Error())

	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}