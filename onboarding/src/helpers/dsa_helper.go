package helpers

import (
	"fmt"
	"net/http"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/helpers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/dtos"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/entities"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/mappers"
)

const (
	ClientName         string = "baas" //
	Running            string = "running"
	Deploying          string = "deploying"
	Terminating        string = "terminating"
	ProvisionDsaPath   string = "/v1/accounts/{accountId}/dsa"
	DeprovisionDsaPath string = "/v1/accounts/{accountId}/{clientName}/{clientSessionId}/dsa"
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

	if dsaStatusResp.ClientName != ClientName {
		return customerrors.NewDsaAlreadyProvisionedByOtherEntityError(fmt.Sprintf("Dsa is already provisioned by %v", dsaStatusResp.ClientName))
	} else if dsaStatusResp.Status == Running && dsaStatusResp.ClientName == ClientName {
		return customerrors.NewDsaAlreadyProvisionedError("Dsa is already provisioned")
	} else if dsaStatusResp.Status == Deploying && dsaStatusResp.ClientName == ClientName {
		return customerrors.NewDsaIsDeployingError("Dsa is being deployed for this account")
	} else if dsaStatusResp.Status == Terminating && dsaStatusResp.ClientName == ClientName {
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
