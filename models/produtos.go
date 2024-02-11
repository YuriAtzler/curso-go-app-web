package models

import "curso/db"

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConnectDb()
	defer db.Close()

	queryProducts, err := db.Query("SELECT * FROM produtos;")
	if err != nil {
		panic(err)
	}
	p := Produto{}
	produtos := []Produto{}

	for queryProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := queryProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err)
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()
	defer db.Close()

	insertProduct, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4);")
	if err != nil {
		panic(err)
	}
	insertProduct.Exec(nome, descricao, preco, quantidade)
}

func DeletaProduto(idProduto string) {
	db := db.ConnectDb()
	defer db.Close()

	deleteProduct, err := db.Prepare("DELETE FROM produtos WHERE id = $1;")
	if err != nil {
		panic(err)
	}
	deleteProduct.Exec(idProduto)
}

func BuscaProduto(idProduto string) Produto {

	db := db.ConnectDb()
	defer db.Close()

	product, err := db.Query("SELECT * FROM produtos WHERE id = $1;", idProduto)
	if err != nil {
		panic(err)
	}
	p := Produto{}

	for product.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := product.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err)
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}

	return p
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()
	defer db.Close()

	updateProduct, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5;")
	if err != nil {
		panic(err)
	}
	updateProduct.Exec(nome, descricao, preco, quantidade, id)
}
