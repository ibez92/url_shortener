package application

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
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

func NewApplication(s Server) *Application {
	return &Application{
		ctx:    context.Background(),
		stop:   make(chan struct{}),
		sigs:   make(chan os.Signal),
		server: s,
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
