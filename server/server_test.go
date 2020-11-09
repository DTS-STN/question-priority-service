package server

import (
	"github.com/DTS-STN/question-priority-service/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"testing"
)

type HandlerServiceMock struct {
	mock.Mock
}

func (m *HandlerServiceMock) HealthCheck(c echo.Context) error {
	args := m.Called()
	return args.Error(1)
}

func (m *HandlerServiceMock) NextQuestion(c echo.Context) error {
	args := m.Called()
	return args.Error(1)
}

// TODO: This doesn't work, need to setup an http client and call the endpoints to run tests
func TestServer(t *testing.T) {
	e := echo.New()

	// Create a Mock for the interface
	handlerMock := new(HandlerServiceMock)
	// Add a mock call request
	handlerMock.On("HealthCheck", e).
		Return(nil)
	handlerMock.On("NextQuestion", e).
		Return(nil)
	// Set the mock to be used by the code
	handlers.HandlerService = handlers.HandlerServiceInterface(handlerMock)

	// TODO: Add Tests

}
