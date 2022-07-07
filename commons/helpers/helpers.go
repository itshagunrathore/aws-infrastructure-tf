package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/response"
)

type Helper interface {
	GetErrorResponse(context *gin.Context, err error)
	GetErrorMessage(resp []byte) string
}
type helper struct {
}

func NewHelper() *helper {
	return &helper{}
}
func (h *helper) GetErrorMessage(resp []byte) string {
	var errStruct struct{ Error string }
	json.Unmarshal(resp, &errStruct)
	return errStruct.Error
}
func (h *helper) GetErrorResponse(context *gin.Context, err error) {
	if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaAlreadyProvisionedError{}) {
		response.ErrorResponseHandler(context, errors.New("DSA is already provisioned"), http.StatusMethodNotAllowed)
		return
	} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaIsDeployingError{}) {
		response.ErrorResponseHandler(context, err, http.StatusConflict)
		return
	} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaAlreadyProvisionedByOtherEntityError{}) {
		response.ErrorResponseHandler(context, err, http.StatusConflict)
		return
	} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.AccountDoesntExistError{}) {
		response.ErrorResponseHandler(context, err, http.StatusNotFound)
		return
	} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaResourceNotFoundError{}) {
		response.ErrorResponseHandler(context, err, http.StatusUnprocessableEntity)
		return
	} else if reflect.TypeOf(err) == reflect.TypeOf(customerrors.DsaNotProvisionedError{}) {
		response.ErrorResponseHandler(context, err, http.StatusBadRequest)
		return
	}
	response.ErrorResponseHandler(context, customerrors.NewInternalServerError(""), http.StatusInternalServerError)
}
