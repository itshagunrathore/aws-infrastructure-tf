package customerrors

type ErrorResponse struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return ErrorResponse{
		Message: msg,
	}
}

func NotFound(msg string) ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return ErrorResponse{
		Message: msg,
	}
}

func Unauthorized(msg string) ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return ErrorResponse{
		Message: msg,
	}
}

func Forbidden(msg string) ErrorResponse {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return ErrorResponse{
		Message: msg,
	}
}

func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return ErrorResponse{
		Message: msg,
	}
}

func DuplicateResource(msg string) ErrorResponse {
	if msg == "" {
		msg = "resource already exists"
	}
	return ErrorResponse{
		Message: msg,
	}
}
