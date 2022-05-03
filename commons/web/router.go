package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
)

type Router struct {
	engine *gin.Engine
	DB *db.PostgresDB
}

func NewRouter() *Router {
	r := gin.Default()

	return &Router{engine: r}
}
