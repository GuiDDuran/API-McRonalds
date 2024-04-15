package pedido

import "Projeto/modelos/produto"

var idPedido = 1

type Pedido struct {
	Id         int 				 `json:"id"`
	Delivery   bool				 `json:"delivery"`
	Produtos   []produto.Produto `json:"produtos"`
	ValorTotal float64			 `json:"valortotal"`
}

func (p *Pedido) criaID(){
	p.Id = idPedido
	idPedido++
}