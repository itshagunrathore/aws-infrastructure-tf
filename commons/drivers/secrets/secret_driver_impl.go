package secrets

import (
	"errors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type SecretsDriver struct {
	GetCloudSecretObject CloudSecretsInterface
	SecretName           string
	Region               string
}

func NewSecretsDriver(getCloudSecretsObject CloudSecretsInterface, secretName string, region string) CloudSecretsDriver {
	return &SecretsDriver{
		GetCloudSecretObject: getCloudSecretsObject,
		SecretName:           secretName,
		Region:               region,
	}
}

func (s *SecretsDriver) GetSecret(request models.GetSecretPayload) (models.SecretResponse, error) {
	secrets := models.SecretResponse{}
	if request.CloudPlatform == models.AWS {
		secretValue, err := s.GetCloudSecretObject.GetAWSSecret(s)
		secrets.Data = secretValue
		if err != nil {
			return models.SecretResponse{}, err
		}
	} else if request.CloudPlatform == models.AZURE {
		return models.SecretResponse{}, errors.New("azure Not Supported")
	} else if request.CloudPlatform == models.GCP {
		return models.SecretResponse{}, errors.New("GCP Not Supported")
	}
	return secrets, nil
}
