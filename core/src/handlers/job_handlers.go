package handlers

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/middlewares"
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

	// call audit here
	jobId, err := handler.service.CreateJob(context, accountId, postJobDto)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(customerrors.JobAlreadyExistsError{}) {
			// return 409 conflict
			middlewares.ErrorHandler(context, err, http.StatusConflict)
			return
		}
		if reflect.TypeOf(err) == reflect.TypeOf(validation.Errors{}) {
			middlewares.ErrorHandler(context, err, http.StatusBadRequest)
			return
			//return 400 bad request
		}
		//if reflect.TypeOf(err) == reflect.TypeOf(customerrors.InternalServerError{}) {
		//	// return 500
		//	return
		//}
	}

	//add newly created job definition in body and return json with 202
	if err != nil {
		panic("error")
		return
	}
	context.JSON(http.StatusAccepted, jobId)
}
