package dtos

type ProvisionDsaDtos struct {
	ClientName      string `json:"clientName"`
	ClientSessionId string `json:"clientSessionId"`
}
type GetDsaStatusDtos struct {
	ClientName   string `json:"clientName"`
	Status       string `json:"status"`
	DsaPrivateIp string `json:"dsaPrivateIp,omitempty"`
	DsaPublicIp  string `json:"dsaPublicIp,omitempty"`
	Error        string `json:"error,omitempty"`
}
