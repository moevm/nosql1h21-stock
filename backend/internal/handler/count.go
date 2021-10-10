package handler

import (
	"context"
	"net/http"

	"nosql1h21-stock-backend/backend/internal/service"
)

type CountHandler struct {
	Service CountService
}

type CountService interface {
	Count(ctx context.Context, by string) ([]service.CountItem, error)
}

func (h *CountHandler) Method() string {
	return http.MethodGet
}

func (h *CountHandler) Path() string {
	return "/count"
}

func (h *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// by := []string(nil)
	// if rawBy := r.FormValue("by"); rawBy != "" {
	// 	by = strings.Split(rawBy, ",")
	// }
	by := r.FormValue("by")

	countItems, err := h.Service.Count(r.Context(), by)
	if err != nil {
		if _, ok := err.(service.ErrInvalidArgument); ok {
			writeResponse(w, r, badRequest{"Invalid `by`"})
			return
		}
		writeResponse(w, r, err)
		return
	}
	writeResponse(w, r, countItems)
}
