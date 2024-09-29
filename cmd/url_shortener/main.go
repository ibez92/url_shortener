package main

import (
	"fmt"

	"github.com/ibez92/url_shortener/internal/application"
	"github.com/ibez92/url_shortener/internal/server"
)

func main() {
	server := server.NewServer()
	_ = application.NewApplication(server)

	fmt.Println("Url shortner init commit")
}
