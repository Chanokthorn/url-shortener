package internal

type CreateURLRequest struct {
	FullURL    string  `json:"fullURL"`
	ExpireDate *string `json:"expireDate"`
}

type GetShortCodeFromURLRequest struct {
	FullURL string
}

type ListURL struct {
	ShortCodeFilter      string
	FullURLKeywordFilter string
}
