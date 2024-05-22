package pedido

import (
	"Projeto/modelos/metricas"
	"Projeto/modelos/produto"
	"errors"
	"math"
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

	if delivery {
		valorTotal += 10
	}

	if valorTotal > 100 {
		valorTotal -= valorTotal * 0.10
	}

	valorTotal = math.Round(valorTotal*100) / 100

	faturamentoTotal += valorTotal

	faturamentoTotal = math.Round(faturamentoTotal*100) / 100

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

func bubbleSort(pedidos []Pedido) {
	n := len(pedidos)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if pedidos[j].ValorTotal > pedidos[j+1].ValorTotal {
				pedidos[j], pedidos[j+1] = pedidos[j+1], pedidos[j]
			}
		}
	}
}

func quickSort(pedidos []Pedido, inicio, fim int) {
	if inicio < fim {
		i, j := inicio, fim
		pivo := pedidos[inicio]

		for i <= j {
			for pedidos[i].ValorTotal < pivo.ValorTotal {
				i++
			}
			for pivo.ValorTotal < pedidos[j].ValorTotal {
				j--
			}
			if i <= j {
				pedidos[i], pedidos[j] = pedidos[j], pedidos[i]
				i++
				j--
			}
		}

		if inicio < j {
			quickSort(pedidos, inicio, j)
		}
		if i < fim {
			quickSort(pedidos, i, fim)
		}
	}
}

func insertionSort(pedidos []Pedido) {
	n := len(pedidos)
	for i := 1; i < n; i++ {
		chave := pedidos[i]
		j := i - 1

		for j >= 0 && pedidos[j].ValorTotal > chave.ValorTotal {
			pedidos[j+1] = pedidos[j]
			j = j - 1
		}
		pedidos[j+1] = chave
	}
}

func ordenarPedidos(pedidos []Pedido, algoritmo string) []Pedido {
	copiaPedidos := make([]Pedido, len(pedidos))
	copy(copiaPedidos, pedidos)

	switch algoritmo {
	case "bubblesort":
		bubbleSort(copiaPedidos)
	case "quicksort":
		quickSort(copiaPedidos, 0, len(copiaPedidos)-1)
	case "insertionsort":
		insertionSort(copiaPedidos)
	}
	return copiaPedidos
}

func (f *FilaPedidos) ListarPedidos(ordenacao string) []Pedido {
	if ordenacao != "" {
		return ordenarPedidos(f.Pedidos, ordenacao)
	}
	return f.Pedidos
}