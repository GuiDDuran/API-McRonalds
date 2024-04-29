package loja_handler


import (
	"fmt"
	"net/http"
)

var Aberta bool

func AbrirLoja(w http.ResponseWriter, r *http.Request) {
	Aberta = true
	fmt.Println("Loja aberta")
	w.WriteHeader(http.StatusOK)
}

func FecharLoja(w http.ResponseWriter, r *http.Request) {
	Aberta = false
	fmt.Println("Loja fechada")
	w.WriteHeader(http.StatusOK)
}