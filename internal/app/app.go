package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"shortener/internal/configs"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	Server  http.Server
	Context context.Context
	Cancel  context.CancelCauseFunc
	Router  *gin.Engine
	Config  *configs.Config
}

func New() *App {
	return &App{
		Router: gin.New(),
		Config: configs.New(),
	}
}

func (app *App) Run() {
	app.StartServer()
	app.GracefulShutdown()
}

func (app *App) StartServer() {
	app.Server = http.Server{
		Addr:    ":" + app.Config.Addr,
		Handler: app.Router,
	}

	go func() {
		fmt.Printf("\nServer start on port: %s\r\n\n", app.Config.Addr)
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
