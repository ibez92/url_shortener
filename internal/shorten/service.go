package shorten

import (
	"github.com/ibez92/url_shortener/internal/shorten/command"
	"github.com/ibez92/url_shortener/internal/shorten/query"
)

type Queries struct {
	GetByShortURL *query.GetByShortURLHandler
}

type Commands struct {
	Create  *command.CreateShortenHandler
	Update  *command.UpdateShortenHandler
	Destroy *command.DestroyShortenHandler
}

type Service struct {
	Queries  Queries
	Commands Commands
}
