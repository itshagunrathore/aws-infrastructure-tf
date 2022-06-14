package service

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
)

func DsmainRestart(event models.Event, DsmainStatus *models.DetailedStatus) string {
	response := utils.DsmainRestart(event.DscIp, event.AccountId, event.TPAId, event.CloudPlatform, event.Region)
	if !response {
		log.Error("DSMAIN restart failed")
		DsmainStatus.StepStatus = "Failed"
		return "DSMAIN restart failed"
	} else {
		DsmainStatus.StepStatus = "Success"
		return "DSMAIN restart Success"
	}
}
