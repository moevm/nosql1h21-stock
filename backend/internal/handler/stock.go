package handler

import (
	"context"
	"errors"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type StockHandler struct {
	Service StockService
}

type StockService interface {
	GetStockInfo(ctx context.Context, ticker string) (*model.Stock, error)
}

func (h *StockHandler) Method() string {
	return http.MethodGet
}

func (h *StockHandler) Path() string {
	return "/stock/{ticker}"
}

func (h *StockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")

	stock, err := h.Service.GetStockInfo(r.Context(), ticker)

	if err != nil {
		if nsi := (service.NoStockInfo{}); errors.As(err, &nsi) {
			writeResponse(w, r, badRequest{nsi.Error()})
			return
		}
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, stock)
}
