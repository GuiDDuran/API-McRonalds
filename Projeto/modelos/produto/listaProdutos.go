package produto

import (
	"Projeto/modelos/metricas"
	"fmt"
)

type No struct {
	Produto  Produto
	Esquerda *No
	Direita  *No
}

type ListaProdutos struct {
	Arvore *No
}

var LProdutos ListaProdutos

func InserirNo(raiz *No, produto Produto) *No {
	if raiz == nil {
		return &No{Produto: produto}
	}

	if produto.Nome < raiz.Produto.Nome {
		raiz.Esquerda = InserirNo(raiz.Esquerda, produto)
	} else {
		raiz.Direita = InserirNo(raiz.Direita, produto)
	}

	return raiz
}

func (l *ListaProdutos) Adicionar(nome, descricao string, valor float64) {
	p := Produto{
		Nome:      nome,
		Descricao: descricao,
		Valor:     valor,
	}
	p.criaID()

	l.Arvore = InserirNo(l.Arvore, p)

	metricas.Metricas.TotalProdutos++
}

func encontrarMinimo(no *No) *No {
	for no.Esquerda != nil {
		no = no.Esquerda
	}
	return no
}

func (l *ListaProdutos) BuscarPorId(idBuscado int) (*Produto, error) {
	var buscarProduto func(*No, int) *Produto
	buscarProduto = func(no *No, idBuscado int) *Produto {
		if no == nil {
			return nil
		}
		if no.Produto.Id == idBuscado {
			return &no.Produto
		}
		if produto := buscarProduto(no.Esquerda, idBuscado); produto != nil {
			return produto
		}
		return buscarProduto(no.Direita, idBuscado)
	}

	produtoEncontrado := buscarProduto(l.Arvore, idBuscado)
	if produtoEncontrado == nil {
		return nil, fmt.Errorf("id %d não localizado", idBuscado)
	}
	return produtoEncontrado, nil
}

func (l *ListaProdutos) Remover(idBuscado int) error {
	_, err := l.BuscarPorId(idBuscado)
	if err != nil {
		return err
	}

	var removerProduto func(*No, int) *No
	removerProduto = func(no *No, idBuscado int) *No {
		if no == nil {
			return nil
		}
		if no.Produto.Id == idBuscado {
			if no.Esquerda == nil {
				return no.Direita
			} else if no.Direita == nil {
				return no.Esquerda
			}
			minNo := encontrarMinimo(no.Direita)
			no.Produto = minNo.Produto
			no.Direita = removerProduto(no.Direita, minNo.Produto.Id)
			return no
		}
		no.Esquerda = removerProduto(no.Esquerda, idBuscado)
		no.Direita = removerProduto(no.Direita, idBuscado)
		return no
	}

	l.Arvore = removerProduto(l.Arvore, idBuscado)

	metricas.Metricas.TotalProdutos--
	return nil
}

func (l *ListaProdutos) ListarProdutos() []Produto {
	var produtosOrdenados []Produto
	var emOrdem func(*No)
	emOrdem = func(no *No) {
		if no != nil {
			emOrdem(no.Esquerda)
			produtosOrdenados = append(produtosOrdenados, no.Produto)
			emOrdem(no.Direita)
		}
	}
	emOrdem(l.Arvore)
	return produtosOrdenados
}
