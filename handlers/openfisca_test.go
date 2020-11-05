package handlers

import (
	"github.com/DTS-STN/question-priority-service/bindings"
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/DTS-STN/question-priority-service/renderings"
	"github.com/DTS-STN/question-priority-service/src/openfisca"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"net/http"
	"net/http/httptest"
	"strings"
)

type openFiscaMock struct {
	mock.Mock
}

func (m *openFiscaMock) SendRequest(traceRequest *bindings.NextQuestionRequest) (renderings.NextQuestionResponse, error) {
	args := m.Called(traceRequest)
	return args.Get(0).(renderings.NextQuestionResponse), args.Error(1)
}

func TestNextQuestion(t *testing.T) {
	const postJSON = `{"key":"value"}`

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodPost, "/trace", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Create the Request and Response for the Mock
	sendRequestData := &bindings.NextQuestionRequest{}
	sendRequestResult := renderings.NextQuestionResponse{
		RequestDate:  100,
		LifeJourneys: []string{"Journey1", "Journey2"},
		BenefitList:  []string{"Benefit1", "Benefit2"},
		QuestionList: []models.Question{
			{"ID1", "Answer1"},
			{"ID2", "Answer2"},
		},
		BenefitEligibility: []models.Benefit{
			{ID: "Benefit1", IsEligible: true},
			{ID: "Benefit2", IsEligible: false},
		},
	}

	// Create a Mock for the interface
	of := new(openFiscaMock)
	// Add a mock call request
	of.On("SendRequest", sendRequestData).
		Return(sendRequestResult, nil)
	// Set the mock to be used by the code
	openfisca.Service = openfisca.OFInterface(of)

	const expectedResult = `{"request_date":100,"life_journeys":["Journey1","Journey2"],"benefit_list":["Benefit1","Benefit2"],"client_response":[{"id":"ID1","Answer":"Answer1"},{"id":"ID2","Answer":"Answer2"}],"benefit_eligibility":[{"id":"Benefit1","is_eligible":true},{"id":"Benefit2","is_eligible":false}]}`
	// Assertions
	if assert.NoError(t, HandlerService.NextQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expectedResult, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
