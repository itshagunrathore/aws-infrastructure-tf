package handlers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/config"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/services"
)

type JobHandlers interface {
	GetJob() error
}

GetAllJob() error
type jobHandlers struct {
	JobService services.JobService
}

func (service jobHandlers) GetJob() error {
	services.JobService.GetJob(dto)
}

func HandleLogging() {
	logger := log.Logger()
	logger.Info("Hello World")
	logger.Error("Not able to reach blog.")

	var configs models.Configurations
	config.GetConfig(&configs)

	logger.Info("the configs are ", configs)
}
