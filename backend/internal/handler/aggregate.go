package handler

import (
	"context"
	"net/http"

	"nosql1h21-stock-backend/backend/internal/service"
)

type CountHandler struct {
	Service AggregateService
}

type AggregateService interface {
	Aggregate(ctx context.Context, mode, property, in string, filter service.FilterRequest) ([]service.CountItem, error)
}

func (h *CountHandler) Method() string {
	return http.MethodGet
}

func (h *CountHandler) Path() string {
	return "/aggregate"
}

func (h *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mode := r.FormValue("mode")
	property := r.FormValue("property")
	in := r.FormValue("in")

	filter := filterRequestFromHttpRequest(r)

	countItems, err := h.Service.Aggregate(r.Context(), mode, property, in, filter)
	if err != nil {
		// if _, ok := err.(service.ErrInvalidArgument); ok {
		// 	writeResponse(w, r, badRequest{"Invalid `by`"})
		// 	return
		// }
		writeResponse(w, r, err)
		return
	}
	writeResponse(w, r, countItems)
}
