package service

import (
	"encoding/json"
	"fmt"
)

func ConfigureAWSApp(event event) (string, error) {
	var payload AwsApp
	payload.ConfigAwsRest.AcctName = event.AwsAccountName
	payload.ConfigAwsRest.RoleName = event.RoleName
	mediaServers, err := GetMedia(event)
	if err != nil {
		return "Failed to get Media server names", err
	}
	var runtimePrefix PrefixList
	var buckets Buckets
	var S3BucketsByRegion BucketsByRegion
	buckets.BucketName = event.BucketName
	for i, _ := range mediaServers {
		runtimePrefix.PrefixId = 0
		runtimePrefix.PrefixName = fmt.Sprintf("m%d/", i+1)
		runtimePrefix.StorageDevices = 1
		buckets.PrefixList = append(buckets.PrefixList, runtimePrefix)
		runtimePrefix = PrefixList{}
	}
	S3BucketsByRegion.Buckets = append(S3BucketsByRegion.Buckets, buckets)
	S3BucketsByRegion.BucketsViewpoint = true
	S3BucketsByRegion.Region = event.Region

	payload.ConfigAwsRest.BucketsByRegion = append(payload.ConfigAwsRest.BucketsByRegion, S3BucketsByRegion)
	fmt.Println(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/components/backup-applications/aws-s3", event.DscIp, event.Port)
	response, err := PostConfigDsc(url, payload)
	if err != nil {
		return "Failed to configure target group", err
	} else {
		return string(response), err
	}
}
func GetAWSApp(event event) (string, error) {
	url := fmt.Sprintf("https://%s:%s/dsa/components/backup-applications/aws-s3", event.DscIp, event.Port)
	response, err := GetConfigDsc(url, event.DscIp, event.Port)
	if err != nil {
		return "Failed to configure target group", err
	} else {
		var jresponse GetSystem
		json.Unmarshal(response, &jresponse)
		return jresponse.Status, err
	}
}
