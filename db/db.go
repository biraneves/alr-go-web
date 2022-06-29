package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {

	conn := "user=postgres dbname=alura_loja password=senhalocal host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {

		panic(err.Error())

	}

	return db

}