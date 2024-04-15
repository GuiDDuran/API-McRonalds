package main

import (
	"Projeto/modelos/produto"
	"fmt"
)

func main(){
	lista := produto.ListaProdutos{}
	fmt.Println(lista)
	lista.Adicionar("Batata", "Frita", 20.0)
	fmt.Println(lista)
	lista.Adicionar("Batata", "Cozida", 25.0)
	fmt.Println(lista)

	err := lista.Remover(1)
	fmt.Println(lista)
	fmt.Println(err)
	err = lista.Remover(4)
	fmt.Println(lista)
	fmt.Println(err)

	produtoBuscado1, err := lista.BuscarPorId(2)
	fmt.Println(produtoBuscado1)
	fmt.Println(err)

	produtoBuscado2, err := lista.BuscarPorId(4)
	fmt.Println(produtoBuscado2)
	fmt.Println(err)

	lista.Adicionar("Mandioca", "Frita", 27.0)
	lista.Adicionar("Banana", "Frita", 15.0)
	lista.Adicionar("Batata Doce", "Frita", 32.0)
	lista.ListarProdutos()
}