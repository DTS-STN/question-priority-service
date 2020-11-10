package questions

import (
	"github.com/DTS-STN/question-priority-service/models"
)

type QuestionInterface interface {
	Questions() []models.Question
	loadQuestions() ([]models.Question, error)
}

type QuestionServiceStruct struct {
	Filename string
}

var QuestionService QuestionInterface
