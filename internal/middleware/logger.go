package middleware

import (
	"net/http"
	"shortener/internal/logger"
)

func LoggerMiddleware(next http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Infof("%s : %v", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
