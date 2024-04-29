package pedidos_handlers

import (
	"Projeto/modelos/pedido"
	"encoding/json"
	"net/http"
)

func AdicionarPedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido struct {
		Delivery    bool  `json:"delivery"`
		ProdutosIDs []int `json:"produtos"`
	}

	err := json.NewDecoder(r.Body).Decode(&novoPedido)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados do pedido", http.StatusBadRequest)
		return
	}

	err = pedido.FPedidos.Adicionar(novoPedido.Delivery, novoPedido.ProdutosIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Pedido cadastrado com sucesso!"))
}

func ObterPedidos(w http.ResponseWriter, r *http.Request) {
	pedidosAtivos := pedido.FPedidos.ListarPedidos()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedidosAtivos)
}
