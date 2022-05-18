package handlers

import (
	"encoding/json"
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/errors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/services"
)

type JobHandlers interface {
	GetJob(context *gin.Context)
	GetAllJob(context *gin.Context)
	PostJob(context *gin.Context)
}

type jobHandlers struct {
	service services.JobService
}

func NewJobHandler(service services.JobService) JobHandlers {
	return &jobHandlers{service}
}

func (handler *jobHandlers) GetJob(context *gin.Context) {
	panic("implement me")
}

func (handler *jobHandlers) GetAllJob(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handler *jobHandlers) PostJob(context *gin.Context) {

	var postJobDto dto.PostJobDto
	accountId := context.Param("account-id")
	jsonData, err := ioutil.ReadAll(context.Request.Body)

	if err != nil {
		//internal server error
		return
	}
	log.Info(string(jsonData))
	json.Unmarshal(jsonData, &postJobDto)
	fmt.Println(postJobDto)
	//if err := context.BindJSON(&postJobDto); err != nil {
	//	log.Error(err)
	//}
	log.Info(postJobDto)
	jobId, err := handler.service.CreateJob(context, accountId, postJobDto)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(errors.JobAlreadyExistsError{}){}) {
			// return 409 conflict
			return
		}
		if error.TypeOf(err) == reflect.TypeOf(Vaidation err){
			//return 404 bad request
		}
		if error.TypeOf(err) == reflect.TypeOf(internal server error){
			// return 500
			return
		}
	}

	//add newly created job definition in body and return json with 202
	if err != nil {
		panic("error")
		return
	}
	context.JSON(http.StatusAccepted, jobId)

}
