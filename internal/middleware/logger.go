package middleware

import (
	"shortener/internal/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MiddlewareLogger(logger logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Infof(
			"%s", "Request "+
				ctx.Request.RequestURI+
				" Response Code "+strconv.Itoa(ctx.Writer.Status()),
		)
		ctx.Next()
	}
}
