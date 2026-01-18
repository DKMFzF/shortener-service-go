package main

import (
	_ "shortener/docs"
	"shortener/internal/app"
)

// @title Shortener Service
// @version 1.0
// @description URL shortener service

// @host localhost:8080
// @BasePath /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app.New().Run()
}
