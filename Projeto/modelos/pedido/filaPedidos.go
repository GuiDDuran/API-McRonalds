package pedido

import "Projeto/modelos/produto"

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