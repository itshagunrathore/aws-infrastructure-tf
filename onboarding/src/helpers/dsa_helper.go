package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/dtos"
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
	if dsaStatusResp.ClientName != ClientName {
		return customerrors.DsaAlreadyProvisionedByOtherEntityError{Message: fmt.Sprintf("Dsa is already provisioned by %v", dsaStatusResp.ClientName)}
	} else if dsaStatusResp.Status == Running && dsaStatusResp.ClientName == ClientName {
		return customerrors.DsaAlreadyProvisionedError{Message: "Dsa is already provisioned"}
	} else if dsaStatusResp.Status == Deploying && dsaStatusResp.ClientName == ClientName {
		return customerrors.DsaIsDeployingError{Message: "Dsa is being deployed for this account"}
	} else if dsaStatusResp.Status == Terminating && dsaStatusResp.ClientName == ClientName {
		return customerrors.DsaAlreadyProvisionedError{Message: "Dsa is getting terminated"}
	} else {
		return nil
	}
}
func (h *dsaHelper) GetErrorMessage(resp *http.Response) (string, error) {
	var errStruct struct{ Errors string }
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, &errStruct)
	return errStruct.Errors, nil
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
		return getDsaStatusDto, customerrors.AccountDoesntExistError{Message: fmt.Sprintf("account %v doesnt exist", accountId)}
	} else if resp.StatusCode == http.StatusInternalServerError {
		msg, err := h.GetErrorMessage(resp)
		if err != nil {
			return getDsaStatusDto, err
		}
		if msg == fmt.Sprintf("no dsa resource exists for account id %v", accountId) {
			return getDsaStatusDto, customerrors.DsaResourceNotFoundError{Message: fmt.Sprintf("no dsa resource exists for account id %v", accountId)}
		} else {
			return getDsaStatusDto, errors.New(msg)
		}
	}
	return getDsaStatusDto, nil
}
