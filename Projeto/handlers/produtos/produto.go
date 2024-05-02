package produtos_handlers

import (
	"Projeto/modelos/produto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
	var novoProduto produto.Produto

	err := json.NewDecoder(r.Body).Decode(&novoProduto)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados do produto", http.StatusBadRequest)
		return
	}

	produto.LProdutos.Adicionar(novoProduto.Nome, novoProduto.Descricao, novoProduto.Valor)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Produto cadastrado com sucesso!"))
}

func ObterProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	produto, err := produto.LProdutos.BuscarPorId(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func RemoverProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = produto.LProdutos.Remover(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Produto removido com sucesso")
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := produto.LProdutos.ListarProdutos()

	if len(todosOsProdutos) == 0 {

		http.Error(w, "Nenhum produto encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todosOsProdutos)
}
