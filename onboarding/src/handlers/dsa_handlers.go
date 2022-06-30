package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/helpers"
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
	helper  helpers.Helper
}

func NewDsaHandler(service services.DsaService, helper helpers.Helper) *dsaHandlers {
	return &dsaHandlers{service, helper}
}
func (h *dsaHandlers) ProvisionDsa(context *gin.Context) {
	log.Infow("request received for provisioning dsa", "baas-trace-id", context.Value("baas-trace-id"))
	accountId := context.Param("account-id")
	err := h.service.ProvisionDsaService(context, accountId)
	if err != nil {
		log.Errorw(err.Error(), "baas-trace-id", context.Value("baas-trace-id"))
		h.helper.GetErrorResponse(context, err)
		return
	}
	response.SuccessResponseHandler(context, "DSA provisoned Successfully", http.StatusOK)

}
func (h *dsaHandlers) GetDsaStatus(context *gin.Context) {
	log.Infow("request received for getting dsa status", "baas-trace-id", context.Value("baas-trace-id"))
	accountId := context.Param("account-id")
	resp, err := h.service.GetDsaStatusService(context, accountId)
	if err != nil {
		log.Errorw(err.Error(), "baas-trace-id", context.Value("baas-trace-id"))
		h.helper.GetErrorResponse(context, err)
		return
	}
	response.SuccessResponseHandler(context, resp, http.StatusOK)
}
func (h *dsaHandlers) DeprovisionDsa(context *gin.Context) {
	log.Infow("request received for deprovisioning dsa", "baas-trace-id", context.Value("baas-trace-id"))
	accountId := context.Param("account-id")
	err := h.service.DeprovisionDsaService(context, accountId)
	if err != nil {
		log.Errorw(err.Error(), "baas-trace-id", context.Value("baas-trace-id"))
		h.helper.GetErrorResponse(context, err)
		return
	}
	response.SuccessResponseHandler(context, "Dsa Deprovisioned Successfully", http.StatusOK)
}
func (h *dsaHandlers) Ping(context *gin.Context) {
	response.SuccessResponseHandler(context, "pong", http.StatusOK)
}
