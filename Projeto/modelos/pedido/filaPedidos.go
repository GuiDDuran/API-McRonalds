package pedido

import (
	"Projeto/modelos/produto"
)

type FilaPedidos struct {
	Pedidos []Pedido `json:"pedidos"`
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

func (f *FilaPedidos) ListarPedidos() []Pedido{
	return f.Pedidos
}