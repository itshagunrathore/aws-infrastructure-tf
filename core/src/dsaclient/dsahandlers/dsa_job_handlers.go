package dsahandlers

import (
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dsaclient/dsaservice"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
)

func CreateDsaJobHandler(createDsaJobRequest dto.CreateDsaJobRequest) {
	fmt.Println("request recieved for create job")

	dsaservice.NewDsaService().CreateDsaJob(createDsaJobRequest)
}
