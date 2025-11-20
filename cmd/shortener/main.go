package main

import (
	"shortener/internal/app"
	"shortener/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := app.New()
	app.Router.Use(gin.Recovery())
	router.SetupRouter(app.Router)
	app.Run()
}
