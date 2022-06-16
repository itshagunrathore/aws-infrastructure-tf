package models

type ProvisionDsaModel struct {
	ClientName string `json:"clientName"`
	ImageId    string `json:"imageId"`
}
type ProvisionDsaResponseModel struct {
	ClientName      string `json:"clientName"`
	ClientSessionId string `json:"clientSessionId"`
	Error           string `json:"error,omitempty"`
}
