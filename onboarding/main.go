package main

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/web"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/repositories"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/routers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/services"
)

func main() {
	log.InitiateLogger("INFO", "dev")

	dbConfig := db.DbConfig{
		Username: "dev_admin",
		Password: "postgre&308",
		Port:     80,
		Host:     "baas-rds-dev-725b87755a61c35c.elb.us-west-2.amazonaws.com",
		DbName:   "baas_dev_db"}
	DB := db.NewDBConnection(dbConfig)
	log.Info(DB)

	dsaClientSessionRepository := repositories.NewDsaClientSessionRepository(DB)
	dsaService := services.NewDsaService(dsaClientSessionRepository)
	//routing
	router := web.NewRouter()
	r := routers.NewRoute(*router)
	r.InitiateDSARoutes(dsaService)
	router.Engine.Run()

}
