package server

import (
	"context"
	"net/http"

	"github.com/ibez92/url_shortener/internal/shorten"
	"github.com/ibez92/url_shortener/internal/shorten/command"
	"github.com/labstack/echo/v4"
)

type server struct {
	e          *echo.Echo
	shortenSvc *shorten.Service
}

func NewServer(shortenSvc *shorten.Service) *server {
	e := echo.New()
	s := &server{
		e:          e,
		shortenSvc: shortenSvc,
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
	s.e.GET("/:shortCode", s.Redirect)
	s.registerAPIv1()
}

func (s *server) registerAPIv1() {
	apiV1 := s.e.Group("/api/v1")
	apiV1.POST("/shorten", s.CreateShorten)
	apiV1.GET("/shorten/:shortURL", s.GetShorten)
	apiV1.PUT("/shorten/:shortURL", s.UpdateShorten)
	apiV1.DELETE("/shorten/:shortURL", s.DeleteShorten)
}

type CreateRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ID        uint64 `json:"id"`
	URL       string `json:"url"`
	ShortCode string `json:"short_code"`
}

func (s *server) Redirect(c echo.Context) error {
	shortCode := c.Param("shortCode")
	shorten, err := s.shortenSvc.Queries.GetByShortURL.Handle(c.Request().Context(), shortCode)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.Redirect(http.StatusTemporaryRedirect, shorten.OrigianlURL)
}

func (s *server) CreateShorten(c echo.Context) error {
	req := &CreateRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error("Invalid request")
		return c.NoContent(http.StatusBadRequest)
	}

	cmd := command.CreateShortenCmd{
		OrigianlURL: req.URL,
	}

	shorten, err := s.shortenSvc.Commands.Create.Handle(c.Request().Context(), cmd)
	if err != nil {
		c.Logger().Errorf("Something went wrong: %s", err)
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	return c.JSON(http.StatusCreated, &ShortenResponse{shorten.ID, shorten.OrigianlURL, shorten.ShortURL})
}

func (s *server) GetShorten(c echo.Context) error {
	shortURL := c.Param("shortURL")
	shorten, err := s.shortenSvc.Queries.GetByShortURL.Handle(c.Request().Context(), shortURL)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, &ShortenResponse{shorten.ID, shorten.OrigianlURL, shorten.ShortURL})
}

func (s *server) UpdateShorten(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"status": "ok"}`)
}

func (s *server) DeleteShorten(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"status": "ok"}`)
}
