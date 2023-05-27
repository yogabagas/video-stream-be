package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode("OK"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
