package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortener/app/usecase/interactor"
)

type URLRedirectController interface {
	Redirect(e echo.Context) error
}

type urlRedirectController struct {
	urlInteractor interactor.URLInteractor
}

func NewURLRedirectController(urlInteractor interactor.URLInteractor) URLRedirectController {
	return &urlRedirectController{urlInteractor: urlInteractor}
}

func (u *urlRedirectController) Redirect(e echo.Context) error {
	shortCode := e.Param("short-code")
	if shortCode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty short code")
	}
	url, statusCode, err := u.urlInteractor.GetRedirectURL(e.Request().Context(), shortCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to get redirect url: %v", err))
	}
	if statusCode == http.StatusFound {
		err = e.Redirect(statusCode, url)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to redirect: %v", err))
		}
		return nil
	}
	if statusCode == http.StatusGone || statusCode == http.StatusNotFound {
		err = e.String(statusCode, "")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to send response: %v", err))
		}
	}
	if statusCode == http.StatusInternalServerError {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("to get redirect url: %v", err))
	}
	return nil
}
