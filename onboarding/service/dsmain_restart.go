package service

import (
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
)

func DsmainRestart(event models.Event, DsmainStatus *models.DetailedStatus) string {
	response := utils.DsmainRestart(event.dsaIp, event.tenantId, event.TPASystemId, event.cloudPlatform, event.Region)()
	if !response {
		fmt.Printf("DSMAIN restart failed")
		DsmainStatus.StepStatus = "Failed"
		return "DSMAIN restart failed"
	} else {
		DsmainStatus.StepStatus = "Success"
		return "DSMAIN restart Success"
	}
}
