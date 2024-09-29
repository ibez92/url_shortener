package main

import (
	"github.com/ibez92/url_shortener/internal/application"
	"github.com/ibez92/url_shortener/internal/server"
)

func main() {
	server := server.NewServer()
	app := application.NewApplication(server)

	app.Run()
}
