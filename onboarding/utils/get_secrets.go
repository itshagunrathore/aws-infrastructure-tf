package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretData struct {
	SshKey string `json:"mongo_user"`
}

var versionStage string = "AWSCURRENT"

func GetSecret(secretName string, region string, cloudPlatform string) (string, error) {
	var secretKey string
	var err error
	switch cloudPlatform {
	case "AWS":
		secretKey, err = GetAwsSecret(secretName, region)
	case "AZURE":
		secretKey, err = GetAzureSecret(secretName, region)
	case "GCP":
		secretKey, err = GetGcpSecret(secretName, region)
	}
	return secretKey, err
}
func GetAwsSecret(secretName string, region string) (string, error) {
	svc := secretsmanager.New(
		session.New(),
		aws.NewConfig().WithRegion(region),
	)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(versionStage),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return result, err
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
		fmt.Println(secretString)
	}
	return secretString, err
}

func GetAzureSecret(secretName string, region string) (string, error) {
	//TODO add Azure code here
	var err error
	return secretName, err
}

func GetGcpSecret(secretName string, region string) (string, error) {
	//TODO add Gcp code here
	var err error
	return secretName, err
}
