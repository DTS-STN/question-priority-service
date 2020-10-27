package renderings

type TraceResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Key string `json:"key"`
}
