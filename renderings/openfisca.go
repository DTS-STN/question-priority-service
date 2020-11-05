package renderings

// NextQuestionResponse is the response returned to the client that contains
// the same information received in the request, with the added property of
// BenefitElegibility, which contains a list of benefits to which the client
// has been deemed eligible or inelegible.
type NextQuestionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	// Date period for request
	RequestDate string `json:"request_date"`
	// Array of life journeys, which represent a subset of benefits
	LifeJourneys []string `json:"life_journeys"`
	// Array of specific benefits to get the questions for
	BenefitList []string `json:"benefit_list"`
	// List of answered priority questions with their answers and the next priority 
	// question with a value of null
	QuestionList []interface{} `json:"client_response"`
	// List of eligible and non-eligible benefits, populated as responses to
	// prioritized questions are recieved
	BenefitElegibility []interface{} `json:"benefit_elegibility"`
}
