package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"shortener/internal/app"
)

func main() {
	if err := app.NewApp().Run(); errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("%v", err)
		os.Exit(1)
	}
}
