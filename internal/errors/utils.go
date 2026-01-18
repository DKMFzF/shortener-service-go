package errors

import (
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	if appErr, ok := err.(*AppError); ok {
		c.Error(appErr)
	} else {
		c.Error(NewInternalError(err))
	}
}

func ResponseWithError(c *gin.Context, appErr *AppError) {
	c.Error(appErr)
	c.Abort()
}
