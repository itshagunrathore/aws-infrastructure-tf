package tests

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/stretchr/testify/assert"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/secrets"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/secrets/mocks"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"testing"
)

func TestAwsGetSecretsSuccess(t *testing.T) {
	mockConstructor := mocks.MockConstructorTestingTNewCloudSecrets(t)
	mockInterface := mocks.NewGetCloudSecretsInterface(mockConstructor)

	driver := secrets.NewSecretsDriver(mockInterface, "test", "test-us")

	response, err := driver.GetSecret(models.GetSecretPayload{CloudPlatform: models.AWS})
	if err != nil {
		log.Infow(err.Error())
	} else {
		log.Infow(response.Data)
	}

	assert.Equal(t, "Success from AWS", response.Data)
	assert.Error(t, nil, err)
}

func TestAureGetSecrets(t *testing.T) {

	mockConstructor := mocks.MockConstructorTestingTNewCloudSecrets(t)
	mockInterface := mocks.NewGetCloudSecretsInterface(mockConstructor)

	driver := secrets.NewSecretsDriver(mockInterface, "test", "test-us")

	response, err := driver.GetSecret(models.GetSecretPayload{CloudPlatform: models.AZURE})
	if err != nil {
		log.Infow(err.Error())
	} else {
		log.Infow(response.Data)
	}

	assert.Equal(t, "", response.Data)
	assert.Error(t, errors.New("azure Not Supported"), err)

}

func TestGCPGetSecrets(t *testing.T) {
	mockConstructor := mocks.MockConstructorTestingTNewCloudSecrets(t)
	mockInterface := mocks.NewGetCloudSecretsInterface(mockConstructor)

	driver := secrets.NewSecretsDriver(mockInterface, "test", "test-us")

	response, err := driver.GetSecret(models.GetSecretPayload{CloudPlatform: models.GCP})
	if err != nil {
		log.Infow(err.Error())
	} else {
		log.Infow(response.Data)
	}
	assert.Equal(t, "", response.Data)
	assert.Error(t, errors.New("GCP Not Supported"), err)
}
func TestDecodeSecret(t *testing.T) {
	secretString := "IntcImtleVwiOlwidmFsdWVcIn0i"
	var str string
	output := secretsmanager.GetSecretValueOutput{SecretBinary: []byte(secretString)}

	res, err := secrets.DecodeSecret(str, &output)
	print(res)
	if err != nil {
		log.Infow(err.Error())
	}
	assert.Equal(t, "{\"key\":\"value\"}", res)
}
