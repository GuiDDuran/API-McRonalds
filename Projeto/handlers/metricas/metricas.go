package metricas_handlers

import (
	"Projeto/modelos/metricas"
	"encoding/json"
	"net/http"
)

func ObterMetricas(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metricas.Metricas)
}
