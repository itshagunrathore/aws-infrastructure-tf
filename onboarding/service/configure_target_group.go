package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

func ConfigureTGT(event models.Event, StatusConfigTGT *models.DetailedStatus) (string, error) {
	var response string
	var err error
	switch event.CloudPlatform {
	case "AWS":
		response, err = ConfigureAwsTGT(event, StatusConfigTGT)
	case "AZURE":
		response, err = ConfigureAzureTGT(event, StatusConfigTGT)
	case "GCP":
		response, err = ConfigureGcpTGT(event, StatusConfigTGT)
	}
	return response, err

}

func ConfigureAwsTGT(event models.Event, StatusConfigTGT *models.DetailedStatus) (string, error) {
	StatusConfigTGT.Step = "ConfigureTargetGroup"
	var payload models.TgtPayload
	var runtimeBuckets models.Buckets
	var runtimePrefix models.PrefixList
	var targetMedia models.TargetMediaBuckets
	var configtgtresponse models.ConfigTGTResponse
	mediaServers, err := dsa.GetMedia(event, &StatusConfigTGT)
	if err != nil {
		return mediaServers[len(mediaServers)-1], err
	}

	for i, media := range mediaServers {
		runtimePrefix.PrefixName = fmt.Sprintf("m%d/", i+1)
		runtimePrefix.StorageDevices = 1
		runtimeBuckets.BucketName = event.BucketName
		runtimeBuckets.PrefixList = append(runtimeBuckets.PrefixList, runtimePrefix)
		targetMedia.BarMediaServer = media
		targetMedia.Buckets = append(targetMedia.Buckets, runtimeBuckets)
		runtimeBuckets = models.Buckets{}
		runtimePrefix = models.PrefixList{}
		payload.TargetMediaBuckets = append(payload.TargetMediaBuckets, targetMedia)
		targetMedia = models.TargetMediaBuckets{}
	}

	payload.AwsAccountName = event.AwsAccountName
	payload.IsEnabled = true
	payload.Region = event.Region
	payload.TargetGroupName = "TG_BAAS"

	log.Info(payload)
	url := fmt.Sprintf("https://%s:%s%s", event.DscIp, event.Port, models.ConfigAwsTGT)
	response, err := dsa.PostConfigDsc(url, payload, &StatusConfigTGT)
	json.Unmarshal(response, &configtgtresponse)
	if err != nil {
		StatusConfigTGT.StepStatus = "Failed"
		StatusConfigTGT.StepResponse = configtgtresponse.Status
		return configtgtresponse.Status, err
	} else {
		log.Info("Configure targetgroup response: " + response)
		StatusConfigTGT.StepStatus = "Success"
		StatusConfigTGT.StepResponse = configtgtresponse.Status
		StatusConfigTGT.Error = err
		return configtgtresponse.Status, err
	}
}

func ConfigureAzureTGT(event models.Event, StatusConfigTGT *models.DetailedStatus) (string, error) {}
func ConfigureGcpTGT(event models.Event, StatusConfigTGT *models.DetailedStatus) (string, error)   {}
