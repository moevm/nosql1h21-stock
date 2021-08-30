package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type IndustriesHandler struct {
	Service IndustriesService
}

type IndustriesService interface {
	GetIndustries(ctx context.Context, sector string) (industries []string, _ error)
}

func (h *IndustriesHandler) Method() string {
	return http.MethodGet
}

func (h *IndustriesHandler) Path() string {
	return "/industries-in-sector/{sector}"
}

func (h *IndustriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sector := chi.URLParam(r, "sector")

	tickers, err := h.Service.GetIndustries(r.Context(), sector)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, tickers)
}
