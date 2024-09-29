package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type server struct {
	e *echo.Echo
}

func NewServer() *server {
	e := echo.New()
	s := &server{
		e: e,
	}

	s.RegisterHTTP()

	return s
}

func (s *server) Start() error {
	return s.e.Start(":3000")
}

func (s *server) Stop(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s *server) RegisterHTTP() {
	s.registerAPIv1()
}

func (s *server) registerAPIv1() {
	apiV1 := s.e.Group("/api/v1")
	apiV1.POST("/shorten", s.CreateShorten)
	apiV1.GET("/shorten/:id", s.GetShorten)
	apiV1.PUT("/shorten/:id", s.UpdateShorten)
	apiV1.DELETE("/shorten/:id", s.DeleteShorten)
}

func (s *server) CreateShorten(c echo.Context) error {
	return c.JSON(http.StatusCreated, `{"status": "ok"}`)
}

func (s *server) GetShorten(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"status": "ok"}`)
}

func (s *server) UpdateShorten(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"status": "ok"}`)
}

func (s *server) DeleteShorten(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"status": "ok"}`)
}
