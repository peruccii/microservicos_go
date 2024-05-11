package router

import (
	"github.com/gin-gonic/gin"
	"github.com/peruccii/microservicos_go/cmd/api/controller"
	"github.com/peruccii/microservicos_go/internal/repositorie"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1/category")
	inMemoryCategoryRepository := repositorie.NewInMemoryCategoryRepository()
	{
		v1.POST("/", func (ctx *gin.Context) {
			controller.CreateCategory(ctx, inMemoryCategoryRepository)
		})
	}
}