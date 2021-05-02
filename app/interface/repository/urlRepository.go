package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"url-shortener/app/domain"
	"url-shortener/app/usecase/repository"
)

type URLRepository interface {
	CreateURL(ctx context.Context, url domain.URL) (int, error)
	GetURLByShortCode(ctx context.Context, shortCode string) (domain.URL, error)
	DeleteURL(ctx context.Context, shortCode string) error
	ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error)
}

type urlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) repository.URLRepository {
	return &urlRepository{db: db}
}

func (u *urlRepository) CreateURL(ctx context.Context, url domain.URL) (int, error) {
	statement := `INSERT INTO url(short_code, full_url, has_expire_date, expire_date, deleted, number_of_hits)
		VALUES (?, ?, ?, ?, ?, ?)`
	insertedID, err := u.db.MustExecContext(
		ctx,
		statement,
		url.ShortCode,
		url.FullURL,
		url.HasExpireDate,
		url.ExpireDate,
		url.Deleted,
		url.NumberOfHits,
	).LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("unable to insert url %v", err)
	}
	return int(insertedID), nil

}

func (u *urlRepository) GetURLByShortCode(ctx context.Context, shortCode string) (domain.URL, error) {
	var url domain.URL
	err := u.db.SelectContext(ctx, &url, "SELECT * FROM url WHERE short_code = ?", shortCode)
	if err != nil {
		return domain.URL{}, err
	}
	return url, nil
}

func (u *urlRepository) DeleteURL(ctx context.Context, shortCode string) error {
	statement := `UPDATE url SET deleted = true WHERE short_code = ?`
	_, err := u.db.MustExecContext(ctx, statement, shortCode).RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (u *urlRepository) ListURL(ctx context.Context, shortCodeFilter string, fullURLKeywordFilter string) ([]domain.URL, error) {
	rows, err := u.db.QueryxContext(
		ctx,
		"SELECT * FROM url WHERE short_code LIKE '%?%' AND full_url LIKE '%?%'",
		shortCodeFilter,
		fullURLKeywordFilter,
	)
	if err != nil {
		return nil, err
	}
	var result []domain.URL
	for rows.Next() {
		var elem domain.URL
		err = rows.StructScan(&elem)
		if err != nil {
			return nil, err
		}
		result = append(result, elem)
	}
	return result, nil
}
