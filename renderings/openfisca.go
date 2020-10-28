package renderings

type NextQuestionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Key string `json:"key"`
}
