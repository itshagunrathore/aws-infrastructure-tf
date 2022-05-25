package dsahandlers

import (
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dsaclient/dsaservice"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
)

func CreateDsaJobHandler(createDsaJobRequest dto.CreateDsaJobRequest) {
	fmt.Println("request recieved for create job")
	createDsaJobRequest.JobObjects[0].ObjectName = "DBC"
	createDsaJobRequest.JobObjects[0].ParentType = "DATABASE"
	createDsaJobRequest.JobObjects[0].ParentName = ""
	createDsaJobRequest.JobObjects[0].ObjectType = "DATABASE"
	dsaservice.CreateDsaJob(createDsaJobRequest)
}
