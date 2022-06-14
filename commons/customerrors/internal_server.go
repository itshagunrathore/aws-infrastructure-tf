package customerrors

type InternalServerError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e InternalServerError) Error() string {
	return e.Message
}

func NewInternalServerError(msg string) InternalServerError {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return InternalServerError{
		Message: msg,
	}
}
