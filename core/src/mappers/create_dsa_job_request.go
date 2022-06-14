package mappers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
)

type CreateDsaJobRequestMapper interface {
	MapToCreateDsaJobRequest(postJobDto dto.PostJobDto, accountId string, jobId int, siteTargetType models.SiteTargetType) dto.CreateDsaJobRequest
}

type createDsaJobRequestMapper struct {
}

func NewCreateDsaJobRequestMapper() CreateDsaJobRequestMapper {
	return &createDsaJobRequestMapper{}
}

func (c createDsaJobRequestMapper) MapToCreateDsaJobRequest(postJobDto dto.PostJobDto, accountId string, jobId int, siteTargetType models.SiteTargetType) dto.CreateDsaJobRequest {
	createDsaJobRequest := dto.CreateDsaJobRequest{}

	createDsaJobRequest.JobName = postJobDto.Name
	createDsaJobRequest.Description = postJobDto.Description
	createDsaJobRequest.JobType = postJobDto.JobType
	createDsaJobRequest.JobSettings = postJobDto.DsaJobDefinition.JobSettings
	createDsaJobRequest.JobObjects = postJobDto.DsaJobDefinition.JobObjects
	createDsaJobRequest.BaaSJobId = jobId
	createDsaJobRequest.AccountId = accountId
	createDsaJobRequest.SiteTargetType = siteTargetType
	return createDsaJobRequest
}
