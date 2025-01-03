package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	healthCheckRouter := NewHealthCheckHandler()
	req := httptest.NewRequest(http.MethodGet, "/health_check", nil)
	writer := httptest.NewRecorder()
	healthCheckRouter.HealthCheck(writer, req)
	resp := writer.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
