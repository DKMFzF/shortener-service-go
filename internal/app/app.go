package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	config "shortener/internal/configs/apps/shortenerConfig"
	"shortener/internal/logger"
	"syscall"
	"time"
)

type App struct {
	Server  http.Server
	Context context.Context
	Cancel  context.CancelCauseFunc
	Router  *gin.Engine
	Config  *config.Config
	Logger  *logger.Logger
}

func New() *App {
	return &App{
		Router: gin.New(),
		Config: config.New(),
		Logger: logger.New(),
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
		app.Logger.Infof("\nServer start on port: %s\r\n\n", app.Config.Addr)
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatalf("panic: %v", err)
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
		app.Logger.Fatalf("panic %v", err)
		panic(err)
	}
}
