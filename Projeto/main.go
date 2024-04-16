package main

import (
	//"Projeto/modelos/pedido"
	"Projeto/modelos/pedido"
	"Projeto/modelos/produto"
	"fmt"
)

func main() {
	pedidos_lista := pedido.FilaPedidos{}
	produtos := produto.ListaProdutos{}
	produtos.Adicionar("batata", "frita", 20.00)
	produtos.Adicionar("ham", "burguer", 20.00)
	pedidos_lista.Adicionar(true, produtos.Produtos, 40.00)
	pedidos_lista.Remover()
	pedidos_lista.Adicionar(true,produtos.Produtos,30.00)
	pedidos_lista.ListarPedidos()
	fmt.Println("\nfim")

	/*
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

			fila := pedido.FilaPedidos{}
			fila.Adicionar(true, lista.Produtos, 15.0)

		    fmt.Println("Fila de Pedidos:")
		    for _, pedido := range fila.Pedidos {
		        fmt.Printf("ID: %d, Delivery: %t, Valor Total: R$ %.2f\n", pedido.Id, pedido.Delivery, pedido.ValorTotal)
		    }*/

}
