package main

import (
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/web"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/routers"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/services"
	"os"
)

func main() {
	log.InitiateLogger("INFO", "dev")
	router := web.NewRouter()
	r := routers.NewRoute(*router)
	dbConfig := db.DbConfig{
		Username: "dev_admin",
		Password: "postgre&308",
		Port:     80,
		Host:     "baas-rds-dev-725b87755a61c35c.elb.us-west-2.amazonaws.com",
		DbName:   "baas_dev_db"}
	DB := db.NewDBConnection(dbConfig)
	jobDefinitionRepository := repositories.NewJobDefinitionRepository(DB)
	customerSiteRepository := repositories.NewCustomerSiteRepository(DB)
	latestJobSessionRepository := repositories.NewLatestJobSessionRepository(DB)
	jobService := services.NewJobService(jobDefinitionRepository, customerSiteRepository, latestJobSessionRepository)
	r.GetJobHandlers(jobService)

	router.Engine.Run()
	//http.ListenAndServe(":8070", r)
	//var cfg *models.Configurations
	//config.ReadConfigInto(&cfg)
	//fmt.Println(cfg)
	//fmt.Println("printing configuration values: ", cfg.DbConfig.Host)
	//fmt.Println(cfg.DbConfig.Username)
	//fmt.Println(cfg.DbConfig.Password)
	//fmt.Println(cfg.DbConfig.SSLEnabled)
	// postgresDb := db.NewDBConnection(cfg.DBConfig)
	// dsahandlers.HandleLogging()
	// services.HandleService()
	// fmt.Printf("Running project: `%s`\n", src.ProjectName())

	// These functions demonstrate two separate checks to detect if the code is being
	// run inside a docker container in debug mode, or production mode!
	//
	// Note: Valid only for docker containers generated using the Makefile command
	// firstCheck()
	// secondCheck()
}

func firstCheck() bool {
	/*
	 * The `debug_mode` environment variable exists only in debug builds, likewise,
	 * `production_mode` variable exists selectively in production builds - use the
	 * existence of these variables to detect container build type (and not values)
	 *
	 * Exactly one of these - `production_mode` or `debug_mode` - is **guaranteed** to
	 * exist for docker builds generated using the Makefile commands!
	 */

	if _, ok := os.LookupEnv("production_mode"); ok {
		fmt.Println("(Check 01): Running in `production` mode!")
		return true
	} else if _, ok := os.LookupEnv("debug_mode"); ok {
		// Could also use a simple `else` statement (above) for docker builds!
		fmt.Println("(Check 01): Running in `debug` mode!")
		return true
	} else {
		fmt.Println("\nP.S. Try running a build generated with the Makefile :)")
		return false
	}
}

func secondCheck() bool {
	/*
	 * There's also a central `__BUILD_MODE__` variable for a dirty checks -- guaranteed
	 * to exist for docker containers generated by the Makefile commands!
	 * The variable will have a value of `production` or `debug` (no capitalization)
	 *
	 * Note: Relates to docker builds generated using the Makefile
	 */

	value := os.Getenv("__BUILD_MODE__")

	// Yes, this if/else could have been written better
	switch value {
	case "production":
		fmt.Println("(Check 02): Running in `production` mode!")
		return true

	case "debug":
		fmt.Println("(Check 02): Running in `debug` mode!")
		return true

	default:
		// Flow ends up here for non-docker builds, or docker builds generated directly
		fmt.Println("Non-makefile build detected :(")
		return false
	}
}
