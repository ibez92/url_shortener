package command

import (
	"context"

	"github.com/ibez92/url_shortener/internal/entity"
)

type ShortenUpdateRepo interface {
	GetByShortURL(ctx context.Context, shortCode string) (*entity.Shorten, error)
	Update(ctx context.Context, shorten *entity.Shorten) error
}

type UpdateShortenHandler struct {
	shortenUpdateRepo ShortenUpdateRepo
}

type UpdateShortenCmd struct {
	ShortCode   string
	OrigianlURL string
}

func NewUpdateShortenHandler(shortenUpdateRepo ShortenUpdateRepo) *UpdateShortenHandler {
	return &UpdateShortenHandler{
		shortenUpdateRepo: shortenUpdateRepo,
	}
}

func (h *UpdateShortenHandler) Handle(ctx context.Context, cmd UpdateShortenCmd) (*entity.Shorten, error) {
	shorten, err := h.shortenUpdateRepo.GetByShortURL(ctx, cmd.ShortCode)
	if err != nil {
		return nil, err
	}

	shorten.OrigianlURL = cmd.OrigianlURL
	if err := h.shortenUpdateRepo.Update(ctx, shorten); err != nil {
		return nil, err
	}

	return shorten, nil
}
