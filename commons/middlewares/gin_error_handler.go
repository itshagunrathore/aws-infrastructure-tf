package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
