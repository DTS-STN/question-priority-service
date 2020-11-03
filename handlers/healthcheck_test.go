package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HealthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
