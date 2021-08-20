package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"
	"sync"
)

const (
	StockPath = "/stock/{ticker}"
)

type StockHandler struct {
	logger          *zerolog.Logger
	service         StocksService
	validTickersMap *sync.Map
}

type StocksService interface {
	GetAllData(ticker string) (model.Stock, error)
}

func NewStockHandler(logger *zerolog.Logger, srv *service.StockService) *StockHandler {
	return &StockHandler{
		logger:  logger,
		service: srv,
	}
}

func (h *StockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")

	if _, ok := h.validTickersMap.Load(ticker); !ok {
		h.logger.Info().Msg("Unknown ticker " + ticker + " in request")
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Unknown ticker"})
		return
	}

	stock, err := h.service.GetAllData(ticker)
	if err != nil {
		h.logger.Error().Err(err).Msg("GetAllData error")
		writeResponse(w, http.StatusInternalServerError, model.Error{Error: "Internal server error"})
		return
	}

	writeResponse(w, http.StatusOK, stock)
}
