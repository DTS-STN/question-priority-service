package questions

import (
	"github.com/DTS-STN/question-priority-service/models"
)

type QuestionInterface interface {
	Questions() []models.Question
	LoadQuestions() ([]models.Question, error)
	GetNext(answers []models.Question) (nextQuestions []models.Question, err error)
}

type QuestionServiceStruct struct {
	Filename string
}

var QuestionService QuestionInterface
