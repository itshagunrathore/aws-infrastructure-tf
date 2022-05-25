package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponseHandler(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func SuccessResponseHandler(c *gin.Context, body interface{}, statusCode int) {
	c.JSON(statusCode, body)
}

func SuccessResponseCreated(c *gin.Context, resourceID string, body interface{}) {
	c.Header("x-resource-id", resourceID)
	c.JSON(http.StatusCreated, body)
}
