package handlers

import (
	"github.com/DTS-STN/question-priority-service/bindings"
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
	sendRequestData := &bindings.NextQuestionRequest{Key: "value"}
	sendRequestResult := renderings.NextQuestionResponse{Key: "value"}

	// Create a Mock for the interface
	of := new(openFiscaMock)
	// Add a mock call request
	of.On("SendRequest", sendRequestData).
		Return(sendRequestResult, nil)
	// Set the mock to be used by the code
	openfisca.Service = openfisca.OFInterface(of)

	const expectedResult = `{"success":true,"message":"","key":"value"}`
	// Assertions
	if assert.NoError(t, HandlerService.NextQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expectedResult, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}