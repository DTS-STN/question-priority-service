package renderings

import (
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/DTS-STN/question-priority-service/src/benefits"
)

// NextQuestionResponse is the response returned to the client that contains
// the same information received in the request, with the added property of
// Benefit Eligibility, which contains a list of benefits to which the client
// has been deemed eligible or inelegible.
type NextQuestionResponse struct {
	// Date period for request in ms since epoch
	RequestDate int64 `json:"request_date"`
	// List of answered priority questions with their answers and the next priority
	// question with a value of null
	QuestionList []models.Question `json:"question_list"`
	// List of eligible and non-eligible benefits, populated as responses to
	// prioritized questions are received
	BenefitEligibility []benefits.Benefit `json:"benefit_eligibility"`
}
