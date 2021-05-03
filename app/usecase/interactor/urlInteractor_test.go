package interactor

import (
	"context"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
	"url-shortener/app/domain"
	"url-shortener/app/usecase/repository"
	"url-shortener/app/usecase/repository/mocks"
)

func Test_urlInteractor_CreateURL(t *testing.T) {
	type fields struct {
		urlRepository repository.URLRepository
		blacklist     []string
	}
	type args struct {
		ctx           context.Context
		fullURL       string
		hasExpireDate bool
		expireDate    time.Time
	}
	urlRepo := new(mocks.URLRepository)
	urlRepo.On("CreateURL", mock.Anything, mock.Anything).Return(nil)
	blacklist := []string{
		"^.*troll.*",
		"^.*forbidden.*",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful creation",
			fields: fields{
				urlRepository: urlRepo,
				blacklist:     blacklist,
			},
			args: args{
				ctx:           nil,
				fullURL:       "https://www.google.com/",
				hasExpireDate: false,
				expireDate:    time.Time{},
			},
			wantErr: false,
		},
		{
			name: "forbidden creation 1	",
			fields: fields{
				urlRepository: urlRepo,
				blacklist:     blacklist,
			},
			args: args{
				ctx:           nil,
				fullURL:       "https://www.sometroller.com/",
				hasExpireDate: false,
				expireDate:    time.Time{},
			},
			wantErr: true,
		},
		{
			name: "forbidden creation 2	",
			fields: fields{
				urlRepository: urlRepo,
				blacklist:     blacklist,
			},
			args: args{
				ctx:           nil,
				fullURL:       "https://www.justaforbiddensite.com/",
				hasExpireDate: false,
				expireDate:    time.Time{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &urlInteractor{
				urlRepository: tt.fields.urlRepository,
				blacklist:     tt.fields.blacklist,
			}
			_, err := u.CreateURL(tt.args.ctx, tt.args.fullURL, tt.args.hasExpireDate, tt.args.expireDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_urlInteractor_GetRedirectURL(t *testing.T) {
	currTime := time.Now()
	type fields struct {
		urlRepository repository.URLRepository
	}
	type args struct {
		ctx       context.Context
		shortCode string
	}
	urlRepo1 := new(mocks.URLRepository)
	url1 := domain.URL{
		ShortCode:     "test1",
		FullURL:       "test1-full-url",
		HasExpireDate: false,
		ExpireDate:    time.Time{},
		Deleted:       false,
		NumberOfHits:  0,
	}
	urlRepo1.On("GetURLByShortCode", mock.Anything, mock.Anything).Return(url1, nil)
	urlRepo1.On("IncreaseURLNoOfHits", mock.Anything, mock.Anything).Return(nil)
	urlRepo2 := new(mocks.URLRepository)
	url2 := domain.URL{
		ShortCode:     "test2",
		FullURL:       "test2-full-url",
		HasExpireDate: false,
		ExpireDate:    time.Time{},
		Deleted:       true,
		NumberOfHits:  0,
	}
	urlRepo2.On("GetURLByShortCode", mock.Anything, mock.Anything).Return(url2, nil)
	urlRepo2.On("IncreaseURLNoOfHits", mock.Anything, mock.Anything).Return(nil)
	urlRepo3 := new(mocks.URLRepository)
	url3 := domain.URL{
		ShortCode:     "test3",
		FullURL:       "test3-full-url",
		HasExpireDate: true,
		ExpireDate:    currTime.Add(time.Hour),
		Deleted:       false,
		NumberOfHits:  0,
	}
	urlRepo3.On("GetURLByShortCode", mock.Anything, mock.Anything).Return(url3, nil)
	urlRepo3.On("IncreaseURLNoOfHits", mock.Anything, mock.Anything).Return(nil)
	urlRepo4 := new(mocks.URLRepository)
	url4 := domain.URL{
		ShortCode:     "test4",
		FullURL:       "test4-full-url",
		HasExpireDate: true,
		ExpireDate:    currTime.Add(-time.Hour),
		Deleted:       false,
		NumberOfHits:  0,
	}
	urlRepo4.On("GetURLByShortCode", mock.Anything, mock.Anything).Return(url4, nil)
	urlRepo4.On("IncreaseURLNoOfHits", mock.Anything, mock.Anything).Return(nil)
	urlRepo5 := new(mocks.URLRepository)
	urlRepo5.On("GetURLByShortCode", mock.Anything, mock.Anything).Return(domain.URL{}, nil)
	urlRepo5.On("IncreaseURLNoOfHits", mock.Anything, mock.Anything).Return(nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		{
			name:   "valid url",
			fields: fields{urlRepository: urlRepo1},
			args:   args{ctx: nil, shortCode: ""},
			want:   "test1-full-url",
			want1:  http.StatusFound,
		},
		{
			name:   "deleted url",
			fields: fields{urlRepository: urlRepo2},
			args:   args{ctx: nil, shortCode: ""},
			want:   "",
			want1:  http.StatusGone,
		},
		{
			name:   "unexpired url",
			fields: fields{urlRepository: urlRepo3},
			args:   args{ctx: nil, shortCode: ""},
			want:   "test3-full-url",
			want1:  http.StatusFound,
		},
		{
			name:   "expired url",
			fields: fields{urlRepository: urlRepo4},
			args:   args{ctx: nil, shortCode: ""},
			want:   "",
			want1:  http.StatusGone,
		},
		{
			name:   "url not found",
			fields: fields{urlRepository: urlRepo5},
			args:   args{ctx: nil, shortCode: ""},
			want:   "",
			want1:  http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &urlInteractor{
				urlRepository: tt.fields.urlRepository,
				blacklist:     []string{},
			}
			got, got1, err := u.GetRedirectURL(tt.args.ctx, tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRedirectURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRedirectURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetRedirectURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
