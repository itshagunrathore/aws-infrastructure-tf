package models

type GetSecretPayload struct {
	CloudPlatform SiteTargetType `json:"cloudPlatform"`
}

type SecretResponse struct {
	Data string `json:"data"`
}
