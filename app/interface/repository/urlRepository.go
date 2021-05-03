package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"url-shortener/app/domain"
	"url-shortener/app/interface/internal"
	"url-shortener/app/usecase/repository"
)

type URLRepository interface {
	CreateURL(ctx context.Context, url domain.URL) error
	GetURLByShortCode(ctx context.Context, shortCode string) (domain.URL, error)
	GetShortCodeByFullURL(ctx context.Context, fullURL string) (string, error)
	DeleteURL(ctx context.Context, shortCode string) error
	ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error)
}

type urlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) repository.URLRepository {
	return &urlRepository{db: db}
}

func (u *urlRepository) CreateURL(ctx context.Context, url domain.URL) error {
	statement := `INSERT INTO url (short_code, full_url, has_expire_date, expire_date, deleted, number_of_hits)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := u.db.MustExecContext(
		ctx,
		statement,
		url.ShortCode,
		url.FullURL,
		url.HasExpireDate,
		url.ExpireDate,
		url.Deleted,
		url.NumberOfHits,
	).RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to insert url %v", err)
	}
	return nil
}

func (u *urlRepository) GetURLByShortCode(ctx context.Context, shortCode string) (domain.URL, error) {
	var url internal.URLPG
	err := u.db.SelectContext(ctx, &url, "SELECT * FROM url WHERE short_code = ?", shortCode)
	if err != nil {
		return domain.URL{}, err
	}
	return *url.ToURL(), nil
}

func (u *urlRepository) GetShortCodeByFullURL(ctx context.Context, fullURL string) (string, error) {
	var res []string
	err := u.db.SelectContext(ctx, &res, "SELECT short_code FROM url WHERE full_url = $1", fullURL)
	if err != nil {
		return "", err
	}
	if len(res) != 1 {
		return "", errors.New("not found")
	}
	return res[0], nil
}

func (u *urlRepository) DeleteURL(ctx context.Context, shortCode string) error {
	statement := `UPDATE url SET deleted = true WHERE short_code = $1`
	res, err := u.db.MustExecContext(ctx, statement, shortCode).RowsAffected()
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("url not found")
	}
	return nil
}

func (u *urlRepository) ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error) {
	rows, err := u.db.QueryxContext(
		ctx,
		"SELECT * FROM url WHERE short_code LIKE $1 AND full_url LIKE $2",
		"%"+shortCodeFilter+"%",
		"%"+fullURLKeywordFilter+"%",
	)
	if err != nil {
		return nil, err
	}
	var result []domain.URL
	for rows.Next() {
		var elem internal.URLPG
		err = rows.StructScan(&elem)
		if err != nil {
			return nil, err
		}
		result = append(result, *elem.ToURL())
	}
	return result, nil
}
