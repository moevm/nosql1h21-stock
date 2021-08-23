package handler

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"
	"sync"
)

const (
	SortPath = "/sort"
)

type SortHandler struct {
	logger          *zerolog.Logger
	service         SortService
	validTickersMap *sync.Map
}

type SortService interface {
	SortData(countries []string, industry string, sector string) (*[]model.ValidTicker, error)
}

func NewSortHandler(logger *zerolog.Logger, srv *service.SortService, validTickersMap *sync.Map) *SortHandler {
	return &SortHandler{
		logger:          logger,
		service:         srv,
		validTickersMap: validTickersMap,
	}
}

func (s *SortHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	sortRequest := model.SortRequest{}
	err := json.NewDecoder(r.Body).Decode(&sortRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	validTickers, err := s.service.SortData(sortRequest.Countries, sortRequest.Sector, sortRequest.Industry)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	writeResponse(w, http.StatusOK, validTickers)
}
