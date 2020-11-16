package questions

import (
	"bytes"
	"errors"
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

type QuestionServiceMock struct {
	mock.Mock
}

func (m *QuestionServiceMock) Questions() []models.Question {
	args := m.Called()
	return args.Get(0).([]models.Question)
}

func (m *QuestionServiceMock) LoadQuestions() ([]models.Question, error) {
	args := m.Called()
	return args.Get(0).([]models.Question), args.Error(1)
}

func (m *QuestionServiceMock) GetNext(answers []models.Question) (nextQuestions []models.Question, err error) {
	args := m.Called()
	return args.Get(0).([]models.Question), args.Error(1)
}

func osOpenMock(path string) (*os.File, error) {
	return os.Open("questions_test.json")
}

// anything that should be run a the end of the unit tests should go here
func setupTests() {
	osOpen = os.Open
	questions = nil
	QuestionService = new(QuestionServiceStruct)
}

func TestQuestions(t *testing.T) {
	setupTests()

	var realQuestionService = QuestionServiceStruct{Filename: ""}

	// Expected result data
	expectedResult := []models.Question{{ID: "1", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"1"}}}

	// Create a Mock for the interface
	qsMock := new(QuestionServiceMock)
	// Add a mock call request
	qsMock.On("LoadQuestions").
		Return(expectedResult, nil)
	// Set the mock to be used by the code
	QuestionService = QuestionInterface(qsMock)

	// redirect to test data
	osOpen = osOpenMock

	// Actual result data
	actual := realQuestionService.Questions()

	// Assertions
	assert.Equal(t, expectedResult, actual)
}

func TestQuestionsNotEqual(t *testing.T) {
	setupTests()

	var realQuestionService = QuestionServiceStruct{Filename: ""}

	// Expected result data
	expectedResult := []models.Question{{ID: "1", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"1"}}}

	// Create a Mock for the interface
	qsMock := new(QuestionServiceMock)
	// Add a mock call request
	qsMock.On("LoadQuestions").
		Return([]models.Question{{ID: "2", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"2"}}}, nil)
	// Set the mock to be used by the code
	QuestionService = QuestionInterface(qsMock)

	// redirect to test data
	osOpen = osOpenMock

	// Actual result data
	actual := realQuestionService.Questions()

	// Assertions
	assert.NotEqual(t, expectedResult, actual)
}

func TestPrefilledQuestions(t *testing.T) {
	setupTests()

	var realQuestionService = QuestionServiceStruct{Filename: ""}

	// Expected result data
	expectedResult := []models.Question{{ID: "2", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"2"}}}

	// prefill test data
	questions = expectedResult

	// Create a Mock for the interface
	qsMock := new(QuestionServiceMock)
	// mock a different result from LoadQuestions, to prove that when questions is populated, LoadQuestions isn't called
	qsMock.On("LoadQuestions").
		Return([]models.Question{{ID: "1", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"1"}}}, nil)
	// Set the mock to be used by the code
	QuestionService = QuestionInterface(qsMock)

	// redirect to test data
	osOpen = osOpenMock

	// Actual result data
	actual := realQuestionService.Questions()

	// Assertions
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedResult, questions)
}

func TestReadFile(t *testing.T) {
	setupTests()

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

func TestLoadQuestions(t *testing.T) {
	setupTests()

	// redirect to test data
	osOpen = osOpenMock

	// Expected result data
	expectedResult := []models.Question{{ID: "1", Description: "are you a resident?", Answer: "", OpenFiscaIds: []string{"1"}}}

	// Actual result data
	actual, err := QuestionService.LoadQuestions()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, actual)
}

// Bug: overriding osOpen is not returning an error when the file is non existent
func TestLoadQuestionsError(t *testing.T) {
	setupTests()

	// redirect to test data
	// want to return an error here
	osOpen = func(path string) (*os.File, error) {
		return &os.File{}, errors.New("missing file")
	}

	// Actual result data
	results, err := QuestionService.LoadQuestions()

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, results)
}
