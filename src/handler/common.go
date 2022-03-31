package handler

import (
	"encoding/json"
	"github.com/reeves122/micro-airlines-api-go/repository"
	"net/http"
)

func HealthCheck(repo repository.IRepository, w http.ResponseWriter, _ *http.Request) {
	if err := repo.HealthCheck(); err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(200)
	_, _ = w.Write([]byte("OK\n"))
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
