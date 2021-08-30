package handler

import (
	"context"
	"net/http"
)

type CountriesHandler struct {
	Service CountriesService
}

type CountriesService interface {
	GetCountries(ctx context.Context) (countries []string, _ error)
}

func (h *CountriesHandler) Method() string {
	return http.MethodGet
}

func (h *CountriesHandler) Path() string {
	return "/countries"
}

func (h *CountriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tickers, err := h.Service.GetCountries(r.Context())

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, tickers)
}
