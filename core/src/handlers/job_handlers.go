package handlers

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/response"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

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
	log.Infow("request received for get job", "baas-trace-id", context.Value("baas-trace-id"))

	accountId := context.Param("account-id")
	jobId, err := strconv.Atoi(context.Param("job-id"))

	if err != nil {
		err = customerrors.NewBadRequest("job-id must be integer")
		response.ErrorResponseHandler(context, err, http.StatusBadRequest)
		return
	}
	getJobDto, err := handler.service.GetJob(context, accountId, jobId)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(customerrors.ResourceNotFound{}) {
			response.SuccessResponseHandler(context, nil, http.StatusNoContent)
			return
		}
		err = customerrors.NewInternalServerError("failed to get job")
		response.ErrorResponseHandler(context, err, http.StatusInternalServerError)
		return
	}
	log.Infow(fmt.Sprintf("get job successfull for account id: %s and job-id: %d", accountId, jobId), "baas-trace-id", context.Value("baas-trace-id"))
	response.SuccessResponseHandler(context, getJobDto, http.StatusOK)
	return
}

func (handler *jobHandlers) GetAllJob(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handler *jobHandlers) PostJob(context *gin.Context) {
	log.Infow("request received for post job", "baas-trace-id", context.Value("baas-trace-id"))
	var postJobDto dto.PostJobDto

	accountId := context.Param("account-id")
	jsonData, err := ioutil.ReadAll(context.Request.Body)

	log.Infow("request body is", "baas-trace-id", context.Value("baas-trace-id"), "body", string(jsonData))
	err = json.Unmarshal(jsonData, &postJobDto)
	if err != nil {
		msg := fmt.Sprintf("error occured while reading request body %v", err.Error())
		log.Errorw(msg, context.Value("baas-trace-id"))
		err = customerrors.NewBadRequest(msg)
		response.ErrorResponseHandler(context, err, http.StatusBadRequest)
		return
	}
	// TODO call audit logic here
	jobId, err := handler.service.CreateJob(context, accountId, postJobDto)

	if err != nil {
		log.Errorw(fmt.Sprintf("error occured %s", err.Error()), "baas-trace-id", context.Value("baas-trace-id"))
		if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DuplicateResourceFound{}) {
			response.ErrorResponseHandler(context, err, http.StatusConflict)
			return
		}
		if reflect.TypeOf(err) == reflect.TypeOf(validation.Errors{}) {
			response.ErrorResponseHandler(context, err, http.StatusBadRequest)
			return
		}
		// creating a generic error here
		err = customerrors.NewInternalServerError("failed to create job")
		response.ErrorResponseHandler(context, err, http.StatusInternalServerError)
		return
	}
	msg := fmt.Sprintf("job created with job-id %v", jobId)
	log.Infow(msg, "baas-trace-id", context.Value("baas-trace-id"))
	response.SuccessResponseAccepted(context, strconv.Itoa(jobId))
}
