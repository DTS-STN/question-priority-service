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
	sendRequest(request *bindings.NextQuestionRequest) (renderings.NextQuestionResponse, error)
}

type OpenFisca struct {}
var openFisca OpenFiscaInterface

// NextQuestion
// @Summary Request Prioritized Questions
// @ID next-question
// @Accept  json
// @Produce  json
// @Success 200 {object} renderings.NextQuestionResponse
// @Failure 400 {object} renderings.NextQuestionResponse
// @Failure 404 {object} renderings.NextQuestionResponse
// @Failure 500 {object} renderings.NextQuestionResponse
// @Param NextQuestion body bindings.NextQuestionRequest true "value"
// @Router /next [post]
func NextQuestion(c echo.Context) (err error) {
	resp := renderings.NextQuestionResponse{}
	NextQuestionRequest := new(bindings.NextQuestionRequest)

	// Bind the request into our request struct
	if err := c.Bind(NextQuestionRequest); err != nil {
		c.Logger().Error(err)
		resp.Success = false
		resp.Message = "Bad Data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Call the open fisca server
	resp, err = openFisca.sendRequest(NextQuestionRequest)
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

func (of OpenFisca) sendRequest(NextQuestionRequest *bindings.NextQuestionRequest) (renderings.NextQuestionResponse, error) {

	//Modify NextQuestionRequest if necessary
	requestBody, err := json.Marshal(NextQuestionRequest)
	if err != nil {
		return renderings.NextQuestionResponse{}, err
	}

	//TODO: Put url in a config
	resp, err := http.Post("https://fd7a43f1-b30f-4895-836d-5b52cede5318.mock.pstmn.io/trace","application/json",  bytes.NewBuffer(requestBody))
	if err != nil {
		return renderings.NextQuestionResponse{}, err
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	temp := &renderings.NextQuestionResponse{}
	err = json.NewDecoder(resp.Body).Decode(temp)

	return *temp, err
}
