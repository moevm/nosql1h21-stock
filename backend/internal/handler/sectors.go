package handler

import (
	"context"
	"net/http"
)

type SectorsHandler struct {
	Service SectorsService
}

type SectorsService interface {
	GetSectors(ctx context.Context) (sectors []string, _ error)
}

func (h *SectorsHandler) Method() string {
	return http.MethodGet
}

func (h *SectorsHandler) Path() string {
	return "/sectors"
}

func (h *SectorsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tickers, err := h.Service.GetSectors(r.Context())

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, tickers)
}
