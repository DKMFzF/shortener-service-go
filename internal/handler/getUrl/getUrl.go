package getUrl

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUrlHandler(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.String(http.StatusBadRequest, "Bad Method. Use GET")
		return
	}

	if !strings.Contains(c.GetHeader("Content-Type"), "text/plain; charset=utf-8") {
		c.String(http.StatusBadRequest, "Bad Header type")
		return
	}

	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "Not param id")
		return
	}

	// TODO: write service logic for get url

	c.String(http.StatusOK, "https://practicum.yandex.ru/")
}
