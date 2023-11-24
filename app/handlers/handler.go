package handlers

type HandlerSuccess struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type HandlerError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewHandlerSuccess(payload any) *HandlerSuccess {
	return &HandlerSuccess{
		Type:    "OK",
		Payload: payload,
	}
}

func NewHandlerError(err error) *HandlerError {
	return &HandlerError{
		Type:    "NOK",
		Message: err.Error(),
	}
}
