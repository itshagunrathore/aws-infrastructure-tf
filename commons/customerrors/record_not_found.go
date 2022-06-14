package customerrors

type ResourceNotFound struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e ResourceNotFound) Error() string {
	return e.Message
}

func NewResourceNotFound(msg string) ResourceNotFound {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return ResourceNotFound{
		Message: msg,
	}
}
