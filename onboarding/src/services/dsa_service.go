package services

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/constants"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	podaccountservice "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/pod_account_service"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/helpers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/mappers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/repositories"
	"gorm.io/gorm"
)

type DsaService interface {
	ProvisionDsaService(context *gin.Context, accountId string) error
	GetDsaStatusService(context *gin.Context, accountId string) (models.DscInstanceDetails, error)
	DeprovisionDsaService(context *gin.Context, accountId string) error
}

type dsaService struct {
	DsaClientSessionRepository repositories.DsaClientSessionRepository
	PodAccountService          podaccountservice.PodAccountService
}

func NewDsaService(r repositories.DsaClientSessionRepository, p podaccountservice.PodAccountService) *dsaService {
	return &dsaService{DsaClientSessionRepository: r, PodAccountService: p}
}
func (d *dsaService) ProvisionDsaService(context *gin.Context, accountId string) error {
	apiPath := fmt.Sprintf("/v1/accounts/%s/dsa", accountId)

	// check if dsa is already provisioned
	dsaStatusResp, err := d.PodAccountService.GetDsaStatus(apiPath, accountId)
	log.Infow(fmt.Sprintf("Response from get dsa service: %v", dsaStatusResp), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil && reflect.TypeOf(err) != reflect.TypeOf(customerrors.DsaResourceNotFoundError{}) {
		return err
	}
	err = helpers.NewDsaHelper().CheckDsaStatus(dsaStatusResp)
	if err != nil {
		return err
	}

	// input for pod acc service
	var provisionDsaModel models.ProvisionDsaModel
	provisionDsaModel.ClientName = constants.ClientName
	// this should be auto populated by dsa prov api but since this is a bug we need to give the image id for dsa manually
	provisionDsaModel.ImageId = "ami-0b81c4b0cbbf63f1a"

	resp, err := d.PodAccountService.ProvisionDsa(apiPath, provisionDsaModel)
	if err != nil {
		return err
	}
	provisionDsaEntity := mappers.NewDsaMapper().MapProvisionDsaResponse(resp, accountId)

	return d.DsaClientSessionRepository.Post(provisionDsaEntity)

}

func (d *dsaService) DeprovisionDsaService(context *gin.Context, accountId string) error {
	// check and get the latest clientSessionId to deprovision for the account
	getDsaClientSessionEntity := mappers.NewDsaMapper().MapDsaClientSessionGetRequest(accountId)
	dsaClientSession, err := d.DsaClientSessionRepository.GetProvisionedAccounts(&getDsaClientSessionEntity)
	fmt.Println(dsaClientSession.ClientSessionId)

	if reflect.TypeOf(err) == reflect.TypeOf(gorm.ErrRecordNotFound) {
		return customerrors.NewDsaNotProvisionedError("Dsa has not been provisioned for this account")
	} else if err != nil {
		return err
	}
	//input for pod-acc-svc
	apiPath := fmt.Sprintf("/v1/accounts/%s/%s/%s/dsa", accountId, constants.ClientName, dsaClientSession.ClientSessionId)

	resp, statusCode, err := d.PodAccountService.DeprovisionDsa(apiPath)
	log.Info(fmt.Sprintf("Response for deprovisioning dsa: %v", resp), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil {
		return err
	}

	dsaClientSessionEntityResp, err := helpers.NewDsaHelper().HandleDeprovisioningResponse(resp, statusCode, dsaClientSession.ClientSessionId)
	if err != nil {
		return err
	}

	return d.DsaClientSessionRepository.Update(dsaClientSessionEntityResp)

}

func (d *dsaService) GetDsaStatusService(context *gin.Context, accountId string) (models.DscInstanceDetails, error) {
	apiPath := fmt.Sprintf("/v1/accounts/%s/dsa", accountId)
	var getDsaStatus models.DscInstanceDetails
	resp, err := d.PodAccountService.GetDsaStatus(apiPath, accountId)
	log.Info(fmt.Sprintf("Response for get dsa status: %v,", resp), "baas-trace-id", context.Value("baas-trace-id"))
	if err != nil {
		return getDsaStatus, err
	}
	return resp, err
}
