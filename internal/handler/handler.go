package handler

import (
	"net/http"
)

// HealthCheck godoc
// @Summary Health check
// @Description Returns 200 OK if the service is up
// @Tags system
// @Success 200 {string} string "ok"
// @Router /healthz [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
