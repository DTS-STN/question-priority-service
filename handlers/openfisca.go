package handlers

import (
	"github.com/DTS-STN/question-priority-service/bindings"
	"github.com/DTS-STN/question-priority-service/renderings"
	"github.com/DTS-STN/question-priority-service/src/openfisca"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
func (h *Handler) NextQuestion(c echo.Context) (err error) {
	resp := renderings.NextQuestionResponse{}
	NextQuestionRequest := new(bindings.NextQuestionRequest)

	// Bind the request into our request struct
	if err := c.Bind(NextQuestionRequest); err != nil {
		c.Logger().Error(err)
		resp.Success = false
		resp.Message = "Bad Data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Call the open fisca service
	resp, err = openfisca.Service.SendRequest(NextQuestionRequest)
	if err != nil {
		c.Logger().Error(err)
		resp.Success = false
		resp.Message = "Can not call OF"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Success = true
	return c.JSON(http.StatusOK, resp)
}
