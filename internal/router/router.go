package router

import (
	"shortener/internal/handler/getUrl"
	"shortener/internal/handler/pong"
	"shortener/internal/handler/short"
	"shortener/internal/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type BaseController struct {
	Router *gin.Engine
	Logger *logger.Logger
}

func New(driver *gin.Engine, logger *logger.Logger) *BaseController {
	return &BaseController{
		Router: driver,
		Logger: logger,
	}
}

func (c *BaseController) SetupRouter() *gin.RouterGroup {
	apiGroup := c.Router.Group("/api/v1")
	{
		apiGroup.GET("/links/:id", getUrl.GetUrlHandler)
		apiGroup.POST("/links", short.ShortHandler)
	}

	c.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	c.Router.GET("/ping", pong.PongHandler)

	return apiGroup
}
