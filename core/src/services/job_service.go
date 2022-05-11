package services

import (
	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
)

type JobService interface {
	GetJob(jobId uint) entities.JobDefinition
	CreateJob(context *gin.Context, accountId string, postJobDto dto.PostJobDto) (int, error)
}

type jobService struct {
	JobDefinitionRepository repositories.JobDefinitionRepository
}

func NewJobService(jd repositories.JobDefinitionRepository) JobService {
	return &jobService{
		JobDefinitionRepository: jd,
	}
}

func (service *jobService) GetJob(jobId uint) entities.JobDefinition {
	return service.JobDefinitionRepository.FindById(jobId)
}

func (service *jobService) CreateJob(context *gin.Context, accountId string, postJobDto dto.PostJobDto) (int, error) {

	jobName := postJobDto.Name

	jobDefinitionEntity := maptoentity(postJobDto)
	return 1, nil
}
