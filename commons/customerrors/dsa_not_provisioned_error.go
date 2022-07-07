package customerrors

type DsaNotProvisionedError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e DsaNotProvisionedError) Error() string {
	return e.Message
}
func NewDsaNotProvisionedError(msg string) DsaNotProvisionedError {
	return DsaNotProvisionedError{
		Message: msg,
	}
}