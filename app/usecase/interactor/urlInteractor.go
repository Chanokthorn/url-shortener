package interactor

import (
	"context"
	"time"
	"url-shortener/app/domain"
	"url-shortener/app/usecase/repository"
)

type URLInteractor interface {
	CreateURL(ctx context.Context, shortCode, fullURL string, hasExpireDate bool, expireDate time.Time) error
	GetRedirectURL(ctx context.Context, shortCode string) (string, error)
	ListURL(ctx context.Context) ([]domain.URL, error)
	GetShortenedURL(ctx context.Context, fullURL string) (string, error)
}

type urlInteractor struct {
	urlRepository repository.URLRepository
}

func NewUrlInteractor(urlRepository repository.URLRepository) *urlInteractor {
	return &urlInteractor{urlRepository: urlRepository}
}
