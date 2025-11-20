package short

import (
	"net/http"
	"net/url"
	"shortener/internal/service/resizeUrl"

	"strings"

	"github.com/gin-gonic/gin"
)

func ShortHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.String(http.StatusBadRequest, "")
		return
	}

	if !strings.Contains(c.GetHeader("Content-Type"), "text/plain; charset=utf-8") {
		c.String(http.StatusBadRequest, "Bad Contet-Type")
		return
	}

	body, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, "Not URL")
		return
	}

	originUrl := string(body)
	if originUrl == "" {
		c.String(http.StatusBadRequest, "Not Url in body")
		return
	}

	if _, err := url.ParseRequestURI(originUrl); err != nil {
		c.String(http.StatusBadRequest, "It's not URL")
		return
	}

	shortCode := resizeUrl.ResizeUrl(originUrl)
	resStr := "http://localhost:8080/" + shortCode

	c.Data(http.StatusCreated, "text/plain; charset=utf-8", []byte(resStr))
}
