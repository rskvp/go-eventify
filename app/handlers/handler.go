package handlers

type HandlerResponse struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Payload any    `json:"payload"`
}

func NewHandlerResponse(status int, payload any) *HandlerResponse {
	return &HandlerResponse{
		Type:    getResponseType(status),
		Status:  status,
		Payload: payload,
	}
}

func getResponseType(status int) string {
	if status >= 200 && status < 300 {
		return "OK"
	}

	return "NOK"
}
