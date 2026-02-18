package app

import (
	"context"
	"errors"
	"net/http"
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
	Controller *router.BaseController
	Config     *config.Config
	Logger     *logger.Logger
}

func New() *App {
	logger := logger.New(
		flags.ParseFlagIsLogsInFile(),  // --writingLogs
		flags.ParseFlagIsDebugLogger(), // --debugLogger
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
	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-sigCtx.Done()
	app.Logger.Infof("shutdown signal received: %v", context.Cause(sigCtx))

	timeout := app.Config.GracefulShutdownTime
	if timeout <= 0 {
		timeout = 5 * time.Second
	}

	shutdownCause := errors.New("graceful shutdown timeout exceeded")
	shutdownCtx, cancel := context.WithTimeoutCause(context.Background(), timeout, shutdownCause)
	defer cancel()

	if err := app.Server.Shutdown(shutdownCtx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(context.Cause(shutdownCtx), shutdownCause) {
			app.Logger.Errorf("shutdown timed out after %s: %v", timeout, err)
			return
		}
		app.Logger.Errorf("shutdown failed: %v", err)
		return
	}

	app.Logger.Infof("server shutdown completed")
}
