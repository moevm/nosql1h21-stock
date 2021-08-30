package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SearchByNameHandler struct {
	Service SearchByNameService
}

type SearchByNameService interface {
	SearchByName(ctx context.Context, nameFragment string) (tickers []string, _ error)
}

func (h *SearchByNameHandler) Method() string {
	return http.MethodGet
}

func (h *SearchByNameHandler) Path() string {
	return "/search-by-name/{nameFragment}"
}

func (h *SearchByNameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nameFragment := chi.URLParam(r, "nameFragment")

	tickers, err := h.Service.SearchByName(r.Context(), nameFragment)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, tickers)
}
