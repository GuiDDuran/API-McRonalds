package pedidos_handlers

import (
    "encoding/json"
    // "fmt"
    "net/http"
    "Projeto/modelos/pedido"
)

func AdicionarPedido(w http.ResponseWriter, r *http.Request) {
    var novoPedido pedido.Pedido

    err := json.NewDecoder(r.Body).Decode(&novoPedido)
    if err != nil {
        http.Error(w, "Erro ao decodificar dados do pedido", http.StatusBadRequest)
        return
    }

    pedido.FPedidos.Adicionar(novoPedido.Delivery, novoPedido.Produtos, novoPedido.ValorTotal)

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/html; charset=UTF-8")
    w.Write([]byte("Pedido cadastrado com sucesso!"))
}

// func ExpedirPedido(w http.ResponseWriter, r *http.Request) {
//     pedido.FPedidos.Expedir()

//     w.WriteHeader(http.StatusOK)
//     w.Header().Set("Content-Type", "application/json")
//     fmt.Println("Pedido expedido com sucesso")
// }

func ObterPedidos(w http.ResponseWriter, r *http.Request) {
    pedidosAtivos := pedido.FPedidos.ListarPedidos()

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(pedidosAtivos)
}