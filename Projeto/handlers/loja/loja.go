package loja_handler

import (
    "fmt"
    "net/http"
    "time"
)

var Aberta bool

func logMessage(message string) {
    var Horario = time.Now().Format("02/01/2006 15:04:05")
    fmt.Printf("%s - %s\n", Horario, message)
}

func AbrirLoja(w http.ResponseWriter, r *http.Request) {
    Aberta = true
    logMessage("Loja aberta")
    w.WriteHeader(http.StatusOK)
}

func FecharLoja(w http.ResponseWriter, r *http.Request) {
    Aberta = false
    logMessage("Loja fechada")
    w.WriteHeader(http.StatusOK)
}