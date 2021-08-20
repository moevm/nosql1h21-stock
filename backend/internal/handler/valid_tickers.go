package handler

import (
	"github.com/rs/zerolog"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"
)

const (
	ValidTickersPath = "/validTickers"
)

type ValidTickersHandler struct {
	logger  *zerolog.Logger
	service ValidTickersService
}

type ValidTickersService interface {
	GetValidTickers() (*[]model.ValidTicker, error)
}

func NewValidTickersHandler(logger *zerolog.Logger, srv *service.ValidTickersService) *ValidTickersHandler {
	return &ValidTickersHandler{
		logger:  logger,
		service: srv,
	}
}

func (h *ValidTickersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	validTickers, err := h.service.GetValidTickers()
	if err != nil {
		h.logger.Error().Err(err).Msg("GetAllData error")
		writeResponse(w, http.StatusInternalServerError, model.Error{Error: "Internal server error"})
		return
	}

	if len(*validTickers) == 0 {
		h.logger.Error().Err(err).Msg("GetValidTickers error")
		writeResponse(w, http.StatusInternalServerError, model.Error{Error: "Valid tickers list empty"})
		return
	}

	writeResponse(w, http.StatusOK, validTickers)
}
