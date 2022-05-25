package customerrors

type DsaError struct {
	ErrorCode string
	Message   string
	Details   interface{}
}

func (d DsaError) Error() string {
	//TODO implement me
	panic("implement me")
}
