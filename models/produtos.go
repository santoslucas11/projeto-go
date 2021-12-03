package models

import (
	"projeto-go/db"
	_ "projeto-go/db"
)

type Produto struct {
	Id int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SearchAllProdutcs() [] Produto  {
	dataBase := db.DbConnection()

	selectAllProducts, err := dataBase.Query("select * from produtos")
	if err != nil{
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAllProducts.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer dataBase.Close()
	return produtos
}
