package service

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/entities"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/models"
)

func OnboardTenant(event models.Event) {
	newSite := entities.Tenant{CustomerName: event.CustomerName}
	repositories.OnboardTenant(newSite)
}
