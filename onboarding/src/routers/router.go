package routers

import (
	helper "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/helpers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/middlewares"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/web"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/handlers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/services"
)

type RouterStruct struct {
	router web.Router
}

func NewRoute(r web.Router) RouterStruct {
	return RouterStruct{r}
}

//add routes here
func (r *RouterStruct) InitiateDSARoutes(dsaService services.DsaService, helper helper.Helper) {
	dsaHandlers := handlers.NewDsaHandler(dsaService, helper)
	r.router.Engine.Use(middlewares.Tracer())
	r.router.Engine.POST("/v1/accounts/:account-id/dsa", dsaHandlers.ProvisionDsa)
	r.router.Engine.GET("/v1/accounts/:account-id/dsa", dsaHandlers.GetDsaStatus)
	r.router.Engine.DELETE("/v1/accounts/:account-id/dsa", dsaHandlers.DeprovisionDsa)
	r.router.Engine.GET("/ping", dsaHandlers.Ping)
}
