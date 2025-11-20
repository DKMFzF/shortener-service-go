package pong

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PongHandler(c *gin.Context) {
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, "pong")
}
