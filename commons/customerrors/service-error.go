package customerrors

type ServiceError struct {
	Message string
	Details interface{} `json:"details,omitempty"`
}

func (e ServiceError) Error() string {
	return e.Message
}
