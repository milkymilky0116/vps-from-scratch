package handler

import (
	"fmt"
	"net/http"
)

type HealthCheckHandler struct{}

func (h *HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello")
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}
