package handlers

import (
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/DTS-STN/question-priority-service/src/questions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"net/http"
	"net/http/httptest"
	"strings"
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

func TestNextQuestion_NoQuestions(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	// Set the mock to be used by the code
	RealQuestionService := questions.QuestionService
	defer func() {
		questions.QuestionService = RealQuestionService
	}()

	cases := []struct {
		name     string
		answers  string
		expected string
		mock     []models.Question
	}{
		{
			name:     "NoQuestions",
			answers:  `{"request_date":100,"life_journeys":["LifeJourney1","LifeJourney2"],"benefit_list":["Benefit1","Benefit2"],"question_list":[]}`,
			expected: `{"request_date":100,"question_list":[{"id":"1","answer":"","description":"","openfiscaids":null}],"benefit_eligibility":[{"id":"1","eligible":false},{"id":"2","eligible":false}]}`,
			mock:     []models.Question{{ID: "1", Answer: ""}},
		},
		{
			name:     "QuestionOneFalse",
			answers:  `{"request_date":100,"life_journeys":["LifeJourney1","LifeJourney2"],"benefit_list":["Benefit1","Benefit2"],"question_list":[{"id":"1","answer":"false"}]}`,
			expected: `{"request_date":100,"question_list":[{"id":"1","answer":"false","description":"","openfiscaids":null},{"id":"2","answer":"","description":"","openfiscaids":null}],"benefit_eligibility":[{"id":"1","eligible":false},{"id":"2","eligible":false}]}`,
			mock:     []models.Question{{ID: "1", Answer: "false"}, {ID: "2"}},
		},
		{
			name:     "QuestionOneTrue",
			answers:  `{"request_date":100,"life_journeys":["LifeJourney1","LifeJourney2"],"benefit_list":["Benefit1","Benefit2"],"question_list":[{"id":"1","answer":"true"}]}`,
			expected: `{"request_date":100,"question_list":[{"id":"1","answer":"true","description":"","openfiscaids":null},{"id":"2","answer":"","description":"","openfiscaids":null}],"benefit_eligibility":[{"id":"1","eligible":true},{"id":"2","eligible":false}]}`,
			mock:     []models.Question{{ID: "1", Answer: "true"}, {ID: "2"}},
		},
		{
			name:     "TwoQuestionsFalse",
			answers:  `{"request_date":100,"life_journeys":["LifeJourney1","LifeJourney2"],"benefit_list":["Benefit1","Benefit2"],"question_list":[{"id":"1","answer":"false"},{"id":"2","answer":"false"}]}`,
			expected: `{"request_date":100,"question_list":[{"id":"1","answer":"false","description":"","openfiscaids":null},{"id":"2","answer":"false","description":"","openfiscaids":null}],"benefit_eligibility":[{"id":"1","eligible":false},{"id":"2","eligible":false}]}`,
			mock:     []models.Question{{ID: "1", Answer: "false"}, {ID: "2", Answer: "false"}},
		},
		{
			name:     "TwoQuestionsTrue",
			answers:  `{"request_date":100,"life_journeys":["LifeJourney1","LifeJourney2"],"benefit_list":["Benefit1","Benefit2"],"question_list":[{"id":"1","answer":"true"},{"id":"2","answer":"true"}]}`,
			expected: `{"request_date":100,"question_list":[{"id":"1","answer":"true","description":"","openfiscaids":null},{"id":"2","answer":"true","description":"","openfiscaids":null}],"benefit_eligibility":[{"id":"1","eligible":true},{"id":"2","eligible":true}]}`,
			mock:     []models.Question{{ID: "1", Answer: "true"}, {ID: "2", Answer: "true"}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup http request using httptest
			req := httptest.NewRequest(http.MethodPost, "/next", strings.NewReader(tc.answers))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			// Create a httptest record
			rec := httptest.NewRecorder()
			// Create a new Echo Context
			c := e.NewContext(req, rec)

			// Create a Mock for the interface
			qsMock := new(QuestionServiceMock)
			// Add a mock call request
			qsMock.On("GetNext").
				Return(tc.mock, nil)
			questions.QuestionService = questions.QuestionInterface(qsMock)

			// Assertions
			if assert.NoError(t, HandlerService.NextQuestion(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				// Here we need to trim new lines since we are parsing a body that could contain them
				assert.Equal(t, tc.expected, strings.TrimSuffix(rec.Body.String(), "\n"))
			}
		})
	}

}
