package web

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	//DB *db.PostgresDB
}

func NewRouter() *Router {
	r := gin.Default()

	return &Router{Engine: r}
}
