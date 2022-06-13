package service

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/entities"
)

func OnboardNewTenant(event models.Event) {
	newSite := entities.CustomerSite{CustomerName: event.CustomerName}
	repositories.OnboardNewTenant(newSite)
}
