package mappers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type GetJobMapper interface {
	ToGetJobDto(jobDefinition entities.JobDefinition) dto.GetJobDto
}

type getJobMapper struct {
}

func NewGetJobMapper() GetJobMapper {
	return &getJobMapper{}
}

func (g getJobMapper) ToGetJobDto(jobDefinition entities.JobDefinition) dto.GetJobDto {
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

	//lastExecutionStartTime := *jobDefinition.LatestJobSession.JobStartTime
	//lastExecutionEndTime := *jobDefinition.LatestJobSession.JobEndTime
	//getJobDto.LastExecutionDetails.StartTime = lastExecutionStartTime
	//getJobDto.LastExecutionDetails.EndTime = lastExecutionEndTime
	// need to find a way to store a job settings and job objects
	return getJobDto
}
