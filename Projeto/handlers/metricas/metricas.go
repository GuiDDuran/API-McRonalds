package metricas_handlers

import (
    "Projeto/modelos/metricas"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

func ObterMetricas(w http.ResponseWriter, r *http.Request) {
	var Horario = time.Now().Format("02/01/2006 15:04:05")
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    fmt.Printf("%s - Obtendo métricas\n", Horario)
    json.NewEncoder(w).Encode(metricas.Metricas)
}