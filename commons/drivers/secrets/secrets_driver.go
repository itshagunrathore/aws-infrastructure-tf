package secrets

import "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"

type CloudSecretsDriver interface {
	GetSecret(request models.GetSecretPayload) (models.SecretResponse, error)
}
type CloudSecretsInterface interface {
	GetAWSSecret(request *SecretsDriver) (string, error)
}
