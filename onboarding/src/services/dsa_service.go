package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	podaccountservice "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/pod_account_service"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/dtos"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/entities"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/mappers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/repositories"
)

const (
	ClientName         string = "baas" //
	ProvisionDsaPath   string = "/v1/accounts/{accountId}/dsa"
	DeprovisionDsaPath string = "/v1/accounts/{accountId}/{clientName}/{clientSessionId}/dsa"
)

type DsaService interface {
	ProvisionDsaService(accountId string) error
	GetDsaStatusService()
	DeprovisionDsaService()
}

type dsaService struct {
	DsaClientSessionRepository repositories.DsaClientSessionRepository
}

func NewDsaService(r repositories.DsaClientSessionRepository) *dsaService {
	return &dsaService{DsaClientSessionRepository: r}
}
func (d *dsaService) ProvisionDsaService(accountId string) error {
	apiPath := fmt.Sprintf("/v1/accounts/%s/dsa", accountId)
	baseurl := viper.GetString("dummyUrl")

	var provisionDsaResponseDto dtos.ProvisionDsaDtos
	var provisionDsaEntity entities.DsaClientSession

	// input for pod acc service
	var provisionDsaModel models.ProvisionDsaModel
	provisionDsaModel.ClientName = ClientName
	podAccountService := podaccountservice.NewPodAccountService()

	resp, err := podAccountService.ProvisionDsa(baseurl, apiPath, false, provisionDsaModel)
	if err != nil {
		return err
	}
	log.Info(resp)
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		json.Unmarshal(body, &provisionDsaResponseDto)
		provisionDsaEntity = mappers.NewDsaMapper().MapProvisionDsaResponse(provisionDsaResponseDto, accountId)
		d.DsaClientSessionRepository.Post(provisionDsaEntity)
		return nil
	}
	//to do handle diff status
	panic(resp)

}
func (d *dsaService) DeprovisionDsaService() {

}
func (d *dsaService) GetDsaStatusService() {

}
