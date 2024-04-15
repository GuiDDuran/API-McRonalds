package produto

import "fmt"

type ListaProdutos struct {
	Produtos []Produto
}

func (l *ListaProdutos) Adicionar(nome, descricao string, valor float64) {
	p := Produto{
		Nome:      nome,
		Descricao: descricao,
		Valor:     valor,
	}
	p.criaID()
	l.Produtos = append(l.Produtos, p)
}

func (l *ListaProdutos) Remover(idBuscado int) error {
	for i, produto := range l.Produtos {
		if idBuscado == produto.Id {
			l.Produtos = append(l.Produtos[:i], l.Produtos[i+1:]...)
			fmt.Printf("Produto de id %d removido\n", idBuscado)
			return nil
		}
	}

	fmt.Printf("Produto de id %d não localizado\n", idBuscado)
	return fmt.Errorf("id %d não localizado", idBuscado)
}

func (l *ListaProdutos) BuscarPorId(idBuscado int) (*Produto, error) {
	for _, produto := range l.Produtos{
		if idBuscado == produto.Id {
			fmt.Printf("Produto de id %d localizado\n", idBuscado)
			return &produto, nil
		}
	}

	fmt.Printf("Produto de id %d não localizado\n", idBuscado)
	return nil, fmt.Errorf("id %d não localizado", idBuscado)
}

func (l *ListaProdutos) ListarProdutos() {
    for _, produto := range l.Produtos {
        fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: R$ %.2f\n", produto.Id, produto.Nome, produto.Descricao, produto.Valor)
    }
}
