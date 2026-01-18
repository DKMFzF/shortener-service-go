package middleware

import (
	"shortener/internal/logger"

	"github.com/gin-gonic/gin"
)

func MiddlewareError(logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Infof("Panic revocerd: %v", r)

			}
		}()

		ctx.Next()
	}
}
