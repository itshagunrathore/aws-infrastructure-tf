package routers

import (
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
func (r *RouterStruct) InitiateDSARoutes(dsaService services.DsaService) {
	dsaHandlers := handlers.NewDsaHandler(dsaService)
	r.router.Engine.Use(middlewares.Tracer())
	r.router.Engine.POST("/v1/accounts/:account-id/dsa", dsaHandlers.ProvisionDsa)
	r.router.Engine.GET("/v1/accounts/:account-id/dsa", dsaHandlers.GetDsaStatus)
	r.router.Engine.DELETE("/v1/accounts/:account-id/dsa", dsaHandlers.DeprovisionDsa)
	r.router.Engine.GET("/ping", dsaHandlers.Ping)
}

// func (r *RouterStruct) GetJobHandlers(jobService services.JobService) { //rename to job dsahandlers

// 	jobHandlers := handlers.NewJobHandler(jobService)
// 	r.router.Engine.Use(middlewares.Tracer())
// 	r.router.Engine.POST("/baas-api/v1/accounts/:account-id/jobs", jobHandlers.PostJob)
// 	r.router.Engine.GET("/baas-api/v1/accounts/:account-id/jobs/:job-id", jobHandlers.GetJob)
// }
