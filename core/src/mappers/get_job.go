package mappers

import (
	"encoding/json"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type GetJobMapper interface {
	ToGetJobDto(jobDefinition entities.JobDefinition) (dto.GetJobDto, error)
}

type getJobMapper struct {
}

func NewGetJobMapper() GetJobMapper {
	return &getJobMapper{}
}

func (g getJobMapper) ToGetJobDto(jobDefinition entities.JobDefinition) (dto.GetJobDto, error) {
	getJobDto := dto.GetJobDto{}

	getJobDto.JobID = jobDefinition.JobId
	getJobDto.Name = jobDefinition.Name
	getJobDto.Description = jobDefinition.Description
	getJobDto.SiteID = jobDefinition.CustomerSite.SiteId
	getJobDto.IsActive = jobDefinition.IsActive
	getJobDto.Priority = jobDefinition.JobPriority
	getJobDto.JobType = jobDefinition.JobType
	getJobDto.SiteTargetType = jobDefinition.CustomerSite.SiteTargetType

	getJobDto.NoOfRetentionCopies = jobDefinition.RetentionCopiesCount
	getJobDto.IsAutoAbortActive = jobDefinition.IsAutoAbortActive
	getJobDto.AutoAbortInMinutes = jobDefinition.AutoAbortInMin
	getJobDto.BackupMechanism = jobDefinition.BackupMechanism
	getJobDto.LastExecutionDetails.Status = jobDefinition.LatestJobSession.LatestStatus

	getJobDto.LastExecutionDetails.BackupSetSize = jobDefinition.LatestJobSession.BackupSetSizeInBytes
	getJobDto.LastExecutionDetails.StartTime = jobDefinition.LatestJobSession.JobStartTime
	getJobDto.LastExecutionDetails.EndTime = jobDefinition.LatestJobSession.JobEndTime

	if jobDefinition.JobObjects != nil {
		var jobObjects []models.JobObjects
		err := json.Unmarshal(jobDefinition.JobObjects, &jobObjects)
		if err != nil {
			return dto.GetJobDto{}, err
		}
		getJobDto.DsaJobDefinition.JobObjects = jobObjects
	}
	//TODO need to check on job_settings and next run time as schedule service is not available yet
	return getJobDto, nil
}
