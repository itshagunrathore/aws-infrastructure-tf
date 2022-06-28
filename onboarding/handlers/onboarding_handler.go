package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/response"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/service"
)

func onboarding(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {

		return
	}
	var payload models.Event

	json.Unmarshal(jsonData, &payload)
	payload.Port = "9090"

	// Check DSC status
	dscState := dscStatus("status")
	var Response models.DetailedStatus
	if dscState == "running" {
		_, err := service.ConfigureSystem(payload, &Response)
		if err == nil {
			Response.OnboardingStatus = "Failed"
			Response.CustomerAccount = payload.AcctName
			log.Error(Response)
			response.ErrorResponseHandler(c, Response, Response.StatusCode)
			return
		} else {
			Response.OnboardingStatus = "Success"
			Response.CustomerAccount = payload.AcctName
			log.Error(Response)
			response.SuccessResponseHandler(c, err, Response.StatusCode)
			return
		}
	} else {
		apiResponse := dscStatus("start")
		log.Info(apiResponse)
		response.ErrorResponseHandler(c, err, Response.StatusCode)
		return
	}
}

func reonboarding(c *gin.Context) {
	// TODO Code for reonboarding here
}

func runjob(c *gin.Context) {
	// Code for runjob here

}

func dscStatus(operation string) string {
	// TODO dsc status code here
	var apiResponse string
	switch operation {
	case "status":
		apiResponse = "running"
	case "start":
		apiResponse = "starting dsc"
	case "stop":
		apiResponse = "stopping dsc"
	}
	return apiResponse
}

func main() {
	// TODO environment need to be fetched from config
	log.InitiateLogger("INFO", "dev")
	router := gin.Default()
	router.POST("/baas-api/v1/onboard", onboarding)
	router.GET("/baas-api/v1/tenant-reconfigure/", reonboarding)
	router.POST("/baas-api/v1/runjob", runjob)
	router.Run("localhost:8080")
}
