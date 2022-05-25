package mappers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type JobDefinitionEntityMapper interface {
	MapToJobDefinitionEntity(postJobDto dto.PostJobDto) entities.JobDefinition
}

type jobDefinitionEntityMapper struct {
}

func NewJobDefinitionEntityMapper() JobDefinitionEntityMapper {
	return &jobDefinitionEntityMapper{}
}

func (j jobDefinitionEntityMapper) MapToJobDefinitionEntity(postJobDto dto.PostJobDto) entities.JobDefinition {
	//TODO implement me
	jobToSave := entities.JobDefinition{}

	jobToSave.Name = postJobDto.Name
	jobToSave.JobType = postJobDto.JobType
	jobToSave.BackupType = postJobDto.BackupType
	jobToSave.BackupMechanism = postJobDto.BackupMechanism
	jobToSave.AutoAbortInMin = postJobDto.AutoAbortInMinutes
	jobToSave.IsActive = postJobDto.IsActive
	// this fields needs to be set correctly
	jobToSave.RetentionSource = "S3"
	if postJobDto.NoOfRetentionCopies == 0 {
		jobToSave.RetentionCopiesCount = 2
	}

	return jobToSave
}
