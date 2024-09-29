package application

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Start() error
	Stop() error
}

type Application struct {
	server Server
	stop   chan struct{}
	sigs   chan os.Signal
}

func NewApplication(s Server) *Application {
	return &Application{
		stop:   make(chan struct{}),
		sigs:   make(chan os.Signal),
		server: s,
	}
}

func (a *Application) Run() error {
	signal.Notify(a.sigs, syscall.SIGTERM, syscall.SIGINT)

	go a.sigWait()
	go a.server.Start()

	fmt.Println("Application started")

	<-a.stop
	a.Stop()

	fmt.Println("Application stoped")

	return nil
}

func (a *Application) Stop() error {
	a.server.Stop()

	return nil
}

func (a *Application) sigWait() {
	<-a.sigs
	a.stop <- struct{}{}
}
