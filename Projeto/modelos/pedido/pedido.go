package pedido

import "Projeto/modelos/produto"

type Pedido struct {
	Id         int
	Delivery   bool
	Produtos   []produto.Produto
	ValorTotal float64
}