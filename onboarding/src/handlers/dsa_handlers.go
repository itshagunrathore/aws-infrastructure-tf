package handlers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/response"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/services"
)

type DsaHandlers interface {
	ProvisionDsa(context *gin.Context)
	GetDsaStatus(context *gin.Context)
	DeprovisionDsa(context *gin.Context)
	Ping(context *gin.Context)
}
type dsaHandlers struct {
	service services.DsaService
}

func NewDsaHandler(service services.DsaService) *dsaHandlers {
	return &dsaHandlers{service}
}
func (h *dsaHandlers) ProvisionDsa(context *gin.Context) {
	log.Infow("request received for provisioning dsa", "baas-trace-id", context.Value("baas-trace-id"))
	accountId := context.Param("account-id")
	err := h.service.ProvisionDsaService(context, accountId)
	if err != nil {
		log.Infow(fmt.Sprintf("Error While provisioning Dsa: %v", err.Error()), "baas-trace-id", context.Value("baas-trace-id"))
		if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaAlreadyProvisionedError{}) {
			response.SuccessResponseHandler(context, "DSA is already provisioned", http.StatusMethodNotAllowed)
			return
		} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaIsDeployingError{}) {
			response.ErrorResponseHandler(context, err, http.StatusConflict)
			return
		} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaAlreadyProvisionedByOtherEntityError{}) {
			response.ErrorResponseHandler(context, err, http.StatusConflict)
			return
		}
		response.ErrorResponseHandler(context, err, http.StatusInternalServerError)
		return
	}
	response.SuccessResponseHandler(context, "DSA provisoned Successfully", http.StatusOK)

}
func (h *dsaHandlers) GetDsaStatus(context *gin.Context) {
	log.Infow("request received for getting dsa status", "baas-trace-id", context.Value("baas-trace-id"))
	accountId := context.Param("account-id")
	resp, err := h.service.GetDsaStatusService(context, accountId)
	if err != nil {
		log.Infow(fmt.Sprintf("Error While getting dsa status: %v", err.Error()), "baas-trace-id", context.Value("baas-trace-id"))
		if reflect.TypeOf(err) == reflect.TypeOf(customerrors.AccountDoesntExistError{}) {
			response.ErrorResponseHandler(context, err, http.StatusNotFound)
			return
		} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaResourceNotFoundError{}) {
			response.ErrorResponseHandler(context, err, http.StatusUnprocessableEntity)
			return
		}
		response.ErrorResponseHandler(context, err, http.StatusInternalServerError)
		return
	}
	response.SuccessResponseHandler(context, resp, http.StatusOK)
}
func (h *dsaHandlers) DeprovisionDsa(context *gin.Context) {
	log.Infow("request received for deprovisioning dsa", "baas-trace-id", context.Value("baas-trace-id"))
	accountId := context.Param("account-id")
	err := h.service.DeprovisionDsaService(context, accountId)
	if err != nil {
		log.Infow(fmt.Sprintf("Error While Deprovisioning Dsa: %v", err.Error()), "baas-trace-id", context.Value("baas-trace-id"))
		if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaNotProvisionedError{}) {
			response.ErrorResponseHandler(context, err, http.StatusBadRequest)
			return
		}
		response.ErrorResponseHandler(context, err, http.StatusInternalServerError)
		return
	}
	response.SuccessResponseHandler(context, "Dsa Deprovisioned Successfully", http.StatusOK)
}
func (h *dsaHandlers) Ping(context *gin.Context) {
	response.SuccessResponseHandler(context, "pong", http.StatusOK)
}
