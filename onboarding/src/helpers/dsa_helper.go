package helpers

import (
	"fmt"
	"net/http"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/constants"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/helpers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/dtos"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/entities"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/mappers"
)

type DsaHelper interface {
	CheckDsaStatus(dsaStatusResp dtos.GetDsaStatusDtos) error
	DeprovisionDsaHelper(resp []byte, statusCode int, clientSessionId string) (entities.DsaClientSession, error)
}

type dsaHelper struct {
}

func NewDsaHelper() *dsaHelper {
	return &dsaHelper{}
}

func (h *dsaHelper) CheckDsaStatus(dsaStatusResp models.DscInstanceDetails) error {
	// we are returning nil cause DSA is not provisoned & we got an error from get dsa status
	// in this case we need to provision dsa
	if dsaStatusResp.ClientName == "" {
		return nil
	}

	if dsaStatusResp.ClientName != constants.ClientName {
		return customerrors.NewDsaAlreadyProvisionedByOtherEntityError(fmt.Sprintf("Dsa is already provisioned by %v", dsaStatusResp.ClientName))
	} else if dsaStatusResp.Status == constants.Running && dsaStatusResp.ClientName == constants.ClientName {
		return customerrors.NewDsaAlreadyProvisionedError("Dsa is already provisioned")
	} else if dsaStatusResp.Status == constants.Deploying && dsaStatusResp.ClientName == constants.ClientName {
		return customerrors.NewDsaIsDeployingError("Dsa is being deployed for this account")
	} else if dsaStatusResp.Status == constants.Terminating && dsaStatusResp.ClientName == constants.ClientName {
		return customerrors.NewDsaAlreadyProvisionedError("Dsa is getting terminated")
	} else {
		return nil
	}
}
func (h *dsaHelper) HandleDeprovisioningResponse(resp []byte, statusCode int, clientSessionId string) (entities.DsaClientSession, error) {
	var dsaClientSessionEntity entities.DsaClientSession
	helper := helpers.NewHelper()
	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		dsaClientSessionEntity = mappers.NewDsaMapper().MapDeprovisionRequestUpdate(clientSessionId)
		return dsaClientSessionEntity, nil
	} else if statusCode == http.StatusNotFound {
		msg := helper.GetErrorMessage(resp)
		return dsaClientSessionEntity, customerrors.NewAccountDoesntExistError(msg)
	} else {
		msg := helper.GetErrorMessage(resp)
		return dsaClientSessionEntity, fmt.Errorf("error deprovisioning dsa: %v", msg)
	}
}
