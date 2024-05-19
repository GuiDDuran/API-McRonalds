package produto

import (
	"Projeto/modelos/metricas"
	"fmt"
)

type ListaProdutos struct {
    Produtos []Produto `json:"produtos"`
}

var LProdutos ListaProdutos

func (l *ListaProdutos) Adicionar(nome, descricao string, valor float64) {
    p := Produto{
        Nome:      nome,
        Descricao: descricao,
        Valor:     valor,
    }
    p.criaID()
    l.Produtos = append(l.Produtos, p)

    metricas.Metricas.TotalProdutos++
}

func (l *ListaProdutos) Remover(idBuscado int) error {
    for i, produto := range l.Produtos {
        if idBuscado == produto.Id {
            l.Produtos = append(l.Produtos[:i], l.Produtos[i+1:]...)
            metricas.Metricas.TotalProdutos--
            return nil
        }
    }
    return fmt.Errorf("id %d não localizado", idBuscado)
}

func (l *ListaProdutos) BuscarPorId(idBuscado int) (*Produto, error) {
    for _, produto := range l.Produtos {
        if idBuscado == produto.Id {
            return &produto, nil
        }
    }
    return nil, fmt.Errorf("id %d não localizado", idBuscado)
}

func (l *ListaProdutos) ListarProdutos() []Produto {
    return l.Produtos
}
