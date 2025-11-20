package router

import (
	"shortener/internal/handler/getUrl"
	"shortener/internal/handler/pong"
	"shortener/internal/handler/short"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.RouterGroup {
	apiGroup := router.Group("/api")

	apiGroup.GET("/ping", pong.PongHandler)
	apiGroup.GET("/:id", getUrl.GetUrlHandler)
	apiGroup.POST("/", short.ShortHandler)

	return apiGroup
}
