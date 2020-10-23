package main

import (
	_ "github.com/DTS-STN/question-priority-service/docs"
	"github.com/DTS-STN/question-priority-service/openfisca"
	"github.com/DTS-STN/question-priority-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger" // echo-swagger middleware

	"fmt"
	"io/ioutil"
	"net/http"
)

// @title Simple Echo Test
// @version 1.0
// @description This is a simple test of the Echo framework.

// @host localhost:8080
// @BasePath /

func main() {

	// Echo instance
	e := echo.New()

	//Declare new custom logger
	logTest := utils.NewLogger()

	// Middleware
	e.Use(logTest.Process)
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.GET("/healthcheck", healthcheck)
	e.POST("/trace", trace)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Healthcheck
// @Summary Returns Healthy
// @Description Returns Healthy
// @ID healthcheck
// @Success 200 {string} string	"Healthy"
// @Router /healthcheck [get]
func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

// OF Trace
// @Summary Send trace request to OpenFisca
// @ID of-trace
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Returns OpenFisca trace response"
// @Router /trace [post]
func trace(c echo.Context) (err error) {

	data := c.Request().Body

	body, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		return
	}

	result, err := openfisca.Trace(body)
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		return
	}

	return c.JSON(http.StatusOK, result)
}
