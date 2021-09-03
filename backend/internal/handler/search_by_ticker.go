package handler

import (
	"context"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"

	"github.com/go-chi/chi/v5"
)

type SearchByTickerHandler struct {
	Service SearchByTickerService
}

type SearchByTickerService interface {
	SearchByTicker(ctx context.Context, tickerFragment string) (stocks []model.StockOverview, _ error)
}

func (h *SearchByTickerHandler) Method() string {
	return http.MethodGet
}

func (h *SearchByTickerHandler) Path() string {
	return "/search-by-ticker/{tickerFragment}"
}

func (h *SearchByTickerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tickerFragment := chi.URLParam(r, "tickerFragment")

	stocks, err := h.Service.SearchByTicker(r.Context(), tickerFragment)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, stocks)
}
