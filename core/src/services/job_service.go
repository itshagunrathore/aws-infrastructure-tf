package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
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
		msg := fmt.Sprintf("database error occurred whilte fetching job error:%s for accountID: %s and jobID: %d", err.Error(), accountId, jobId)
		log.Errorw(msg, "baas-trace-id", context.Value("baas-trace-id"))
		if err == gorm.ErrRecordNotFound {
			return dto.GetJobDto{}, customerrors.NewResourceNotFound(fmt.Sprintf("job with job-id:%d not found for account %s", jobId, accountId))
		}
		return dto.GetJobDto{}, err
	}

	getJobDto, err := mappers.NewGetJobMapper().ToGetJobDto(jobDefinitionEntity)
	if err != nil {
		msg := fmt.Sprintf("error occurred while unmarshling job_objects to json error:%v", err.Error())
		log.Errorw(msg, "baas-trace-id", context.Value("baas-trace-id"))
		return dto.GetJobDto{}, err
	}
	return getJobDto, nil
}

func (service *jobService) CreateJob(context *gin.Context, accountId string, postJobDto dto.PostJobDto) (int, error) {

	err := postJobDto.Validate()
	if err != nil {
		log.Errorw(err.Error(), "baas-trace-id", context.Value("baas-trace-id"))
		return 0, err
	}

	// TODO should have type for parentType and objectType from enums
	if postJobDto.BackupMechanism == models.DSA && len(postJobDto.DsaJobDefinition.JobObjects) == 0 {
		log.Infow("job objects are not provided creating default job objects", "baas-trace-id", context.Value("baas-trace-id"))
		defaultJobObjects := []models.JobObjects{
			models.JobObjects{
				ObjectName: "DBC",
				ObjectType: "DATABASE",
				ParentName: "",
				ParentType: "DATABASE",
				IncludeAll: true,
			},
		}
		postJobDto.DsaJobDefinition.JobObjects = defaultJobObjects
	}
	jobName := postJobDto.Name
	isPresent, err := service.checkJobAlreadyExists(accountId, jobName)
	if err != nil {
		msg := fmt.Sprintf("error occurred while checking job already present database error:%s for accountID: %s", err.Error(), accountId)
		log.Errorw(msg, "baas-trace-id", context.Value("baas-trace-id"))
		return 0, err
	}

	if isPresent {
		msg := fmt.Sprintf("job with job-name:%s already exists for account with id %s", jobName, accountId)
		log.Errorw(msg, "baas-trace-id", context.Value("baas-trace-id"))
		return 0, customerrors.NewDuplicateResource(msg)
	}

	customerSite, err := service.CustomerSiteRepository.Get(accountId)
	if err != nil {
		msg := fmt.Sprintf("error occurred while fetching customer site from database error:%s for accountID: %s", err.Error(), accountId)
		log.Errorw(msg, "baas-trace-id", context.Value("baas-trace-id"))
		return 0, err
	}

	jobToSave, err := mappers.NewJobDefinitionEntityMapper().MapToJobDefinitionEntity(postJobDto)
	if err != nil {
		msg := fmt.Sprintf("error occurred while marshling job_objects to json error:%v", err.Error())
		log.Errorw(msg, "baas-trace-id", context.Value("baas-trace-id"))
		return 0, err
	}

	jobToSave.CustomerSiteId = customerSite.CustomerSiteId
	if customerSite.SiteTargetType == models.AWS {
		jobToSave.RetentionSource = models.S3
	} else if customerSite.SiteTargetType == models.AZURE {
		jobToSave.RetentionSource = models.Blob
	} else if customerSite.SiteTargetType == models.GCP {
		jobToSave.RetentionSource = models.CloudStorage
	}

	jobId, err := service.JobDefinitionRepository.Save(jobToSave)

	if err != nil {
		msg := fmt.Sprintf("error occurred while saving job_definition database error:%s for accountID: %s", err.Error(), accountId)
		log.Errorw(msg, "baas-trace-id", context.Value("baas-trace-id"))
		return 0, err
	}

	createDsaJobRequest := mappers.NewCreateDsaJobRequestMapper().MapToCreateDsaJobRequest(postJobDto, accountId, jobId, customerSite.SiteTargetType)
	service.triggerDsaJobCreation(context, createDsaJobRequest)

	log.Infow(fmt.Sprintf("job created successfully with jobId %d", jobId), "baas-trace-id", context.Value("baas-trace-id"))
	return jobId, nil
}

func (service *jobService) checkJobAlreadyExists(accountId string, jobName string) (bool, error) {
	jobDefinition, err := service.JobDefinitionRepository.FindByAccountIdAndJobName(accountId, jobName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	// TODO if the previous job is job creation failed state then we will update the same job need to confirm on this
	if jobDefinition.Name != "" {
		return true, nil
	}
	return false, nil
}

func (service *jobService) triggerDsaJobCreation(context *gin.Context, createDsaJobRequest dto.CreateDsaJobRequest) {
	log.Infow("Triggering dsa job creation in go routine", "baas-trace-id", context.Value("baas-trace-id"))
	go dsahandlers.CreateDsaJobHandler(context, createDsaJobRequest)
}
