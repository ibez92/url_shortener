package repository

import (
	"context"
	"sync"

	"github.com/ibez92/url_shortener/internal/entity"
	"github.com/ibez92/url_shortener/internal/pkg/shorturl"
)

type ShortenMemoryRepo struct {
	mu              sync.RWMutex
	shortens        map[uint64]*entity.Shorten
	autoIncrementID uint64
}

func NewShortenMemoryRepo() *ShortenMemoryRepo {
	return &ShortenMemoryRepo{
		shortens:        make(map[uint64]*entity.Shorten),
		autoIncrementID: 1,
	}
}

func (r *ShortenMemoryRepo) GetByShortURL(ctx context.Context, shortURL string) (*entity.Shorten, error) {
	id := shorturl.IdByShortURL(shortURL) + 1

	r.mu.RLock()
	defer r.mu.RUnlock()

	if shorten, ok := r.shortens[id]; ok {
		return shorten, nil
	}

	return nil, entity.ErrShortenNotFound
}

func (r *ShortenMemoryRepo) Create(ctx context.Context, shorten *entity.Shorten) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.shortens[shorten.ID]; ok {
		return entity.ErrShortenAlreadyExists
	}

	shorten.ID = r.autoIncrementID
	r.autoIncrementID += 1
	shorten.ShortURL = shorturl.ShortURLByID(shorten.ID - 1)
	r.shortens[shorten.ID] = shorten

	return nil
}

func (r *ShortenMemoryRepo) Update(ctx context.Context, shorten *entity.Shorten) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.shortens[shorten.ID]; !ok {
		return entity.ErrShortenNotFound
	}

	r.shortens[shorten.ID] = shorten

	return nil
}

func (r *ShortenMemoryRepo) Destroy(ctx context.Context, shorten *entity.Shorten) error {
	r.mu.Lock()
	delete(r.shortens, shorten.ID)
	r.mu.Unlock()

	return nil
}
