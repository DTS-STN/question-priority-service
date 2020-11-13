package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Healthcheck
// @Summary Returns Healthy
// @Description Returns Healthy
// @ID healthcheck
// @Success 200 {string} string	"Healthy"
// @Router /healthcheck [get]
func (h *Handler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}
