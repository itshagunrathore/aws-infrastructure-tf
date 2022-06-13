package customerrors

type AccountDoesntExistError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e AccountDoesntExistError) Error() string {
	return e.Message
}
