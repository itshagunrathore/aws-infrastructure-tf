package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/services"
)

type JobHandlers interface {
	GetJob(context *gin.Context) error
	GetAllJob(context *gin.Context) error
	PostJob(context *gin.Context) error
}

type jobHandlers struct {
	service services.JobService
}

func NewJobHandler(service services.JobService) JobHandlers {
	return &jobHandlers{service}
}

func (handler *jobHandlers) GetJob(context *gin.Context) error {
	return nil
}

func (handler *jobHandlers) GetAllJob(context *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (handler *jobHandlers) PostJob(context *gin.Context) error {

	var postJobDto dto.PostJobDto
	accountId := context.Param("account-id")

	if err := context.BindJSON(&postJobDto); err != nil {
		return nil
	}

	jobId, err := handler.service.CreateJob(context, accountId, postJobDto)

	if err != nil {

	}
	context.JSON(http.StatusAccepted, jobId)
	return nil
}
