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
}