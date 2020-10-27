package handlers

import (
	"github.com/DTS-STN/question-priority-service/bindings"
	"github.com/DTS-STN/question-priority-service/renderings"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"net/http"
	"net/http/httptest"
	"strings"

	"fmt"
)

var (
	postJSON = `{"key":"value"}`
	expectedResult = `{"success":true,"message":"","key":"value"}`
)

type openFiscaMock struct {
	mock.Mock
}

func (m *openFiscaMock) sendRequest(traceRequest *bindings.TraceRequest) (renderings.TraceResponse, error) {
	args := m.Called(traceRequest)
	return args.Get(0).(renderings.TraceResponse), args.Error(1)
}

func TestTrace(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/trace", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	of := new(openFiscaMock)

	sendRequestData := &bindings.TraceRequest{Key: "value"}
	sendRequestResult := renderings.TraceResponse{Key: "value"}

	of.On("sendRequest", sendRequestData).
		Return(sendRequestResult, nil)

	openFisca = OpenFiscaInterface(of)

	// Assertions
	if assert.NoError(t, Trace(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Print(rec.Body.String())
		assert.Equal(t, expectedResult, rec.Body.String())
	}
}
