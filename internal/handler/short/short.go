package short

import (
	"net/http"
	"shortener/internal/errors"
	"shortener/internal/service/resizeUrl"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type MetaInfo struct {
	Title       string `json:"title" valid:"optional"`
	Description string `json:"description" valid:"optional"`
}

type UrlToShort struct {
	Url string `json:"url" valid:"url"`
}

type Request struct {
	Meta *MetaInfo   `json:"meta" valid:"required"`
	Data *UrlToShort `json:"data" valid:"required"`
}

type Response struct {
	Data *UrlToShort `json:"data" valid:"required"`
}

// Short godoc
// @Summary сокращение url
// @Description Возвращает сокращенный url и meta информацию о нём
// @Accept json
// @Produce json
// @Tags system
// @Param request body Request true "Параметры запроса"
// @Success 200 {object} Response "Успешный ответ"
// @Router /api/v1/links/ [post]
func ShortHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.Error(errors.NewMethodNotAllowed(c.Request.Method))
		return
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.NewBadRequest(
			"Invalid JSON format or structure",
			err,
		))
		return
	}

	if _, err := valid.ValidateStruct(req); err != nil {
		c.Error(errors.NewBadRequest(
			"Valid faild: "+err.Error(),
			err,
		))
		return
	}

	shortCode, err := resizeUrl.ResizeUrl(req.Data.Url)
	if err != nil {
		c.Error(errors.NewInternalError(err))
		return
	}

	resData := Response{
		Data: &UrlToShort{
			Url: "http://localhost:8080/" + shortCode,
		},
	}
	c.JSON(http.StatusCreated, resData)
}
