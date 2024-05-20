package loja_handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var Aberta bool
var IntervaloExpedicao = 30 * time.Second

func logMessage(message string) {
    var Horario = time.Now().Format("02/01/2006 15:04:05")
    fmt.Printf("%s - %s\n", Horario, message)
}

func AbrirLoja(w http.ResponseWriter, r *http.Request) {
	intervaloStr := r.URL.Query().Get("intervalo")
	if intervaloStr != "" {
		intervalo, err := strconv.Atoi(intervaloStr)
		if err != nil || intervalo <= 0 {
			http.Error(w, "Parâmetro 'intervalo' inválido", http.StatusBadRequest)
			return
		}
		IntervaloExpedicao = time.Duration(intervalo) * time.Second
		logMessage(fmt.Sprintf("Intervalo de expedição definido para %d segundos", intervalo))
	} else {
		IntervaloExpedicao = 30 * time.Second
	}

	Aberta = true
	logMessage("Loja aberta")
	w.WriteHeader(http.StatusOK)
}

func FecharLoja(w http.ResponseWriter, r *http.Request) {
    Aberta = false
    logMessage("Loja fechada")
    w.WriteHeader(http.StatusOK)
}