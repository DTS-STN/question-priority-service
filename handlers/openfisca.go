package handlers

import (
	"github.com/DTS-STN/question-priority-service/bindings"
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/DTS-STN/question-priority-service/renderings"
	"github.com/labstack/echo/v4"
	"net/http"
)

// NextQuestion
// @Summary Request Prioritized Questions
// @ID next-question
// @Accept  json
// @Produce  json
// @Success 200 {object} renderings.NextQuestionResponse
// @Failure 400 {object} renderings.QPSError
// @Failure 404 {object} renderings.QPSError
// @Failure 500 {object} renderings.QPSError
// @Param NextQuestion body bindings.NextQuestionRequest 1604599804740 "Journey 1" ["Benefit 1"] [models.Question] [models.Benefit]
// @Router /next [post]
func (h *Handler) NextQuestion(c echo.Context) (err error) {
	var nextQuestionResponse = new(renderings.NextQuestionResponse)
	nextQuestionRequest := new(bindings.NextQuestionRequest)

	// Bind the request into our request struct
	if err := c.Bind(nextQuestionRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, nextQuestionResponse)
	}

	///////////////////////////////////
	// TODO: Implement fetching dynamic questions

	nextQuestionResponse.RequestDate = nextQuestionRequest.RequestDate
	nextQuestionResponse.BenefitEligibility = []models.Benefit{{ID: "1", IsEligible: false}, {ID: "2", IsEligible: false}}

	// This is not permanent and will only be used as Phase 0.1 to return hardcoded content
	if len(nextQuestionRequest.QuestionList) == 0 {
		nextQuestionResponse.QuestionList = []models.Question{{ID: "1"}}
	} else if len(nextQuestionRequest.QuestionList) == 1 && nextQuestionRequest.QuestionList[0].ID == "1" {
		// Since theres 1 question answered and its question 1, lets append 2 and return benefit 1 is eligible if answer is true

		// Append next question to end of result list and send it back
		nextQuestionResponse.QuestionList = append(nextQuestionRequest.QuestionList, models.Question{ID: "2"})

		// Only one question answered, lets check the answer
		if nextQuestionRequest.QuestionList[0].Answer == "true" {
			// Find the write benefit and set it to true
			for i, _ := range nextQuestionResponse.BenefitEligibility {
				if nextQuestionResponse.BenefitEligibility[i].ID == "1" {
					nextQuestionResponse.BenefitEligibility[i].IsEligible = true
					break
				}
			}
		}
	} else if len(nextQuestionRequest.QuestionList) == 1 && nextQuestionRequest.QuestionList[0].ID == "2" {
		// Since theres 1 question answered and its question 2, lets append 1 and return benefit 2 is eligible if answer is true

		// Append next question to end of result list and send it back
		nextQuestionResponse.QuestionList = append(nextQuestionRequest.QuestionList, models.Question{ID: "1"})

		// Only one question answered, lets check the answer
		if nextQuestionRequest.QuestionList[0].Answer == "true" {
			// Find the write benefit and set it to true
			for i, _ := range nextQuestionResponse.BenefitEligibility {
				if nextQuestionResponse.BenefitEligibility[i].ID == "2" {
					nextQuestionResponse.BenefitEligibility[i].IsEligible = true
					break
				}
			}
		}
	} else {
		//Since we only have 2 questions, lets just set asked questions back into the response
		nextQuestionResponse.QuestionList = nextQuestionRequest.QuestionList
		// More than one question answered, lets check the benefits
		for i, _ := range nextQuestionResponse.BenefitEligibility {
			// For each Benefit (i), check the answer if true
			if nextQuestionResponse.BenefitEligibility[i].ID == "1" {
				for j, _ := range nextQuestionRequest.QuestionList {
					if nextQuestionRequest.QuestionList[j].ID == "1" && nextQuestionRequest.QuestionList[j].Answer == "true" {
						nextQuestionResponse.BenefitEligibility[i].IsEligible = true
						break
					}
				}
			} else if nextQuestionResponse.BenefitEligibility[i].ID == "2" {
				for j, _ := range nextQuestionRequest.QuestionList {
					if nextQuestionRequest.QuestionList[j].ID == "2" && nextQuestionRequest.QuestionList[j].Answer == "true" {
						nextQuestionResponse.BenefitEligibility[i].IsEligible = true
						break
					}
				}
			}
		}
	}
	///////////////////////////////////
	return c.JSON(http.StatusOK, nextQuestionResponse)
}
