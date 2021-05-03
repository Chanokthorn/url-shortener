package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"url-shortener/app/interface/internal"
	"url-shortener/app/usecase/interactor"
)

type URLClientController interface {
	CreateURL(e echo.Context) error
	GetShortCodeFromFullURL(e echo.Context) error
}

type urlClientController struct {
	urlInteractor interactor.URLInteractor
}

func NewURLClientController(URLInteractor interactor.URLInteractor) URLClientController {
	return &urlClientController{urlInteractor: URLInteractor}
}

func (u *urlClientController) CreateURL(e echo.Context) error {
	var (
		expireDate   time.Time
		hasExireDate bool
	)
	request := new(internal.CreateURLRequest)
	err := e.Bind(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if request.ExpireDate != nil {
		hasExireDate = true
		expireDate, err = time.Parse(time.RFC3339, *request.ExpireDate)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid date format")
	}
	shortCode, err := u.urlInteractor.CreateURL(e.Request().Context(), request.FullURL, hasExireDate, expireDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to create url: %v", err))
	}
	err = e.String(http.StatusOK, shortCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse change to JSON")
	}
	return nil
}

func (u *urlClientController) GetShortCodeFromFullURL(e echo.Context) error {
	fullURL := e.QueryParam("full-url")
	if fullURL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty full-url query param")
	}
	shortCode, err := u.urlInteractor.GetShortCode(e.Request().Context(), fullURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to get short code: %v", err))
	}
	err = e.String(http.StatusOK, shortCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse change to JSON")
	}
	return nil
}
