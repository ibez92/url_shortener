package server

import "github.com/labstack/echo/v4"

type server struct {
	e *echo.Echo
}

func NewServer() *server {
	e := echo.New()

	return &server{
		e: e,
	}
}

func (s *server) Start() error {
	return nil
}

func (s *server) Stop() error {
	return nil
}
