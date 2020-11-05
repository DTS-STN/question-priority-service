package handlers

import (
	"github.com/DTS-STN/question-priority-service/bindings"
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/DTS-STN/question-priority-service/renderings"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2/json"
	"testing"

	"net/http"
	"net/http/httptest"
	"strings"
)

func TestNextQuestion_NoQuestions(t *testing.T) {
	// Request data
	nextQuestionRequest := bindings.NextQuestionRequest{
		RequestDate:  100,
		LifeJourneys: []string{"LifeJourney1", "LifeJourney2"},
		BenefitList:  []string{"Benefit1", "Benefit2"},
		QuestionList: []models.Question{},
	}

	request, err := json.Marshal(nextQuestionRequest)
	if err != nil {
		assert.Fail(t, "Could not parse Test Request into JSON")
	}

	// Expected result data
	sendRequestResult := renderings.NextQuestionResponse{
		RequestDate: 100,
		QuestionList: []models.Question{
			{"1", ""},
		},
		BenefitEligibility: []models.Benefit{
			{ID: "1", IsEligible: false},
			{ID: "2", IsEligible: false},
		},
	}

	expectedResult, err := json.Marshal(sendRequestResult)
	if err != nil {
		assert.Fail(t, "Could not parse Test Expected Result into JSON")
	}

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodPost, "/next", strings.NewReader(string(request)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.NextQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, string(expectedResult), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestNextQuestion_QuestionOneFalse(t *testing.T) {
	// Request Data
	nextQuestionRequest := bindings.NextQuestionRequest{
		RequestDate:  100,
		LifeJourneys: []string{"LifeJourney1", "LifeJourney2"},
		BenefitList:  []string{"Benefit1", "Benefit2"},
		QuestionList: []models.Question{
			{ID: "1", Answer: "false"},
		},
	}

	request, err := json.Marshal(nextQuestionRequest)
	if err != nil {
		assert.Fail(t, "Could not parse Test Request into JSON")
	}

	// Expected Result data
	sendRequestResult := renderings.NextQuestionResponse{
		RequestDate: 100,
		QuestionList: []models.Question{
			{"1", "false"},
			{"2", ""},
		},
		BenefitEligibility: []models.Benefit{
			{ID: "1", IsEligible: false},
			{ID: "2", IsEligible: false},
		},
	}

	expectedResult, err := json.Marshal(sendRequestResult)
	if err != nil {
		assert.Fail(t, "Could not parse Test Expected Result into JSON")
	}

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodPost, "/next", strings.NewReader(string(request)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.NextQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, string(expectedResult), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestNextQuestion_QuestionOneTrue(t *testing.T) {
	// Request Data
	nextQuestionRequest := bindings.NextQuestionRequest{
		RequestDate:  100,
		LifeJourneys: []string{"LifeJourney1", "LifeJourney2"},
		BenefitList:  []string{"Benefit1", "Benefit2"},
		QuestionList: []models.Question{
			{ID: "1", Answer: "true"},
		},
	}

	request, err := json.Marshal(nextQuestionRequest)
	if err != nil {
		assert.Fail(t, "Could not parse Test Request into JSON")
	}

	// Expected Result data
	sendRequestResult := renderings.NextQuestionResponse{
		RequestDate: 100,
		QuestionList: []models.Question{
			{"1", "true"},
			{"2", ""},
		},
		BenefitEligibility: []models.Benefit{
			{ID: "1", IsEligible: true},
			{ID: "2", IsEligible: false},
		},
	}

	expectedResult, err := json.Marshal(sendRequestResult)
	if err != nil {
		assert.Fail(t, "Could not parse Test Expected Result into JSON")
	}

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodPost, "/next", strings.NewReader(string(request)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.NextQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, string(expectedResult), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestNextQuestion_TwoQuestionsFalse(t *testing.T) {
	// Request Data
	nextQuestionRequest := bindings.NextQuestionRequest{
		RequestDate:  100,
		LifeJourneys: []string{"LifeJourney1", "LifeJourney2"},
		BenefitList:  []string{"Benefit1", "Benefit2"},
		QuestionList: []models.Question{
			{ID: "1", Answer: "false"},
			{ID: "2", Answer: "false"},
		},
	}

	request, err := json.Marshal(nextQuestionRequest)
	if err != nil {
		assert.Fail(t, "Could not parse Test Request into JSON")
	}

	// Expected Result data
	sendRequestResult := renderings.NextQuestionResponse{
		RequestDate: 100,
		QuestionList: []models.Question{
			{"1", "false"},
			{"2", "false"},
		},
		BenefitEligibility: []models.Benefit{
			{ID: "1", IsEligible: false},
			{ID: "2", IsEligible: false},
		},
	}

	expectedResult, err := json.Marshal(sendRequestResult)
	if err != nil {
		assert.Fail(t, "Could not parse Test Expected Result into JSON")
	}

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodPost, "/next", strings.NewReader(string(request)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.NextQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, string(expectedResult), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestNextQuestion_TwoQuestionsTrue(t *testing.T) {
	// Request Data
	nextQuestionRequest := bindings.NextQuestionRequest{
		RequestDate:  100,
		LifeJourneys: []string{"LifeJourney1", "LifeJourney2"},
		BenefitList:  []string{"Benefit1", "Benefit2"},
		QuestionList: []models.Question{
			{ID: "1", Answer: "true"},
			{ID: "2", Answer: "true"},
		},
	}

	request, err := json.Marshal(nextQuestionRequest)
	if err != nil {
		assert.Fail(t, "Could not parse Test Request into JSON")
	}

	// Expected Result data
	sendRequestResult := renderings.NextQuestionResponse{
		RequestDate: 100,
		QuestionList: []models.Question{
			{"1", "true"},
			{"2", "true"},
		},
		BenefitEligibility: []models.Benefit{
			{ID: "1", IsEligible: true},
			{ID: "2", IsEligible: true},
		},
	}

	expectedResult, err := json.Marshal(sendRequestResult)
	if err != nil {
		assert.Fail(t, "Could not parse Test Expected Result into JSON")
	}

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodPost, "/next", strings.NewReader(string(request)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.NextQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, string(expectedResult), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
