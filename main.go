package main

import (
	_ "github.com/DTS-STN/question-priority-service/docs"
	"github.com/DTS-STN/question-priority-service/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger" // echo-swagger middleware

)

// @title Question Prioritization Service
// @version 1.0
// @description This is a service to return questions that when answered will return elegible benefits.

// @host localhost:8080
// @BasePath /

func main() {

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