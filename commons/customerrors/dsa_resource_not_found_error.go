package customerrors

type DsaResourceNotFoundError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e DsaResourceNotFoundError) Error() string {
	return e.Message
}
