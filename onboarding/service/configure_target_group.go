package service

import (
	"fmt"
)

func ConfigureTGT(event event) (string, error) {
	var payload TgtPayload
	var runtimeBuckets Buckets
	var runtimePrefix PrefixList
	var targetMedia TargetMediaBuckets
	mediaServers, err := GetMedia(event)
	if err != nil {
		return "Failed to get Media server names", err
	}

	for i, media := range mediaServers {
		fmt.Println(i, media)
		runtimePrefix.PrefixName = fmt.Sprintf("m%d/", i+1)
		runtimePrefix.StorageDevices = 1
		runtimeBuckets.BucketName = event.BucketName
		runtimeBuckets.PrefixList = append(runtimeBuckets.PrefixList, runtimePrefix)
		targetMedia.BarMediaServer = media
		targetMedia.Buckets = append(targetMedia.Buckets, runtimeBuckets)
		runtimeBuckets = Buckets{}
		runtimePrefix = PrefixList{}
		payload.TargetMediaBuckets = append(payload.TargetMediaBuckets, targetMedia)
		targetMedia = TargetMediaBuckets{}
	}

	payload.AwsAccountName = event.AwsAccountName
	payload.IsEnabled = true
	payload.Region = event.Region
	payload.TargetGroupName = "TG_BAAS_GO3"

	fmt.Println(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/components/target-groups/s3", event.DscIp, event.Port)
	response, err := PostConfigDsc(url, payload)
	if err != nil {
		return "Failed to configure target group", err
	} else {
		return string(response), err
	}
}
