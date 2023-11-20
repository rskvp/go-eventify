package handlers

type HandlerResponse struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Payload any    `json:"payload"`
}

func NewHandlerResponse(status int, payload any) *HandlerResponse {
	return &HandlerResponse{
		Type:    "OK",
		Status:  status,
		Payload: payload,
	}
}
