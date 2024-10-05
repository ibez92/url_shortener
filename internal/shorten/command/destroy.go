package command

import (
	"context"

	"github.com/ibez92/url_shortener/internal/entity"
)

type ShortenDestroyRepo interface {
	GetByShortURL(ctx context.Context, shortCode string) (*entity.Shorten, error)
	Destroy(ctx context.Context, shorten *entity.Shorten) error
}

type DestroyShortenHandler struct {
	shortenDestroyRepo ShortenDestroyRepo
}

type DestroyShortenCmd struct {
	ShortCode string
}

func NewDestroyShortenHandler(shortenDestroyRepo ShortenDestroyRepo) *DestroyShortenHandler {
	return &DestroyShortenHandler{
		shortenDestroyRepo: shortenDestroyRepo,
	}
}

func (h *DestroyShortenHandler) Handle(ctx context.Context, cmd DestroyShortenCmd) error {
	shorten, err := h.shortenDestroyRepo.GetByShortURL(ctx, cmd.ShortCode)
	if err != nil {
		return err
	}

	if err := h.shortenDestroyRepo.Destroy(ctx, shorten); err != nil {
		return err
	}

	return nil
}
