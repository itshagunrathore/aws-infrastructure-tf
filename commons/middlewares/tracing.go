package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Tracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("baas-trace-id", uuid.New())
	}
}
