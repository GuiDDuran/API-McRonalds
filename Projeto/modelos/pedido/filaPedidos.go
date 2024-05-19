package pedido

import (
	"Projeto/modelos/metricas"
	"Projeto/modelos/produto"
	"errors"
)

type FilaPedidos struct {
	Pedidos []Pedido `json:"pedidos"`
}

var FPedidos FilaPedidos
var valorTotal float64
var faturamentoTotal float64
var TicketMedio float64

func (f *FilaPedidos) Adicionar(delivery bool, produtosIDs []int) error {
	if len(produtosIDs) == 0 {
		return errors.New("a lista de produtos não pode ser vazia")
	}

	var produtos []produto.Produto
	valorTotal = 0

	for _, id := range produtosIDs {
		prod, err := produto.LProdutos.BuscarPorId(id)
		if err != nil {
			return err
		}
		produtos = append(produtos, *prod)
		valorTotal += prod.Valor
	}

	if delivery{
		valorTotal += 10
	}

	if valorTotal > 100 {
		valorTotal -= valorTotal * 0.10
	}

	faturamentoTotal += valorTotal

	p := Pedido{
		Delivery:   delivery,
		Produtos:   produtos,
		ValorTotal: valorTotal,
	}
	p.criaID()
	f.Pedidos = append(f.Pedidos, p)

	metricas.Metricas.PedidosAndamento++

	return nil
}

func (f *FilaPedidos) Expedir() {
	if len(f.Pedidos) > 0 {
		f.Pedidos = f.Pedidos[1:]
	}

	metricas.Metricas.PedidosAndamento--
	metricas.Metricas.PedidosEncerrados++
	metricas.Metricas.FaturamentoTotal = faturamentoTotal
	metricas.Metricas.TicketMedio = faturamentoTotal / float64(metricas.Metricas.PedidosEncerrados)
}

func (f *FilaPedidos) ListarPedidos() []Pedido {
	return f.Pedidos
}