package main

import (
	"github.com/DTS-STN/question-priority-service/utils"
	_ "github.com/DTS-STN/question-priority-service/docs"
	"github.com/swaggo/echo-swagger" // echo-swagger middleware
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	"fmt"
	"io/ioutil"
	"log"
	"bytes"
	"encoding/json"
	"net/http"
	"sync/atomic"
)

// @title Simple Echo Test
// @version 1.0
// @description This is a simple test of the Echo framework.

// @host localhost:1234
// @BasePath /

// Declare a local userCount
var userCount int64

// increments the userCount using the atomic library so it can only be done one at a time
func incUserCount() int64 {
	return atomic.AddInt64(&userCount, 1)
}

// returns the current userCount
func getUserCount() int64 {
	return atomic.LoadInt64(&userCount)
}

func main() {

	// Echo instance
	e := echo.New()

	//Declare new custom logger
	logTest := utils.NewLogger()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Middleware
	e.Use(logTest.Process)
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.GET("/hello", hello)
	e.POST("/trace", trace)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// Hello
// @Summary Returns Hello, World!
// @Description Returns Hello, World!
// @ID hello
// @Success 200 {string} string	"ok"
// @Router /hello [get]
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
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
		log.Printf("Error reading body: %v", err)
		// http.Error("can't read body", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:5000/trace", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error in req: ", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in request: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Printf("Error reading body: %v", err)
		//http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	json := json.RawMessage(body)

	return c.JSON(http.StatusOK, json)
}
