package getUrl

import (
	"net/http"
	"shortener/internal/errors"

	"github.com/gin-gonic/gin"
)

// TODO: integration with Elastic Search

func GetUrlHandler(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.Error(errors.NewMethodNotAllowed(c.Request.Method))
		return
	}

	id := c.Param("id")
	if id == "" {
		c.Error(errors.NewValidationError("ID parameter is required", nil))
		return
	}

	c.String(http.StatusOK, "https://practicum.yandex.ru/")
}
