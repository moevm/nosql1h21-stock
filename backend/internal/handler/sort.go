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

func (h *SortHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/*query := r.URL.Query()
	countryFilters, present := query["country"]

	industryFilters, present := query["industry"]
	if length:= len(industryFilters); !present || length == 0 || length > 1{
		h.logger.Info().Msg("Invalid argument for request")
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	sectorFilters, present := query["sector"]
	if length:= len(sectorFilters); !present || length == 0 || length > 1 {
		h.logger.Info().Msg("Invalid argument for request")
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	validTickers ,err := h.service.SortData(countryFilters, industryFilters[0], sectorFilters[0])

	if err !=nil {
		_ = validTickers
	}

	writeResponse(w, http.StatusOK, validTickers)*/

	u := model.SortRequest{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	validTickers, err := h.service.SortData(u.Countries, u.Industry, u.Sector)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Invalid argument for request"})
		return
	}

	writeResponse(w, http.StatusOK, validTickers)
}
