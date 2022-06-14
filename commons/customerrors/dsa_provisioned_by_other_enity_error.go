package customerrors

type DsaAlreadyProvisionedByOtherEntityError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e DsaAlreadyProvisionedByOtherEntityError) Error() string {
	return e.Message
}
func NewDsaAlreadyProvisionedByOtherEntityError(msg string) DsaAlreadyProvisionedByOtherEntityError {
	return DsaAlreadyProvisionedByOtherEntityError{
		Message: msg,
	}
}
