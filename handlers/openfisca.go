package handlers

import (
	"bytes"
	"github.com/DTS-STN/question-priority-service/bindings"
	"github.com/DTS-STN/question-priority-service/renderings"
	"github.com/labstack/echo/v4"
	"gopkg.in/square/go-jose.v2/json"

	"net/http"
	"fmt"
)

type OpenFiscaInterface interface {
	sendRequest(*bindings.TraceRequest) (renderings.TraceResponse, error)
}

type OpenFisca struct {}
var openFisca OpenFiscaInterface

// OF Trace
// @Summary Send trace request to OpenFisca
// @ID of-trace
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Returns OpenFisca trace response"
// @Router /trace [post]
func Trace(c echo.Context) (err error) {
	resp := renderings.TraceResponse{}
	traceRequest := new(bindings.TraceRequest)

	// Bind the request into our request struct
	if err := c.Bind(traceRequest); err != nil {
		c.Logger().Error(err)
		resp.Success = false
		resp.Message = "Bad Data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Call the open fisca server
	resp, err = openFisca.sendRequest(traceRequest)
	if err != nil {
		c.Logger().Error(err)
		resp.Success = false
		resp.Message = "Can not call OF"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Success = true
	fmt.Print(resp)
	return c.JSON(http.StatusOK, resp)
}

func (of OpenFisca) sendRequest(traceRequest *bindings.TraceRequest) (renderings.TraceResponse, error) {

	//Modify TraceRequest if necessary
	requestBody, err := json.Marshal(traceRequest)
	if err != nil {
		return renderings.TraceResponse{}, err
	}

	//TODO: Put url in a config
	resp, err := http.Post("https://fd7a43f1-b30f-4895-836d-5b52cede5318.mock.pstmn.io/trace","application/json",  bytes.NewBuffer(requestBody))
	if err != nil {
		return renderings.TraceResponse{}, err
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	temp := &renderings.TraceResponse{}
	err = json.NewDecoder(resp.Body).Decode(temp)

	return *temp, err
}
