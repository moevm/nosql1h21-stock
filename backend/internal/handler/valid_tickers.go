package handler

import (
	"github.com/rs/zerolog"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"
)

const (
	ValidTickersPath = "/validData"
)

type ValidTickersHandler struct {
	logger  *zerolog.Logger
	service ValidTickersService
}

type ValidTickersService interface {
	GetValidData() (*model.ValidData, error)
}

func NewValidTickersHandler(logger *zerolog.Logger, srv *service.ValidTickersService) *ValidTickersHandler {
	return &ValidTickersHandler{
		logger:  logger,
		service: srv,
	}
}

func (h *ValidTickersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	validData, err := h.service.GetValidData()
	if err != nil {
		h.logger.Error().Err(err).Msg("GetValidData error")
		writeResponse(w, http.StatusInternalServerError, model.Error{Error: "Internal server error"})
		return
	}

	if len(validData.Tickers) == 0 {
		h.logger.Error().Err(err).Msg("GetValidData error")
		writeResponse(w, http.StatusInternalServerError, model.Error{Error: "Valid tickers list empty"})
		return
	}

	if len(validData.Sectors) == 0 {
		h.logger.Error().Err(err).Msg("GetValidData error")
		writeResponse(w, http.StatusInternalServerError, model.Error{Error: "Valid tickers list empty"})
		return
	}

	writeResponse(w, http.StatusOK, validData)
}
