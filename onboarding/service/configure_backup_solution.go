package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

func ConfigureAWSApp(event models.Event, StatusAWSApp *models.DetailedStatus) (string, error) {
	StatusAWSApp.Step = "ConfigureBackupSolution"
	var payload models.AwsApp
	payload.ConfigAwsRest.AcctName = event.AwsAccountName
	payload.ConfigAwsRest.RoleName = event.RoleName
	mediaServers, err := StatusAWSApp.GetMedia(event)
	if err != nil {
		return "Failed to get Media server names", err
	}
	var configawsappresponse models.ConfigAWSAppResponse
	var runtimePrefix models.PrefixList
	var buckets models.Buckets
	var S3BucketsByRegion models.BucketsByRegion
	buckets.BucketName = event.BucketName
	for i, _ := range mediaServers {
		runtimePrefix.PrefixId = 0
		runtimePrefix.PrefixName = fmt.Sprintf("m%d/", i+1)
		runtimePrefix.StorageDevices = 1
		buckets.PrefixList = append(buckets.PrefixList, runtimePrefix)
		runtimePrefix = models.PrefixList{}
	}
	S3BucketsByRegion.Buckets = append(S3BucketsByRegion.Buckets, buckets)
	S3BucketsByRegion.BucketsViewpoint = true
	S3BucketsByRegion.Region = event.Region

	payload.ConfigAwsRest.BucketsByRegion = append(payload.ConfigAwsRest.BucketsByRegion, S3BucketsByRegion)
	log.Info(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/components/backup-applications/aws-s3", event.DscIp, event.Port)
	response, err := dsa.PostConfigDsc(url, payload, &StatusAWSApp)
	json.Unmarshal(response, &configawsappresponse)
	if err != nil {
		StatusAWSApp.StepResponse = configawsappresponse.Status
		StatusAWSApp.StepStatus = "Failed"
		return "Failed to configure target group", err
	} else {
		StatusAWSApp.StepResponse = configawsappresponse.Status
		StatusAWSApp.StepStatus = "Success"
		return string(response), err
	}
}
func GetAWSApp(event models.Event, StatusGetAWSApp *models.DetailedStatus) ([]byte, error) {
	url := fmt.Sprintf("https://%s:%s/dsa/components/backup-applications/aws-s3", event.DscIp, event.Port)
	StatusGetAWSApp.Step = "ConfigureAWSApp"
	response, err := dsa.GetConfigDsc(url, &StatusGetAWSApp)
	if err != nil {
		log.Info("Back to GetAWSApp function with error")
		StatusGetAWSApp.StepStatus = "Failed"
		StatusGetAWSApp.Error = err
		log.Info("Dsa new response : %v", StatusGetAWSApp)
		return response, err
	} else {
		var getSystemResponse models.GetSystem
		json.Unmarshal(response, &getSystemResponse)
		StatusGetAWSApp.StepStatus = "Success"
		StatusGetAWSApp.StepResponse = getSystemResponse.Status
		StatusGetAWSApp.Error = err
		StatusGetAWSApp.StatusCode = 200
		fmt.Printf("Dsa new response : %v", StatusGetAWSApp)
		data, _ := json.Marshal(getSystemResponse)
		log.Info(string(data))
		return response, err
	}
}
