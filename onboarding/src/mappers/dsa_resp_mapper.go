package mappers

import (
	"time"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/dtos"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/entities"
)

type DsaRespMapper interface {
	MapProvisionDsaResponse() entities.DsaClientSession
	MapDsaClientSessionGetRequest(accountId string) entities.DsaClientSession
}

type dsaRespMapper struct {
}

func NewDsaMapper() *dsaRespMapper {
	return &dsaRespMapper{}
}

func (m *dsaRespMapper) MapProvisionDsaResponse(provisionDsaResponseDto dtos.ProvisionDsaDtos, accountId string) entities.DsaClientSession {
	var dsaProvisioningDtos entities.DsaClientSession
	dsaProvisioningDtos.ClientSessionId = provisionDsaResponseDto.ClientSessionId
	dsaProvisioningDtos.TimeCreated = time.Now().UTC()
	dsaProvisioningDtos.TimeUpdated = time.Now().UTC()
	dsaProvisioningDtos.AccountId = accountId
	return dsaProvisioningDtos
}
func (m *dsaRespMapper) MapDsaClientSessionGetRequest(accountId string) entities.DsaClientSession {
	var dsaClientSessionEntity entities.DsaClientSession
	dsaClientSessionEntity.AccountId = accountId
	dsaClientSessionEntity.IsDeleted = false
	return dsaClientSessionEntity
}
