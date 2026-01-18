package getUrl

import (
	"net/http"
	"shortener/internal/errors"

	"github.com/gin-gonic/gin"
)

func GetUrlHandler(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.Error(errors.NewMethodNotAllowed(c.Request.Method))
		return
	}

	contentType := c.GetHeader("Content-Type")
	if contentType != "" && contentType != "text/plain; charset=utf-8" {
		c.Error(errors.NewBadRequest(
			"Invalid Content-Type. Expected text/plain; charset=utf-8 or empty",
			nil,
		))
		return
	}

	id := c.Param("id")
	if id == "" {
		c.Error(errors.NewValidationError("ID parameter is required", nil))
		return
	}

	c.String(http.StatusOK, "https://practicum.yandex.ru/")
}
