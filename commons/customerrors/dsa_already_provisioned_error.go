package customerrors

type DsaAlreadyProvisionedError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e DsaAlreadyProvisionedError) Error() string {
	return e.Message
}
