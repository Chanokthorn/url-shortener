package domain

import "time"

type URL struct {
	ShortCode     string
	FullURL       string
	HasExpireDate bool
	ExpireDate    time.Time
	Deleted       bool
	NumberOfHits  int
}

func (u *URL) IsEmpty() bool {
	return *u == URL{}
}
