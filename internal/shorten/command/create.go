package command

import (
	"context"

	"github.com/ibez92/url_shortener/internal/entity"
)

type ShortenWrite interface {
	Create(ctx context.Context, shorten *entity.Shorten) error
}

type CreateShortenHandler struct {
	shortenWrite ShortenWrite
}

type CreateShortenCmd struct {
	OrigianlURL string
}

func NewCreateShortenHandler(shortenWrite ShortenWrite) *CreateShortenHandler {
	return &CreateShortenHandler{
		shortenWrite: shortenWrite,
	}
}

func (h *CreateShortenHandler) Handle(ctx context.Context, cmd CreateShortenCmd) (*entity.Shorten, error) {
	shorten, err := entity.NewShorten(cmd.OrigianlURL)
	if err != nil {
		return nil, err
	}

	if err := h.shortenWrite.Create(ctx, shorten); err != nil {
		return nil, err
	}

	return shorten, nil
}
