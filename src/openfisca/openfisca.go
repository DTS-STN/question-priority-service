package openfisca

import (
	"github.com/DTS-STN/question-priority-service/bindings"
	"github.com/DTS-STN/question-priority-service/renderings"

	"bytes"
	"encoding/json"
	"net/http"
)

// OFInterface mainly used for testing
type OFInterface interface {
	SendRequest(request *bindings.NextQuestionRequest) (renderings.NextQuestionResponse, error)
}

type OFService struct{}

// Service that others can use to interact with OpenFisca functions
var Service OFInterface

func (of OFService) SendRequest(NextQuestionRequest *bindings.NextQuestionRequest) (renderings.NextQuestionResponse, error) {

	//Modify NextQuestionRequest if necessary
	requestBody, err := json.Marshal(NextQuestionRequest)
	if err != nil {
		return renderings.NextQuestionResponse{}, err
	}

	//TODO: Put url in a config
	resp, err := http.Post("https://fd7a43f1-b30f-4895-836d-5b52cede5318.mock.pstmn.io/trace", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return renderings.NextQuestionResponse{}, err
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	temp := &renderings.NextQuestionResponse{}
	err = json.NewDecoder(resp.Body).Decode(temp)

	return *temp, err
}
