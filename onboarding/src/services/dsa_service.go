package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	podaccountservice "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/pod_account_service"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/dtos"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/entities"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/helpers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/mappers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/repositories"
	"gorm.io/gorm"
)

const (
	ClientName         string = "baas" //
	Running            string = "running"
	Deploying          string = "deploying"
	ProvisionDsaPath   string = "/v1/accounts/{accountId}/dsa"
	DeprovisionDsaPath string = "/v1/accounts/{accountId}/{clientName}/{clientSessionId}/dsa"
)

type DsaService interface {
	ProvisionDsaService(context *gin.Context, accountId string) error
	GetDsaStatusService(context *gin.Context, accountId string) (dtos.GetDsaStatusDtos, error)
	DeprovisionDsaService(context *gin.Context, accountId string) error
}

type dsaService struct {
	DsaClientSessionRepository repositories.DsaClientSessionRepository
}

func NewDsaService(r repositories.DsaClientSessionRepository) *dsaService {
	return &dsaService{DsaClientSessionRepository: r}
}
func (d *dsaService) ProvisionDsaService(context *gin.Context, accountId string) error {
	apiPath := fmt.Sprintf("/v1/accounts/%s/dsa", accountId)
	baseurl := viper.GetString("dummyUrl")
	var provisionDsaResponseDto dtos.ProvisionDsaDtos
	var provisionDsaEntity entities.DsaClientSession
	mappers := mappers.NewDsaMapper()

	// check if dsa is already provisioned
	dsaStatusResp, err := d.GetDsaStatusService(context, accountId)
	log.Infow(fmt.Sprintf("Response from get dsa service: %v", dsaStatusResp), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil {
		return err
	}
	err = helpers.NewDsaHelper().CheckDsaStatus(dsaStatusResp)
	if err != nil {
		return err
	}

	// input for pod acc service
	var provisionDsaModel models.ProvisionDsaModel
	provisionDsaModel.ClientName = ClientName

	podAccountService := podaccountservice.NewPodAccountService()
	resp, err := podAccountService.ProvisionDsa(baseurl, apiPath, false, provisionDsaModel)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Infow(fmt.Sprintf("pod-account-service response, statusCode: %v, body %v", resp.StatusCode, string(body)), "baas-trace-id", context.Value("baas-trace-id"))

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		json.Unmarshal(body, &provisionDsaResponseDto)
		provisionDsaEntity = mappers.MapProvisionDsaResponse(provisionDsaResponseDto, accountId)
		err := d.DsaClientSessionRepository.Post(provisionDsaEntity)
		if err != nil {
			return err
		}
		return nil
	} else if resp.StatusCode == http.StatusMethodNotAllowed {
		return customerrors.DsaAlreadyProvisionedError{Message: fmt.Sprintf("Dsa already provisioned by %v", ClientName)}
	} else {
		msg, err := helpers.NewDsaHelper().GetErrorMessage(resp)
		if err != nil {
			return err
		}
		return fmt.Errorf("error provisioning dsa: %v", msg)
	}
}

func (d *dsaService) DeprovisionDsaService(context *gin.Context, accountId string) error {

	// check and get the latest clientSessionId to deprovision for the account
	dsaClientSessionEntity := mappers.NewDsaMapper().MapDsaClientSessionGetRequest(accountId)
	dsaClientSession, err := d.DsaClientSessionRepository.Get(dsaClientSessionEntity)
	if reflect.TypeOf(err) == reflect.TypeOf(gorm.ErrRecordNotFound) {
		return customerrors.DsaNotProvisionedError{Message: "Dsa has not been provisioned for this account"}
	} else if err != nil {
		return err
	}

	//input for pod-acc-svc
	apiPath := fmt.Sprintf("/v1/accounts/%s/%s/%s/dsa", accountId, ClientName, dsaClientSession.ClientSessionId)
	baseurl := viper.GetString("dummyUrl")
	podAccSvc := podaccountservice.NewPodAccountService()

	resp, err := podAccSvc.DeprovisionDsa(baseurl, apiPath, false)
	log.Info(fmt.Sprintf("Response for deprovisioning dsa: %v", resp), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		dsaClientSessionEntity.IsDeleted = true
		dsaClientSessionEntity.TimeUpdated = time.Now().UTC()
		err := d.DsaClientSessionRepository.Update(dsaClientSessionEntity)
		if err != nil {
			return err
		}
		return nil
	} else if resp.StatusCode == http.StatusNotFound {
		return customerrors.AccountDoesntExistError{Message: fmt.Sprintf("Account %v doesnt exist", accountId)}
	} else {
		msg, err := helpers.NewDsaHelper().GetErrorMessage(resp)
		if err != nil {
			return err
		}
		return fmt.Errorf("error deprovisioning dsa: %v", msg)
	}

}

func (d *dsaService) GetDsaStatusService(context *gin.Context, accountId string) (dtos.GetDsaStatusDtos, error) {
	apiPath := fmt.Sprintf("/v1/accounts/%s/dsa", accountId)
	baseurl := viper.GetString("dummyUrl")
	var getDsaStatusDto dtos.GetDsaStatusDtos

	podAccSvc := podaccountservice.NewPodAccountService()
	resp, err := podAccSvc.GetDsaStatus(baseurl, apiPath, false)
	log.Info(fmt.Sprintf("Response for get dsa status: %v", resp), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil {
		return getDsaStatusDto, err
	}
	getDsaStatusDto, err = helpers.NewDsaHelper().GetDsaStatusHelper(resp, accountId)
	if err != nil {
		return getDsaStatusDto, err
	}
	return getDsaStatusDto, nil
}
