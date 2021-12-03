package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)


func DbConnection() *sql.DB {
	conexao := //some code here
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
