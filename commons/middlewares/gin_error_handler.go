package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, err error) {
	if err == ErrNotFound {
		// 404
		c.JSON(http.StatusNotFound, gin.H{"error": ErrNotFound.Error()})
	} else if err == ErrInternalServerError {
		// 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServerError.Error()})
	} // etc...
}
