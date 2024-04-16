package pedido

import (
	"Projeto/modelos/produto"
	"fmt"
)

type FilaPedidos struct {
	Pedidos []Pedido
}

func (f *FilaPedidos) Adicionar(delivery bool, produtos []produto.Produto, valorTotal float64){
	p := Pedido{
		Delivery: delivery,
		Produtos: produtos,
		ValorTotal: valorTotal,
	}
	p.criaID()
	f.Pedidos = append(f.Pedidos, p)
}
func (f *FilaPedidos) Remover() {
    if len(f.Pedidos) > 0 {
        f.Pedidos = f.Pedidos[1:]
    }
}



func (f *FilaPedidos) ListarPedidos(){

	for _, pedido := range f.Pedidos {
        fmt.Printf("ID: %d, Delivery: %t, Valor: R$ %.2f\n", pedido.Id, pedido.Delivery, pedido.ValorTotal)
        fmt.Println("Produtos:")
        for _, produto := range pedido.Produtos {
            fmt.Printf("Nome: %s, Descrição: %s, Valor: R$ %.2f\n", produto.Nome, produto.Descricao, produto.Valor)
        }
        fmt.Println("------------")
    }

}