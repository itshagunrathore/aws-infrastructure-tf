package customerrors

type DsaIsDeployingError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e DsaIsDeployingError) Error() string {
	return e.Message
}
func NewDsaIsDeployingError(msg string) DsaIsDeployingError {
	return DsaIsDeployingError{
		Message: msg,
	}
}
