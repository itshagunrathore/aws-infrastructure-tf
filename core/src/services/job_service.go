package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dsaclient/dsahandlers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/mappers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
	"gorm.io/gorm"
)

type JobService interface {
	GetJob(context *gin.Context, accountId string, jobId int) (dto.GetJobDto, error)
	CreateJob(context *gin.Context, accountId string, postJobDto dto.PostJobDto) (int, error)
}

type jobService struct {
	JobDefinitionRepository    repositories.JobDefinitionRepository
	CustomerSiteRepository     repositories.CustomerSiteRepository
	LatestJobSessionRepository repositories.LatestJobSessionRepository
}

func NewJobService(jd repositories.JobDefinitionRepository, customerSite repositories.CustomerSiteRepository, latestJobSession repositories.LatestJobSessionRepository) JobService {
	return &jobService{
		JobDefinitionRepository:    jd,
		CustomerSiteRepository:     customerSite,
		LatestJobSessionRepository: latestJobSession,
	}
}

func (service *jobService) GetJob(context *gin.Context, accountId string, jobId int) (dto.GetJobDto, error) {
	jobDefinitionEntity, err := service.JobDefinitionRepository.FindByAccountIdAndJobId(accountId, jobId)
	if err != nil {
		return dto.GetJobDto{}, err
	}
	getJobDto := mappers.NewGetJobMapper().ToGetJobDto(jobDefinitionEntity)
	return getJobDto, nil
}

func (service *jobService) CreateJob(context *gin.Context, accountId string, postJobDto dto.PostJobDto) (int, error) {

	err := postJobDto.Validate()
	// TODO should have type for parentType and objectType
	if postJobDto.BackupMechanism == models.DSA && postJobDto.DsaJobDefinition.JobObjects[0].ObjectName == "" {
		postJobDto.DsaJobDefinition.JobObjects[0].ObjectName = "DBC"
		postJobDto.DsaJobDefinition.JobObjects[0].ParentType = "DATABASE"
		postJobDto.DsaJobDefinition.JobObjects[0].ParentName = ""
		postJobDto.DsaJobDefinition.JobObjects[0].ObjectType = "DATABASE"
	}

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	jobName := postJobDto.Name

	isPresent, err := service.checkJobAlreadyExists(context, accountId, jobName)

	if err != nil {
		return 0, err
	}

	if isPresent {
		return 0, customerrors.JobAlreadyExistsError{JobName: jobName, AccountId: accountId}
	}

	customerSite, err := service.CustomerSiteRepository.Get(accountId)

	if err != nil {
		return 0, err
	}
	//fmt.Println(customerSite)
	jobToSave := mappers.NewJobDefinitionEntityMapper().MapToJobDefinitionEntity(postJobDto)
	jobToSave.CustomerSiteId = customerSite.CustomerSiteId
	jobId, err := service.JobDefinitionRepository.Save(jobToSave)
	createDsaJobRequest := mappers.NewCreateDsaJobRequestMapper().MapToCreateDsaJobRequest(postJobDto, accountId, 0)
	defer service.triggerDsaJobCreation(createDsaJobRequest)
	if err != nil {
		return 0, err
	}
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

	return jobId, nil
}

func (service *jobService) checkJobAlreadyExists(context *gin.Context, accountId string, jobName string) (bool, error) {
	jobDefinition, err := service.JobDefinitionRepository.FindByAccountIdAndJobName(accountId, jobName)
	fmt.Println(jobDefinition)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	if jobDefinition.Name != "" {
		return true, nil
	}
	return false, nil
}

func (service *jobService) triggerDsaJobCreation(createDsaJobRequest dto.CreateDsaJobRequest) {

	go dsahandlers.CreateDsaJobHandler(createDsaJobRequest) //even api
	return
}
