package podaccountservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/config"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	httpclient "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/http"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/helpers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type PodAccountService interface {
	ProvisionDsa(path string, provisionDsaModel models.ProvisionDsaModel) (models.ProvisionDsaResponseModel, error)
	GetDsaStatus(path, accountId string) (models.DscInstanceDetails, error)
	DeprovisionDsa(path string) ([]byte, int, error)
}

type podAccountService struct {
	endpoint   string
	httpClient httpclient.HttpClient
}

func NewPodAccountService() *podAccountService {
	endpoint := config.GetConfig("podAccountService.endpoint")
	fmt.Println("The endpoint received: ", endpoint)
	httpClient := httpclient.NewHttpClient(false)
	return &podAccountService{endpoint: endpoint, httpClient: httpClient}
}

func (p *podAccountService) ProvisionDsa(path string, provisionDsaModel models.ProvisionDsaModel) (models.ProvisionDsaResponseModel, error) {
	var buf bytes.Buffer
	var provisionDsaResponseModel models.ProvisionDsaResponseModel
	helper := helpers.NewHelper()
	err := json.NewEncoder(&buf).Encode(provisionDsaModel)
	if err != nil {
		return provisionDsaResponseModel, err
	}
	resp, statusCode, err := p.httpClient.Post(p.endpoint+path, buf)
	if err != nil {
		return provisionDsaResponseModel, err
	}
	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		json.Unmarshal(resp, &provisionDsaResponseModel)
		return provisionDsaResponseModel, nil
	} else if statusCode == http.StatusMethodNotAllowed {
		msg := helper.GetErrorMessage(resp)
		return provisionDsaResponseModel, customerrors.NewDsaAlreadyProvisionedError(msg)
	} else {
		msg := helper.GetErrorMessage(resp)
		fmt.Println(string(resp))
		return provisionDsaResponseModel, fmt.Errorf("error provisioning dsa: %v", msg)
	}
}

func (p *podAccountService) GetDsaStatus(path, accountId string) (models.DscInstanceDetails, error) {
	var getDsaStatusDto models.DscInstanceDetails
	resp, statusCode, err := p.httpClient.Get(p.endpoint + path)

	if err != nil {
		return getDsaStatusDto, err
	}

	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		json.Unmarshal(resp, &getDsaStatusDto)
		return getDsaStatusDto, nil
	} else if statusCode == http.StatusNotFound {
		return getDsaStatusDto, customerrors.NewAccountDoesntExistError(fmt.Sprintf("account %v doesnt exist", accountId))
	} else if statusCode == http.StatusInternalServerError {
		msg := helpers.NewHelper().GetErrorMessage(resp)
		if msg == fmt.Sprintf("no dsa resource exists for account id %v", accountId) {
			return getDsaStatusDto, customerrors.NewDsaResourceNotFoundError(fmt.Sprintf("no dsa resource exists for account id %v", accountId))
		} else {
			return getDsaStatusDto, errors.New(msg)
		}
	} else {
		msg := helpers.NewHelper().GetErrorMessage(resp)
		return getDsaStatusDto, errors.New(msg)
	}
}

func (p *podAccountService) DeprovisionDsa(path string) ([]byte, int, error) {
	resp, statusCode, err := p.httpClient.Delete(p.endpoint + path)
	if err != nil {
		return nil, 0, err
	}
	return resp, statusCode, nil

}
