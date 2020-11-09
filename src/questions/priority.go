package questions

import (
	"github.com/DTS-STN/question-priority-service/models"
)

// This function returns the list of questions
func GetNext(answers []models.Question) (nextQuestions []models.Question, err error) {
	// This is not permanent and will only be used as Phase 0.1 to return hardcoded content
	if len(answers) == 0 {
		nextQuestions = []models.Question{{ID: "1"}}
	} else if len(answers) == 1 && answers[0].ID == "1" {
		// Since theres 1 question answered and its question 1, lets append 2 and return benefit 1 is eligible if answer is true

		// Append next question to end of result list and send it back
		nextQuestions = append(answers, models.Question{ID: "2"})

	} else if len(answers) == 1 && answers[0].ID == "2" {
		// Since theres 1 question answered and its question 2, lets append 1 and return benefit 2 is eligible if answer is true

		// Append next question to end of result list and send it back
		nextQuestions = append(answers, models.Question{ID: "1"})
	} else {
		//Since we only have 2 questions, lets just set asked questions back into the response
		nextQuestions = answers
	}
	return
}
