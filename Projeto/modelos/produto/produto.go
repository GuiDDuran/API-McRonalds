package produto
var idProduto = 1
type Produto struct {
	Id        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}

func (p *Produto) criaID(){
	p.Id = idProduto
	idProduto++
}
