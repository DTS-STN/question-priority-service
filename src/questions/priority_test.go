package questions

import (
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	// Mock Data
	mockData := []models.Question{{ID: "1", Answer: ""}, {ID: "2", Answer: ""}, {ID: "3", Answer: ""}}
	// Create a Mock for the interface
	qsMock := new(QuestionServiceMock)
	// Add a mock call request
	qsMock.On("Questions").
		Return(mockData)
	// Set the mock to be used by the code
	realQuestionService := QuestionService
	QuestionService = QuestionInterface(qsMock)

	return func(t *testing.T) {
		osOpen = os.Open
		questions = nil
		QuestionService = realQuestionService
	}
}

func TestGetNext(t *testing.T) {
	cases := []struct {
		name     string
		answers  []models.Question
		expected []models.Question
	}{
		{"NoAnswers", []models.Question{}, []models.Question{{ID: "1"}}},
		{"OneAnswer", []models.Question{{ID: "1"}}, []models.Question{{ID: "2"}}},
		{"OneAnswerSecond", []models.Question{{ID: "2"}}, []models.Question{{ID: "1"}}},
		{"BothAnswers", []models.Question{{ID: "1"}, {ID: "2"}}, nil},
	}

	// Store old Global to be able to use it
	realQuestionService := QuestionService

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := realQuestionService.GetNext(tc.answers)
			// Assertions
			if assert.NoError(t, err) {
				// Here we need to trim new lines since we are parsing a body that could contain them
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
