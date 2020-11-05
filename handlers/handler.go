package handlers

import "github.com/labstack/echo/v4"

type HandlerServiceInterface interface {
	HealthCheck(c echo.Context) error
	NextQuestion(c echo.Context) (err error)
}

type Handler struct {
}

var HandlerService HandlerServiceInterface = new(Handler)