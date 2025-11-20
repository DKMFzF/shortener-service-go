package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	Server  http.Server
	Context context.Context
	Cancel  context.CancelCauseFunc
	Router  *gin.Engine
}

func New() *App {
	return &App{
		Router: gin.New(),
	}
}

func (app *App) Run() {
	app.StartServer()
	app.GracefulShutdown()
}

func (app *App) StartServer() {
	app.Server = http.Server{
		Addr:    ":8080",
		Handler: app.Router,
	}

	go func() {
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (app *App) GracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
