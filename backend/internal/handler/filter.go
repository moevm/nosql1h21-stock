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
	SortPath = "/filter"
)

type FilterHandler struct {
	logger          *zerolog.Logger
	service         FilterService
	validTickersMap *sync.Map
}

type FilterService interface {
	FilterData(countries []string, industry string, sector string) (*[]model.ValidTicker, error)
}

func NewFilterHandler(logger *zerolog.Logger, srv *service.FilterService, validTickersMap *sync.Map) *FilterHandler {
	return &FilterHandler{
		logger:          logger,
		service:         srv,
		validTickersMap: validTickersMap,
	}
}

func (s *FilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	sortRequest := model.FilterRequest{}
	err := json.NewDecoder(r.Body).Decode(&sortRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	validTickers, err := s.service.FilterData(sortRequest.Countries, sortRequest.Sector, sortRequest.Industry)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	writeResponse(w, http.StatusOK, validTickers)
}
