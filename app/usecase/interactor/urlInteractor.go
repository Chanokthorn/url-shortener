package interactor

import (
	"context"
	"fmt"
	"github.com/teris-io/shortid"
	"net/http"
	"regexp"
	"time"
	"url-shortener/app/domain"
	"url-shortener/app/usecase/repository"
)

type URLInteractor interface {
	CreateURL(ctx context.Context, fullURL string, hasExpireDate bool, expireDate time.Time) (string, error)
	GetRedirectURL(ctx context.Context, shortCode string) (string, int, error)
	ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error)
	GetShortCode(ctx context.Context, fullURL string) (string, error)
	DeleteURL(ctx context.Context, shortCode string) error
}

type urlInteractor struct {
	urlRepository repository.URLRepository
	blacklist     []string
}

func NewUrlInteractor(urlRepository repository.URLRepository, blacklist []string) URLInteractor {
	return &urlInteractor{
		urlRepository: urlRepository,
		blacklist:     blacklist,
	}
}

func (u *urlInteractor) CreateURL(ctx context.Context, fullURL string, hasExpireDate bool, expireDate time.Time) (string, error) {
	for _, pattern := range u.blacklist {
		matched, err := regexp.MatchString(pattern, fullURL)
		if err != nil {
			return "", fmt.Errorf("unable to perform regex matching: %v", err)
		}
		if matched {
			return "", fmt.Errorf("full url had been blacklisted")
		}
	}
	shortCode, err := shortid.Generate()
	if err != nil {
		return "", err
	}
	url := domain.URL{
		ShortCode:     shortCode,
		FullURL:       fullURL,
		HasExpireDate: hasExpireDate,
		ExpireDate:    expireDate,
		Deleted:       false,
		NumberOfHits:  0,
	}
	err = u.urlRepository.CreateURL(ctx, url)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}

func (u *urlInteractor) GetRedirectURL(ctx context.Context, shortCode string) (string, int, error) {
	url, err := u.urlRepository.GetURLByShortCode(ctx, shortCode)
	if err != nil {
		return "", http.StatusNotFound, err
	}
	if url.Deleted {
		return "", http.StatusGone, nil
	}
	if url.HasExpireDate {
		if time.Now().After(url.ExpireDate) {
			return "", http.StatusGone, nil
		}
	}
	if url.IsEmpty() {
		return "", http.StatusNotFound, nil
	}
	err = u.urlRepository.IncreaseURLNoOfHits(ctx, shortCode)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return url.FullURL, http.StatusFound, nil
}

func (u *urlInteractor) ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error) {
	return u.urlRepository.ListURL(ctx, shortCodeFilter, fullURLKeywordFilter)
}

func (u *urlInteractor) GetShortCode(ctx context.Context, fullURL string) (string, error) {
	return u.urlRepository.GetShortCodeByFullURL(ctx, fullURL)
}

func (u *urlInteractor) DeleteURL(ctx context.Context, shortCode string) error {
	return u.urlRepository.DeleteURL(ctx, shortCode)
}
