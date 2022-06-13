package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
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
	GetDsaStatusHelper(resp *http.Response, accountId string) (dtos.GetDsaStatusDtos, error)
	GetErrorMessage(resp *http.Response) (string, error)
}

type dsaHelper struct {
}

func NewDsaHelper() *dsaHelper {
	return &dsaHelper{}
}

func (h *dsaHelper) CheckDsaStatus(dsaStatusResp dtos.GetDsaStatusDtos) error {
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
func (h *dsaHelper) GetErrorMessage(resp *http.Response) (string, error) {
	var errStruct struct{ Error string }
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, &errStruct)
	return errStruct.Error, nil
}
func (h *dsaHelper) GetDsaStatusHelper(resp *http.Response, accountId string) (dtos.GetDsaStatusDtos, error) {
	var getDsaStatusDto dtos.GetDsaStatusDtos
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return getDsaStatusDto, err
		}
		json.Unmarshal(body, &getDsaStatusDto)
		return getDsaStatusDto, nil
	} else if resp.StatusCode == http.StatusNotFound {
		return getDsaStatusDto, customerrors.NewAccountDoesntExistError(fmt.Sprintf("account %v doesnt exist", accountId))
	} else if resp.StatusCode == http.StatusInternalServerError {
		msg, err := h.GetErrorMessage(resp)
		if err != nil {
			return getDsaStatusDto, err
		}
		if msg == fmt.Sprintf("no dsa resource exists for account id %v", accountId) {
			return getDsaStatusDto, customerrors.NewDsaResourceNotFoundError(fmt.Sprintf("no dsa resource exists for account id %v", accountId))
		} else {
			return getDsaStatusDto, errors.New(msg)
		}
	}
	return getDsaStatusDto, nil
}
func (h *dsaHelper) DeprovisionDsaHelper(resp *http.Response, clientSessionId string) (entities.DsaClientSession, error) {
	var dsaClientSessionEntity entities.DsaClientSession

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		dsaClientSessionEntity = mappers.NewDsaMapper().MapDeprovisionRequestUpdate(clientSessionId)
		return dsaClientSessionEntity, nil
	} else if resp.StatusCode == http.StatusNotFound {
		msg, err := h.GetErrorMessage(resp)
		if err != nil {
			return dsaClientSessionEntity, err
		}
		return dsaClientSessionEntity, customerrors.NewAccountDoesntExistError(msg)
	} else {
		msg, err := h.GetErrorMessage(resp)
		if err != nil {
			return dsaClientSessionEntity, err
		}
		return dsaClientSessionEntity, fmt.Errorf("error deprovisioning dsa: %v", msg)
	}
}
func (h *dsaHelper) ProvisionDsaHelper(resp *http.Response, accountId string, context *gin.Context) (entities.DsaClientSession, error) {
	var provisionDsaResponseDto dtos.ProvisionDsaDtos
	var provisionDsaEntity entities.DsaClientSession
	mappers := mappers.NewDsaMapper()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return provisionDsaEntity, err
	}
	log.Infow(fmt.Sprintf("pod-account-service response, statusCode: %v, body %v", resp.StatusCode, string(body)), "baas-trace-id", context.Value("baas-trace-id"))

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		json.Unmarshal(body, &provisionDsaResponseDto)
		provisionDsaEntity = mappers.MapProvisionDsaResponse(provisionDsaResponseDto, accountId)

		return provisionDsaEntity, nil
	} else if resp.StatusCode == http.StatusMethodNotAllowed {
		return provisionDsaEntity, customerrors.NewDsaAlreadyProvisionedError(fmt.Sprintf("Dsa already provisioned by %v", ClientName))
	} else {
		msg, err := h.GetErrorMessage(resp)
		if err != nil {
			return provisionDsaEntity, err
		}
		return provisionDsaEntity, fmt.Errorf("error provisioning dsa: %v", msg)
	}
}
