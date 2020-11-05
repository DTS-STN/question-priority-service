package bindings

import "github.com/DTS-STN/question-priority-service/models"

// NextQuestionRequest is the request sent by the client that contains the information
// required for the QPS to determine the next question.
type NextQuestionRequest struct {
	// Date period for request in ms since epoch
	RequestDate int64 `json:"request_date"`
	// Array of life journeys, which represent a subset of benefits
	LifeJourneys []string `json:"life_journeys"`
	// Array of specific benefits to get the questions for
	BenefitList []string `json:"benefit_list"`
	// List of answered priority questions
	QuestionList []models.Question `json:"client_response"`
}
