package services

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/config"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	podaccountservice "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/pod_account_service"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
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
	GetDsaStatusService(context *gin.Context, accountId string) (models.DscInstanceDetails, error)
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
	baseurl := config.GetConfig("dummyUrl")

	// check if dsa is already provisioned
	dsaStatusResp, err := d.GetDsaStatusService(context, accountId)
	log.Infow(fmt.Sprintf("Response from get dsa service: %v", dsaStatusResp.Error), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil && reflect.TypeOf(err) != reflect.TypeOf(customerrors.DsaResourceNotFoundError{}) {
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
	resp, err := podAccountService.ProvisionDsa(baseurl, apiPath, provisionDsaModel)
	if err != nil {
		return err
	}
	provisionDsaEntity := mappers.NewDsaMapper().MapProvisionDsaResponse(resp, accountId)

	return d.DsaClientSessionRepository.Post(provisionDsaEntity)

}

func (d *dsaService) DeprovisionDsaService(context *gin.Context, accountId string) error {

	// check and get the latest clientSessionId to deprovision for the account
	getDsaClientSessionEntity := mappers.NewDsaMapper().MapDsaClientSessionGetRequest(accountId)
	dsaClientSession, err := d.DsaClientSessionRepository.Get(getDsaClientSessionEntity)
	if reflect.TypeOf(err) == reflect.TypeOf(gorm.ErrRecordNotFound) {
		return customerrors.NewDsaNotProvisionedError("Dsa has not been provisioned for this account")
	} else if err != nil {
		return err
	}

	//input for pod-acc-svc
	apiPath := fmt.Sprintf("/v1/accounts/%s/%s/%s/dsa", accountId, ClientName, dsaClientSession.ClientSessionId)
	baseurl := config.GetConfig("dummyUrl")
	podAccSvc := podaccountservice.NewPodAccountService()

	resp, statusCode, err := podAccSvc.DeprovisionDsa(baseurl, apiPath)
	log.Info(fmt.Sprintf("Response for deprovisioning dsa: %v", statusCode), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil {
		return err
	}

	dsaClientSessionEntityResp, err := helpers.NewDsaHelper().DeprovisionDsaHelper(resp, statusCode, dsaClientSession.ClientSessionId)
	if err != nil {
		return err
	}

	return d.DsaClientSessionRepository.Update(dsaClientSessionEntityResp)

}

func (d *dsaService) GetDsaStatusService(context *gin.Context, accountId string) (models.DscInstanceDetails, error) {
	apiPath := fmt.Sprintf("/v1/accounts/%s/dsa", accountId)
	baseurl := config.GetConfig("dummyUrl")
	var getDsaStatus models.DscInstanceDetails

	podAccSvc := podaccountservice.NewPodAccountService()
	resp, err := podAccSvc.GetDsaStatus(baseurl, apiPath, accountId)
	log.Info(fmt.Sprintf("Response for get dsa status: %v", resp), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil {
		return getDsaStatus, err
	}
	return resp, err
}
