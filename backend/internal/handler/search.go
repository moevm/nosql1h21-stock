package handler

import (
	"context"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"
	"strings"
)

type SearchHandler struct {
	Service SearchService
}

type SearchService interface {
	Search(ctx context.Context, r service.SearchRequest) (stocks []model.StockOverview, _ error)
}

func (h *SearchHandler) Method() string {
	return http.MethodGet
}

func (h *SearchHandler) Path() string {
	return "/search"
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rawCountries := r.FormValue("countries")
	countries := []string(nil)
	if rawCountries != "" {
		countries = strings.Split(rawCountries, ",")
	}

	searchRequest := service.SearchRequest{
		Fragment:  r.FormValue("fragment"),
		Sector:    r.FormValue("sector"),
		Industry:  r.FormValue("industry"),
		Countries: countries,
	}
	stocks, err := h.Service.Search(r.Context(), searchRequest)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, stocks)
}
