package secrets

import (
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
)

type CloudSecrets struct {
}

func NewGetCloudSecrets() CloudSecretsInterface {
	return &CloudSecrets{}
}

func (cs *CloudSecrets) GetAWSSecret(request *SecretsDriver) (string, error) {
	//secretName := "tc268cb7-baas-reader"
	//region := "us-west-2"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(request.Region))

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(request.SecretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.Infow("error path", err.Error())
			log.Error(aerr.Code(), aerr.Error())
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Infow(err.Error())
		}
		return "", err
	}

	//Decrypts secret using the associated KMS CMK.
	//Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secrets string
	if result.SecretString != nil {
		secrets = *result.SecretString
	} else {
		secrets, err := DecodeSecret(secrets, result)
		if err != nil {
			return secrets, err
		}
	}
	return secrets, nil
}
func DecodeSecret(secrets string, result *secretsmanager.GetSecretValueOutput) (string, error) {
	decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
	length, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
	if err != nil {
		log.Infow("Base64 Decode Error:", err)
		return secrets, err
	}
	decodedBinarySecret := string(decodedBinarySecretBytes[:length])
	err = json.Unmarshal([]byte(decodedBinarySecret), &secrets)
	return secrets, err
}
