package main

import (
	"log"

	"github.com/ibez92/url_shortener/internal/application"
)

func main() {
	app := application.NewApplication()

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start application: %s\n", err)
	}
}
