package customerrors

type DuplicateResourceFound struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e DuplicateResourceFound) Error() string {
	return e.Message
}
func NewDuplicateResource(msg string) DuplicateResourceFound {
	if msg == "" {
		msg = "resource already exists"
	}
	return DuplicateResourceFound{
		Message: msg,
	}
}
