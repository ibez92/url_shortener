package entity

import (
	"context"
	"errors"
)

type Shorten struct {
	ID          uint64
	OrigianlURL string
	ShortURL    string
}

var (
	ErrShortenNotFound      = errors.New("the shorten not found")
	ErrShortenAlreadyExists = errors.New("the shorten already exists")

	ErrShortenOriginalURLBlank = errors.New("shorten original_url can't be blank")
)

type ShortenRepo interface {
	GetByShortURL(ctx context.Context, id uint64) (*Shorten, error)
	Save(ctx context.Context, shorten *Shorten) error
	Destroy(ctx context.Context, shorten *Shorten) error
}

func NewShorten(originalURL string) (*Shorten, error) {
	if originalURL == "" {
		return nil, ErrShortenOriginalURLBlank
	}

	return &Shorten{
		OrigianlURL: originalURL,
	}, nil
}
