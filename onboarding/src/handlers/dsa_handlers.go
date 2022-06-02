package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	err := h.service.ProvisionDsaService(accountId)
	if err != nil {
		response.ErrorResponseHandler(context, err, http.StatusInternalServerError)
		return
	}
	response.SuccessResponseHandler(context, "DSA provisoned Successfully", http.StatusOK)

}
func (h *dsaHandlers) GetDsaStatus(context *gin.Context) {
	log.Infow("request received for getting dsa status", "baas-trace-id", context.Value("baas-trace-id"))

}
func (h *dsaHandlers) DeprovisionDsa(context *gin.Context) {
	log.Infow("request received for deprovisioning dsa", "baas-trace-id", context.Value("baas-trace-id"))
}
func (h *dsaHandlers) Ping(context *gin.Context) {
	response.SuccessResponseHandler(context, "pong", http.StatusOK)
}
