package customerrors

type DsaResourceNotFoundError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e DsaResourceNotFoundError) Error() string {
	return e.Message
}
func NewDsaResourceNotFoundError(msg string) DsaResourceNotFoundError {
	return DsaResourceNotFoundError{
		Message: msg}
}
