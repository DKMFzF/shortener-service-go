package middleware

import (
	"shortener/internal/errors"
	"shortener/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorHandler struct {
	logger *logger.Logger
}

func NewErrorHandler(logger *logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h *ErrorHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Если нет ошибок - пропускаем
		if len(c.Errors) == 0 {
			return
		}

		// Берем последнюю ошибку
		err := c.Errors.Last().Err

		var appErr *errors.AppError

		switch e := err.(type) {
		case *errors.AppError:
			appErr = e
		case validator.ValidationErrors:
			// Обработка ошибок валидации
			validationErrors := make(map[string]string)
			for _, ve := range e {
				validationErrors[ve.Field()] = ve.Tag()
			}
			appErr = errors.NewValidationError("Validation failed", e)
			appErr.Meta = validationErrors
		default:
			// Логируем внутренние ошибки
			h.logger.Warnf("Internal error", "error", e)
			appErr = errors.NewInternalError(e)
		}

		// Отправляем ошибку клиенту
		c.JSON(appErr.Status, errors.ErrorResponse{
			Errors: []errors.AppError{*appErr},
		})
	}
}
