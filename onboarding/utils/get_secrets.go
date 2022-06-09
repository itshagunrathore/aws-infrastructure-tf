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

var (
	secretName   string = "pod-tenant-%s-%s_sshkey"
	region       string = "us-west-2"
	versionStage string = "AWSCURRENT"
)

func GetSecret(tenantId string, TPASystemId string, cloudPlatform string) string {
	var secretKey string
	switch cloudPlatform {
	case "AWS":
		secretKey = GetAwsSecret(tenantId, TPASystemId)
	case "AZURE":
		secretKey = GetAzureSecret(tenantId, TPASystemId)
	case "GCP":
		secretKey = GetGcpSecret(tenantId, TPASystemId)
	}
	return secretKey
}
func GetAwsSecret(tenantId string, TPASystemId string) string {
	svc := secretsmanager.New(
		session.New(),
		aws.NewConfig().WithRegion(region),
	)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(fmt.Sprintf(secretName, tenantId, TPASystemId)),
		VersionStage: aws.String(versionStage),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		panic(err.Error())
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
		fmt.Println(secretString)
	}
	return secretString
}

func GetAzureSecret(tenantId string, TPASystemId string) string {
	//TODO add Azure code here
}

func GetGcpSecret(tenantId string, TPASystemId string) string {
	//TODO add Gcp code here
}
