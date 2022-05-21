package services

import (
	"fmt"
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

	err := postJobDto.Validate()

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	jobName := postJobDto.Name

	_, err = service.checkJobAlreadyExists(context, accountId, jobName)

	if err != nil {
		return 0, err
	}
	//jobDefinitionEntity := maptoentity(postJobDto)
	// map to entity and trigger async flow for job creation on dsa
	//create a job definition in database and return new job with status as in progress

	//mulitple cases :
	// job lifecycle INPROGRESS, FAILED, SUCCESS.
	// if previous is in progress then return 409 bad request with message saying already in progress.
	// if previous failed then accept the request and make it inprogress.
	// if previous success return 409

	// failure cases :
	// failed to trigger async workflow. nothing will happen
	// triggered async workflow but not able to get dsa up. within specific retries. then also can be retriggered.
	// failed at dsa side then also can be retried with correct input or after rectifying dsa error.

	return 1, nil
}

func (service *jobService) checkJobAlreadyExists(context *gin.Context, accountId string, jobName string) (bool, error) {
	jobDefinition, err := service.JobDefinitionRepository.FindByAccountIdAndJobName(accountId, jobName)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if jobDefinition.Name != "" {
		return true, nil
	}
	return false, nil
}
