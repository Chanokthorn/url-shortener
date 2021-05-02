package repository

import (
	"context"
	"url-shortener/app/domain"
)

type URLRepository interface {
	CreateURL(ctx context.Context, url domain.URL) (int, error)
	GetURLByShortCode(ctx context.Context, shortCode string) (domain.URL, error)
	DeleteURL(ctx context.Context, shortCode string) error
	ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error)
}
