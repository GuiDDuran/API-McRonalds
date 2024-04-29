package pedido

import (
	"Projeto/modelos/produto"
)

type FilaPedidos struct {
	Pedidos []Pedido `json:"pedidos"`
}

var FPedidos FilaPedidos

func (f *FilaPedidos) Adicionar(delivery bool, produtosIDs []int) error {
	var produtos []produto.Produto
	var valorTotal float64

	for _, id := range produtosIDs {
		prod, err := produto.LProdutos.BuscarPorId(id)
		if err != nil {
			return err // Se houver um erro, retorne imediatamente
		}
		produtos = append(produtos, *prod)
		valorTotal += prod.Valor
	}

	p := Pedido{
		Delivery:   delivery,
		Produtos:   produtos,
		ValorTotal: valorTotal,
	}
	p.criaID()
	f.Pedidos = append(f.Pedidos, p)

	return nil
}

func (f *FilaPedidos) Expedir() {
	if len(f.Pedidos) > 0 {
		f.Pedidos = f.Pedidos[1:]
	}
}

func (f *FilaPedidos) ListarPedidos() []Pedido {
	return f.Pedidos
}
