package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"
)

const (
	Path = "/stock/{ticker}"
)

type Handler struct {
	logger  *zerolog.Logger
	service Service
}

type Service interface {
	GetAllData(ticker string) (model.Stock, error)
}

func New(logger *zerolog.Logger, srv *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: srv,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")

	stock, err := h.service.GetAllData(ticker)
	if err != nil {
		h.logger.Error().Err(err).Msg("Get all data error")
		writeResponse(w, http.StatusInternalServerError, model.Error{Error: "Internal server error"})
		return
	}

	writeResponse(w, http.StatusOK, stock)
}
