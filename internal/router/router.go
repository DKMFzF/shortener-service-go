package router

import (
	"shortener/internal/handler/getUrl"
	"shortener/internal/handler/pong"
	"shortener/internal/handler/short"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(router *gin.Engine) *gin.RouterGroup {
	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/:id", getUrl.GetUrlHandler)
		apiGroup.POST("/", short.ShortHandler)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", pong.PongHandler)

	return apiGroup
}
