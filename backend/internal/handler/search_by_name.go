package handler

import (
	"context"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"

	"github.com/go-chi/chi/v5"
)

type SearchByNameHandler struct {
	Service SearchByNameService
}

type SearchByNameService interface {
	SearchByName(ctx context.Context, nameFragment string) (stocks []model.StockOverview, _ error)
}

func (h *SearchByNameHandler) Method() string {
	return http.MethodGet
}

func (h *SearchByNameHandler) Path() string {
	return "/search-by-name/{nameFragment}"
}

func (h *SearchByNameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nameFragment := chi.URLParam(r, "nameFragment")

	stocks, err := h.Service.SearchByName(r.Context(), nameFragment)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, stocks)
}
