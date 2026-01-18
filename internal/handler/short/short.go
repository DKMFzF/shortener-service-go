package short

import (
	"net/http"
	"net/url"
	"shortener/internal/errors"
	"shortener/internal/service/resizeUrl"

	"github.com/gin-gonic/gin"
)

func ShortHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.Error(errors.NewMethodNotAllowed(c.Request.Method))
		return
	}

	contentType := c.GetHeader("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		c.Error(errors.NewBadRequest(
			"Invalid Content-Type. Expected text/plain; charset=utf-8",
			nil,
		))
		return
	}

	body, err := c.GetRawData()
	if err != nil {
		c.Error(errors.NewBadRequest("Failed to read request body", err))
		return
	}

	originUrl := string(body)
	if originUrl == "" {
		c.Error(errors.NewValidationError("URL is required in request body", nil))
		return
	}

	if _, err := url.ParseRequestURI(originUrl); err != nil {
		c.Error(errors.NewValidationError("Invalid URL format", err))
		return
	}

	shortCode, err := resizeUrl.ResizeUrl(originUrl)
	if err != nil {
		c.Error(errors.NewInternalError(err))
		return
	}

	resStr := "http://localhost:8080/" + shortCode
	c.Data(http.StatusCreated, "text/plain; charset=utf-8", []byte(resStr))
}
