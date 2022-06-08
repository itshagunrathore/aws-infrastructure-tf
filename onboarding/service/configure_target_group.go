package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

func ConfigureTGT(event models.Event, StatusConfigTGT *models.DetailedStatus) (string, error) {
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
		fmt.Println(i, media)
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
	payload.TargetGroupName = "TG_BAAS_GO3"

	fmt.Println(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/components/target-groups/s3", event.DscIp, event.Port)
	response, err := dsa.PostConfigDsc(url, payload, &StatusConfigTGT)
	json.Unmarshal(response, &configtgtresponse)
	if err != nil {
		StatusConfigTGT.StepStatus = "Failed"
		StatusConfigTGT.StepResponse = configtgtresponse.Status
		return configtgtresponse.Status, err
	} else {
		fmt.Println(response)
		StatusConfigTGT.StepStatus = "Success"
		StatusConfigTGT.StepResponse = configtgtresponse.Status
		StatusConfigTGT.Error = err
		return configtgtresponse.Status, err
	}
}
