package repository

import (
	"context"
	"url-shortener/app/domain"
)

type URLRepository interface {
	CreateURL(ctx context.Context, url domain.URL) error
	GetURLByShortCode(ctx context.Context, shortCode string) (domain.URL, error)
	GetShortCodeByFullURL(ctx context.Context, fullURL string) (string, error)
	DeleteURL(ctx context.Context, shortCode string) error
	ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error)
	IncreaseURLNoOfHits(ctx context.Context, shortCode string) error
}
