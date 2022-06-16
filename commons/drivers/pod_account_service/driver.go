package podaccountservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/config"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	httpClient "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/http"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/helpers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type PodAccountService interface {
	ProvisionDsa(url, path string, secure bool, provisionDsaModel models.ProvisionDsaModel) ([]byte, error)
}

type podAccountService struct {
	endpoint   string
	httpClient httpClient.HttpClient
}

func NewPodAccountService() *podAccountService {
	endpoint := config.GetConfig("podAccountService.endpoint")
	fmt.Println("The endpoint received: ", endpoint)
	httpClient := httpClient.NewHttpClient(false)
	return &podAccountService{endpoint: endpoint, httpClient: httpClient}
}

func (p *podAccountService) ProvisionDsa(url, path string, provisionDsaModel models.ProvisionDsaModel) (models.ProvisionDsaResponseModel, error) {
	var buf bytes.Buffer
	var provisionDsaResponseModel models.ProvisionDsaResponseModel
	helper := helpers.NewHelper()
	err := json.NewEncoder(&buf).Encode(provisionDsaModel)
	if err != nil {
		return provisionDsaResponseModel, err
	}
	resp, statusCode, err := p.httpClient.Post(url+path, buf)
	if err != nil {
		return provisionDsaResponseModel, err
	}
	if statusCode == http.StatusOK || statusCode == http.StatusAccepted {
		json.Unmarshal(resp, &provisionDsaResponseModel)
		return provisionDsaResponseModel, nil
	} else if statusCode == http.StatusMethodNotAllowed {
		msg := helper.GetErrorMessage(resp)
		return provisionDsaResponseModel, customerrors.NewDsaAlreadyProvisionedError(msg)
	} else {
		msg := helper.GetErrorMessage(resp)
		return provisionDsaResponseModel, fmt.Errorf("error provisioning dsa: %v", msg)
	}
}

// func (p *podAccountService) GetAccountInfoById(accountId string) (models.AccountDetails, error) {
// 	url := p.endpoint + "/v1/accounts/" + accountId
// 	resp, err := p.httpClient.Get(url)
// 	if err != nil {
// 		return models.AccountDetails{}, err
// 	}
// 	var accountDetails models.AccountDetails
// 	if err := json.Unmarshal(resp, &accountDetails); err != nil { // Parse []byte to the go struct pointer
// 		return models.AccountDetails{}, err
// 	}
// 	return accountDetails, nil
// }

// func (p *podAccountService) GetDscInstanceInfo(accountId string) (models.DscInstanceDetails, error) {
// 	url := p.endpoint + "/v1/accounts/" + accountId + "/dsa"
// 	resp, err := p.httpClient.Get(url)
// 	if err != nil {
// 		return models.DscInstanceDetails{}, err
// 	}
// 	var dscInstanceDetails models.DscInstanceDetails
// 	if err := json.Unmarshal(resp, &dscInstanceDetails); err != nil { // Parse []byte to the go struct pointer
// 		return models.DscInstanceDetails{}, err
// 	}
// 	return dscInstanceDetails, nil
// }

func (p *podAccountService) GetDsaStatus(url, path, accountId string) (models.DscInstanceDetails, error) {
	var getDsaStatusDto models.DscInstanceDetails
	resp, statusCode, err := p.httpClient.Get(url + path)
	if err != nil {
		return getDsaStatusDto, err
	}

	if statusCode == http.StatusOK || statusCode == http.StatusAccepted {
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
	}
	return getDsaStatusDto, nil
}
func (p *podAccountService) DeprovisionDsa(url, path string) ([]byte, int, error) {
	resp, statusCode, err := p.httpClient.Delete(url + path)
	if err != nil {
		return nil, 0, err
	}
	return resp, statusCode, nil

}
