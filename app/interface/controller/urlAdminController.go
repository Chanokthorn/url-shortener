package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortener/app/usecase/interactor"
)

type URLAdminController interface {
	ListURL(e echo.Context) error
	DeleteURL(e echo.Context) error
}

type urlAdminController struct {
	urlInteractor interactor.URLInteractor
}

func NewURLAdminController(urlInteractor interactor.URLInteractor) URLAdminController {
	return &urlAdminController{urlInteractor: urlInteractor}
}

func (u *urlAdminController) ListURL(e echo.Context) error {
	shortCodeFilter := e.QueryParam("short-code-filter")
	fullURLKeywordFilter := e.QueryParam("full-url-keyword-filter")
	urlList, err := u.urlInteractor.ListURL(e.Request().Context(), shortCodeFilter, fullURLKeywordFilter)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to list url: %v", err))
	}
	err = e.JSON(http.StatusOK, urlList)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse change to JSON")
	}
	return nil
}

func (u *urlAdminController) DeleteURL(e echo.Context) error {
	shortCode := e.Param("short-code")
	if shortCode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty short code")
	}
	err := u.urlInteractor.DeleteURL(e.Request().Context(), shortCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to delete url: %v", err))
	}
	err = e.String(http.StatusOK, "deleted")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable send response: %v", err))
	}
	return nil
}
