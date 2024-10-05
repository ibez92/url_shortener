package application

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ibez92/url_shortener/internal/repository"
	"github.com/ibez92/url_shortener/internal/server"
	"github.com/ibez92/url_shortener/internal/shorten"
	"github.com/ibez92/url_shortener/internal/shorten/command"
	"github.com/ibez92/url_shortener/internal/shorten/query"
)

type Server interface {
	Start() error
	Stop(context.Context) error
}

type Application struct {
	ctx    context.Context
	server Server
	stop   chan struct{}
	sigs   chan os.Signal
}

func NewApplication() *Application {
	shortenRepo := repository.NewShortenMemoryRepo()
	shortenSvc := &shorten.Service{
		Queries: shorten.Queries{
			GetByShortURL: query.NewGetByShortURLHandler(shortenRepo),
		},
		Commands: shorten.Commands{
			Create:  command.NewCreateShortenHandler(shortenRepo),
			Update:  command.NewUpdateShortenHandler(shortenRepo),
			Destroy: command.NewDestroyShortenHandler(shortenRepo),
		},
	}

	server := server.NewServer(shortenSvc)

	return &Application{
		ctx:    context.Background(),
		stop:   make(chan struct{}),
		sigs:   make(chan os.Signal),
		server: server,
	}
}

func (a *Application) Run() error {
	signal.Notify(a.sigs, syscall.SIGTERM, syscall.SIGINT)

	go a.sigWait()
	go a.server.Start()

	log.Default().Println("Application started")

	<-a.stop
	a.Stop()

	log.Default().Println("Application stopped")

	return nil
}

func (a *Application) Stop() error {
	if err := a.server.Stop(a.ctx); err != nil {
		log.Fatal()
	}

	return nil
}

func (a *Application) sigWait() {
	<-a.sigs
	a.stop <- struct{}{}
}
