package handler

import (
	"context"
	"net/http"
	"strings"
)

type FilterHandler struct {
	Service FilterService
}

type FilterService interface {
	Filter(ctx context.Context, countries []string, sector, industry string) (tickers []string, _ error)
}

func (h *FilterHandler) Method() string {
	return http.MethodGet
}

func (h *FilterHandler) Path() string {
	return "/filter"
}

func (h *FilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rawCountries := r.FormValue("countries")
	countries := []string(nil)
	if rawCountries != "" {
		countries = strings.Split(rawCountries, ",")
	}
	sector := r.FormValue("sector")
	industry := r.FormValue("industry")

	tickers, err := h.Service.Filter(r.Context(), countries, sector, industry)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, tickers)
}
