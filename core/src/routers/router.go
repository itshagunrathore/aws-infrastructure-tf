package routers

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/web"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories"
)

type RouterStruct struct {
	web.Router
}

func NewRoute(r RouterStruct) RouterStruct {
	return r
}

func (r *RouterStruct) GetRoute() {

	DB := NewDatabase(config.Read())
	jobDefinitionRepository := repositories.NewJobDefinitionRepository(DB)
	productService := services.NewProductService(productRepository)
	productHandlers := handlers.NewHttpHandler(productService)


}
