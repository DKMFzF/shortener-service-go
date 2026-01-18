package pong

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Pong godoc
// @Summary Проверка доступности сервиса
// @Description Возвращает pong
// @Tags system
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PongHandler(c *gin.Context) {
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, "pong")
}
