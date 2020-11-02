package server

import (
	"github.com/DTS-STN/question-priority-service/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Main(args []string){
	// Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/healthcheck", handlers.HealthCheck)
	e.POST("/next", handlers.NextQuestion)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
