package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SearchByTickerHandler struct {
	Service SearchByTickerService
}

type SearchByTickerService interface {
	SearchByTicker(ctx context.Context, tickerFragment string) (tickers []string, _ error)
}

func (h *SearchByTickerHandler) Method() string {
	return http.MethodGet
}

func (h *SearchByTickerHandler) Path() string {
	return "/search-by-ticker/{tickerFragment}"
}

func (h *SearchByTickerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tickerFragment := chi.URLParam(r, "tickerFragment")

	tickers, err := h.Service.SearchByTicker(r.Context(), tickerFragment)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, tickers)
}
