package mappers

import (
<<<<<<< HEAD
=======
	"encoding/json"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
>>>>>>> cb617580f7e540b5109a595dbdc81d6aa6c40d39
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type GetJobMapper interface {
<<<<<<< HEAD
	ToGetJobDto(jobDefinition entities.JobDefinition) dto.GetJobDto
=======
	ToGetJobDto(jobDefinition entities.JobDefinition) (dto.GetJobDto, error)
>>>>>>> cb617580f7e540b5109a595dbdc81d6aa6c40d39
}

type getJobMapper struct {
}

func NewGetJobMapper() GetJobMapper {
	return &getJobMapper{}
}

<<<<<<< HEAD
func (g getJobMapper) ToGetJobDto(jobDefinition entities.JobDefinition) dto.GetJobDto {
=======
func (g getJobMapper) ToGetJobDto(jobDefinition entities.JobDefinition) (dto.GetJobDto, error) {
>>>>>>> cb617580f7e540b5109a595dbdc81d6aa6c40d39
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
<<<<<<< HEAD
	getJobDto.LastExecutionDetails.BackupSetSize = jobDefinition.LatestJobSession.BackupSetSizeInBytes

	//lastExecutionStartTime := *jobDefinition.LatestJobSession.JobStartTime
	//lastExecutionEndTime := *jobDefinition.LatestJobSession.JobEndTime
	//getJobDto.LastExecutionDetails.StartTime = lastExecutionStartTime
	//getJobDto.LastExecutionDetails.EndTime = lastExecutionEndTime
	// need to find a way to store a job settings and job objects
	return getJobDto
=======

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
>>>>>>> cb617580f7e540b5109a595dbdc81d6aa6c40d39
}
