package router

import (
	"net/http"

	"shortener/internal/handler/getUrl"
	"shortener/internal/handler/pong"
	"shortener/internal/handler/short"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(`GET /ping`, http.HandlerFunc(pong.PongHandler))
	mux.Handle(`POST /`, http.HandlerFunc(short.ShortHandler))
	mux.Handle(`GET /{id}`, http.HandlerFunc(getUrl.GetUrlHandler))

	return mux
}
