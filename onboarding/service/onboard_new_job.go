package service

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/entities"
)

func OnboardNewJob(event models.Event) {
	newJob := entities.Tenant{CustomerName: event.CustomerName}
	repositories.OnboardNewTenant(newJob)
}
