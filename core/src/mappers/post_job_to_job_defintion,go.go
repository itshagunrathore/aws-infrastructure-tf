package mappers

import "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"

type JobDefinitionEntityMapper interface {
	MapToJobDefinitionEntity(postJobDto dto.PostJobDto)
}

type jobDefinitionEntityMapper struct {
}

func NewJobDefintionEntityMapper() JobDefinitionEntityMapper {
	return &jobDefinitionEntityMapper{}
}

func (j jobDefinitionEntityMapper) MapToJobDefinitionEntity(postJobDto dto.PostJobDto) {
	//TODO implement me
}
