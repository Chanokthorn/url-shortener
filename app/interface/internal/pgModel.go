package internal

import (
	"database/sql"
	"url-shortener/app/domain"
)

type URLPG struct {
	ShortCode     sql.NullString `db:"short_code"`
	FullURL       sql.NullString `db:"full_url"`
	HasExpireDate sql.NullBool   `db:"has_expire_date"`
	ExpireDate    sql.NullTime   `db:"expire_date"`
	Deleted       sql.NullBool   `db:"deleted"`
	NumberOfHits  sql.NullInt64  `db:"number_of_hits"`
}

func (u *URLPG) getURL() *domain.URL {
	return &domain.URL{
		ShortCode:     u.ShortCode.String,
		FullURL:       u.FullURL.String,
		HasExpireDate: u.HasExpireDate.Bool,
		ExpireDate:    u.ExpireDate.Time,
		Deleted:       u.Deleted.Bool,
		NumberOfHits:  int(u.NumberOfHits.Int64),
	}
}
