package routers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/web"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/handlers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/services"
)

type RouterStruct struct {
	router web.Router
}

func NewRoute(r RouterStruct) RouterStruct {
	return r
}

func (r *RouterStruct) GetRoute() {

	DB := NewDatabase(config.Read())
	jobDefinitionRepository := repositories.NewJobDefinitionRepository(DB)
	jobService := services.NewJobService(jobDefinitionRepository)
	jobHandlers := handlers.NewJobHandler(jobService)
	r.router.POST("/baas-api/v1/accounts/:account-id/jobs", jobHandlers.PostJob)
}
