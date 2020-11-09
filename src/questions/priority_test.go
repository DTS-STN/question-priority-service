package questions

import (
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNext(t *testing.T) {
	// Request data
	answersObject := []models.Question{}

	// Expected result data
	expectedNextQuestion := []models.Question{{ID: "1"}}

	actual, err := GetNext(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expectedNextQuestion, actual)
	}
}

func TestGetNext_OneAnswer(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "1"}}

	// Expected result data
	expectedNextQuestion := []models.Question{{ID: "1"}, {ID: "2"}}

	actual, err := GetNext(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expectedNextQuestion, actual)
	}
}

func TestGetNext_OneAnswer_Two(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "2"}}

	// Expected result data
	expectedNextQuestion := []models.Question{{ID: "2"}, {ID: "1"}}

	actual, err := GetNext(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expectedNextQuestion, actual)
	}
}

func TestGetNext_OneAnswer_Both(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "1"}, {ID: "2"}}

	// Expected result data
	expectedNextQuestion := []models.Question{{ID: "1"}, {ID: "2"}}

	actual, err := GetNext(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expectedNextQuestion, actual)
	}
}
