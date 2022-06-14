package dsahandlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dsaclient/dsaservice"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
)

func CreateDsaJobHandler(context *gin.Context, createDsaJobRequest dto.CreateDsaJobRequest) {
	log.Infow("request received for creating dsa job", "baas-trace-id", context.Value("baas-trace-id"))

	dsaservice.NewDsaService().CreateDsaJob(context, createDsaJobRequest)
}
