package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
