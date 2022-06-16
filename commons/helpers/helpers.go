package helpers

import (
	"encoding/json"
)

type Helper interface {
}
type helper struct {
}

func NewHelper() *helper {
	return &helper{}
}
func (h *helper) GetErrorMessage(resp []byte) string {
	var errStruct struct{ Error string }
	json.Unmarshal(resp, &errStruct)
	return errStruct.Error
}
