package dsaservice

import (
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
)

func CreateDsaJob(request dto.CreateDsaJobRequest) {
	host := "tdicam2118dev00.gateway.dev.cloud.teradata.com"
	dsaDriver := dsa.NewDsaDriver(host, "443", "", "", "", "")

	targetGroup, err := dsaDriver.GetTargetGroup()
	userName := "baradmin"
	password := "V#cbjioln0367192"
	dsaRestJobPaylaod := models.RestJobPayload{}

	dsaRestJobPaylaod.RestJobDefinitionModel.JobName = request.JobName
	dsaRestJobPaylaod.RestJobDefinitionModel.JobType = request.JobType
	dsaRestJobPaylaod.RestJobDefinitionModel.JobDescription = request.Description
	dsaRestJobPaylaod.RestJobDefinitionModel.DataDictionaryType = "DATA"
	dsaRestJobPaylaod.RestJobDefinitionModel.SourceSystem = systemName
	dsaRestJobPaylaod.RestJobDefinitionModel.TargetGroupName = targetGroup
	dsaRestJobPaylaod.RestJobDefinitionModel.SrcUserName = userName
	dsaRestJobPaylaod.RestJobDefinitionModel.SrcUserPassword = password
	dsaRestJobPaylaod.RestJobSettingsModel = request.JobSettings
	dsaRestJobPaylaod.RestJobObjectsModels = request.JobObjects

	fmt.Println("creating dsa job")
	dsaDriver.CreateJob(dsaRestJobPaylaod)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(systemName, targetGroup)
	// get system name
	// get target-group-name
	// get password
	// preparepaylaod
	// trigger job_creation
	// update status
}
