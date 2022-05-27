package routers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/middlewares"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/web"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/handlers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/services"
)

type RouterStruct struct {
	router web.Router
}

func NewRoute(r web.Router) RouterStruct {
	return RouterStruct{r}
}

func (r *RouterStruct) GetJobHandlers(jobService services.JobService) { //rename to job dsahandlers

	jobHandlers := handlers.NewJobHandler(jobService)
	r.router.Engine.Use(middlewares.Tracer())
	r.router.Engine.POST("/baas-api/v1/accounts/:account-id/jobs", jobHandlers.PostJob)
	r.router.Engine.GET("/baas-api/v1/accounts/:account-id/jobs/:job-id", jobHandlers.GetJob)
}
