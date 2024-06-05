package pedidos_handlers

import (
	"Projeto/modelos/pedido"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func logMessage(message string) {
	Horario := time.Now().Format("02/01/2006 15:04:05")
    fmt.Printf("%s - %s\n", Horario, message)
}

func AdicionarPedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido struct {
		Delivery    bool  `json:"delivery"`
		ProdutosIDs []int `json:"produtos"`
	}

	err := json.NewDecoder(r.Body).Decode(&novoPedido)
	if err != nil {
		logMessage("Erro ao decodificar dados do pedido")
		http.Error(w, "Erro ao decodificar dados do pedido", http.StatusBadRequest)
		return
	}

	if len(novoPedido.ProdutosIDs) == 0 {
		logMessage("A lista de produtos não pode ser vazia")
		http.Error(w, "A lista de produtos não pode ser vazia", http.StatusBadRequest)
		return
	}

	err = pedido.FPedidos.Adicionar(novoPedido.Delivery, novoPedido.ProdutosIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	logMessage("Pedido cadastrado com sucesso!")
	w.Write([]byte("Pedido cadastrado com sucesso!"))
}

func ObterPedidos(w http.ResponseWriter, r *http.Request) {
	ordenacao := r.URL.Query().Get("ordenacao")

	pedidosAtivos := pedido.FPedidos.ListarPedidos(ordenacao)

	if len(pedidosAtivos) == 0 {
		logMessage("Não há pedidos ativos")
		http.Error(w, "Não há pedidos ativos", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	logMessage("Obtendo pedidos")
	json.NewEncoder(w).Encode(pedidosAtivos)
}