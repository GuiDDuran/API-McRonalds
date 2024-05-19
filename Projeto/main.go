package main

import (
	loja_handler "Projeto/handlers/loja"
	metricas_handlers "Projeto/handlers/metricas"
	pedidos_handlers "Projeto/handlers/pedidos"
	produtos_handlers "Projeto/handlers/produtos"
	"Projeto/processamento"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/produto", produtos_handlers.AdicionarProduto).Methods("POST")
	r.HandleFunc("/produto/{id}", produtos_handlers.ObterProduto).Methods("GET")
	r.HandleFunc("/produto/{id}", produtos_handlers.RemoverProduto).Methods("DELETE")

	r.HandleFunc("/produtos", produtos_handlers.ListarProdutos).Methods("GET")

	r.HandleFunc("/pedido", pedidos_handlers.AdicionarPedido).Methods("POST")
	r.HandleFunc("/pedidos", pedidos_handlers.ObterPedidos).Methods("GET")

	r.HandleFunc("/abrir", loja_handler.AbrirLoja).Methods("POST")
	r.HandleFunc("/fechar", loja_handler.FecharLoja).Methods("POST")

	r.HandleFunc("/metricas", metricas_handlers.ObterMetricas).Methods("GET")

	go processamento.ProcessarPedidos()
	go processamento.ContarTempoDeFuncionamento()

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}