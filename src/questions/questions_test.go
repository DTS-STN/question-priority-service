package questions

import (
	"bytes"
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func osOpenMock(aString string) (*os.File, error) {
	return os.Open("questions_test.json")
}

func cleanUp() {
	osOpen = os.Open
	questions = nil
}

func TestQuestions(t *testing.T) {
	defer cleanUp()

	// redirect to test data
	osOpen = osOpenMock

	// Expected result data
	expectedResult := []models.Question{{ID: "1", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"1"}}}

	// Actual result data
	actual := Questions()

	// Assertions
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedResult, questions)
}

func TestPrefilledQuestions(t *testing.T) {
	defer cleanUp()

	// prefill test data
	questions = []models.Question{{ID: "2", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"2"}}}

	// redirect to test data
	osOpen = osOpenMock

	// Expected result data
	expectedResult := []models.Question{{ID: "2", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"2"}}}

	// Actual result data
	actual := Questions()

	// Assertions
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedResult, questions)
}

func TestReadFile(t *testing.T) {
	defer cleanUp()

	var buffer bytes.Buffer
	buffer.WriteString("test read data. testing to see if readFile works")

	// expected results
	expected := buffer.Bytes()

	// actual results
	content, err := readFile(&buffer)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, expected, content)
}
