package controller

import "url-shortener/app/usecase/interactor"

type URLRedirectController interface {
}

type urlRedirectController struct {
	urlInteractor interactor.URLInteractor
}

func NewURLRedirectController(urlInteractor interactor.URLInteractor) URLRedirectController {
	return &urlRedirectController{urlInteractor: urlInteractor}
}
