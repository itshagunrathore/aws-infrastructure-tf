package customerrors

type BadRequest struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e BadRequest) Error() string {
	return e.Message
}
func NewBadRequest(msg string) BadRequest {
	if msg == "" {
		msg = "resource already exists"
	}
	return BadRequest{
		Message: msg,
	}
}
