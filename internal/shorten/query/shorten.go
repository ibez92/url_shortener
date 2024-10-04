package query

import (
	"context"

	"github.com/ibez92/url_shortener/internal/entity"
)

type ShortenRead interface {
	GetByShortURL(ctx context.Context, shortURL string) (*entity.Shorten, error)
}

type GetByShortURLHandler struct {
	shortenRead ShortenRead
}

func NewGetByShortURLHandler(shortenRead ShortenRead) *GetByShortURLHandler {
	return &GetByShortURLHandler{
		shortenRead: shortenRead,
	}
}

func (h *GetByShortURLHandler) Handle(ctx context.Context, shortURL string) (*entity.Shorten, error) {
	return h.shortenRead.GetByShortURL(ctx, shortURL)
}
