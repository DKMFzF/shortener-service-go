package router

import (
	"net/http"
	"shortener/internal/handler"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(`GET /ping`, http.HandlerFunc(handler.PongHandler))

	return mux
}
