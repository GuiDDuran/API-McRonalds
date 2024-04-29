package metricas_handlers

import (
	"encoding/json"
	"net/http"
	"Projeto/modelos/metricas"
)

func ObterMetricas(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metricas.Metricas)
}