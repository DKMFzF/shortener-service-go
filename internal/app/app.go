package app

import (
	"net/http"
	"shortener/internal/router"
)

type App struct {
	Server http.Server
}

func NewApp() *App {
	return &App{
		Server: http.Server{
			Addr:    ":8080",
			Handler: router.Router(),
		},
	}
}

func (app *App) Run() error {
	return app.Server.ListenAndServe()
}
