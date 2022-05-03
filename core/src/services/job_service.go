package services

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
)

type JobService interface {
	GetJob(jobId uint) entities.JobDefinition
}

type jobService struct {
	JobDefinitionRepository repositories.JobDefinitionRepository
}

func NewJobService(jd repositories.JobDefinitionRepository) JobService {
	return &jobService{
		JobDefinitionRepository: jd,
	}
}

func (repo *jobService) GetJob(jobId uint) entities.JobDefinition {
	return repo.JobDefinitionRepository.FindById(jobId)
}


func HandleService() {
	logger := log.Logger()
	logger.Info("Hello World from service")
	logger.Error("Not able to reach blog. from service")
}

// func GetAllJobsService(site_id string) []dto.GetAllJobsDto {
// 	return dao.GetAllJobs(site_id)
// }
