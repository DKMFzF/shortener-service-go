package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	config "shortener/internal/configs/apps/shortenerConfig"
	"shortener/internal/configs/flags"
	"shortener/internal/logger"
	"shortener/internal/middleware"
	"shortener/internal/router"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type App struct {
	Server     http.Server
	Context    context.Context
	Cancel     context.CancelCauseFunc
	Controller *router.BaseController
	Config     *config.Config
	Logger     *logger.Logger
}

func New() *App {
	logger := logger.New(
		flags.ParseFlagIsLogsInFile(), // check --writingLogs
	)
	return &App{
		Controller: router.New(gin.New(), logger),
		Config:     config.New(),
		Logger:     logger,
	}
}

func (app *App) Run() {
	app.StartServer()
	app.GracefulShutdown()
}

func (app *App) _Bootstrap() *App {
	app.Server = http.Server{
		Addr:    ":" + app.Config.Addr,
		Handler: app.Controller.Router,
	}

	app.Controller.Router.Use(
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
		middleware.MiddlewareLogger(*app.Logger),
		middleware.NewErrorHandler(app.Logger).Handle(),
	)
	app.Controller.SetupRouter()

	return app
}

func (app *App) StartServer() {
	app._Bootstrap()
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
