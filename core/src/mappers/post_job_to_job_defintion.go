package mappers

import (
	"encoding/json"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type JobDefinitionEntityMapper interface {
	MapToJobDefinitionEntity(postJobDto dto.PostJobDto) (entities.JobDefinition, error)
}

type jobDefinitionEntityMapper struct {
}

func NewJobDefinitionEntityMapper() JobDefinitionEntityMapper {
	return &jobDefinitionEntityMapper{}
}

func (j jobDefinitionEntityMapper) MapToJobDefinitionEntity(postJobDto dto.PostJobDto) (entities.JobDefinition, error) {
	//TODO implement me
	jobToSave := entities.JobDefinition{}

	jobToSave.Name = postJobDto.Name
	jobToSave.JobType = postJobDto.JobType
	jobToSave.BackupType = postJobDto.BackupType
	jobToSave.Description = postJobDto.Description
	jobToSave.BackupMechanism = postJobDto.BackupMechanism
	jobToSave.AutoAbortInMin = postJobDto.AutoAbortInMinutes
	jobToSave.RetentionSource = models.S3
	jobToSave.Status = "IN_PROGRESS"
	if postJobDto.IsActive != nil {
		jobToSave.IsActive = *postJobDto.IsActive
	} else {
		jobToSave.IsActive = true
	}
	if postJobDto.NoOfRetentionCopies == 0 {
		jobToSave.RetentionCopiesCount = 2
	}
	if postJobDto.Priority == 0 {
		jobToSave.JobPriority = 5
	} else {
		jobToSave.JobPriority = postJobDto.Priority
	}
	if postJobDto.BackupMechanism == models.DSA {
		jobObjects, err := json.Marshal(postJobDto.DsaJobDefinition.JobObjects)
		if err != nil {
			return entities.JobDefinition{}, err
		}
		jobToSave.JobObjects = jobObjects
	}

	return jobToSave, nil
}
