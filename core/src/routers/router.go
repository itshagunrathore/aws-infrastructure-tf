package routers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/web"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/handlers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/services"
)

type RouterStruct struct {
	router web.Router
}

func NewRoute(r web.Router) RouterStruct {
	return RouterStruct{r}
}

func (r *RouterStruct) GetRoute() { //rename to job handlers
	dbConfig := db.DbConfig{
		Username: "dev_admin",
		Password: "postgre&308",
		Port:     80,
		Host:     "baas-rds-dev-725b87755a61c35c.elb.us-west-2.amazonaws.com",
		DbName:   "baas_dev_db"}
	DB := db.NewDBConnection(dbConfig)
	jobDefinitionRepository := repositories.NewJobDefinitionRepository(DB)
	jobService := services.NewJobService(jobDefinitionRepository)
	jobHandlers := handlers.NewJobHandler(jobService)
	r.router.Engine.POST("/baas-api/v1/accounts/:account-id/jobs", jobHandlers.PostJob)

}
