package handler

import (
	"context"
	"io"
	"net/http"
)

type ImportHandler struct {
	Service ImportService
}

type ImportService interface {
	Import(ctx context.Context, jsonEncoded io.Reader) error
}

func (h *ImportHandler) Method() string {
	return http.MethodPost
}

func (h *ImportHandler) Path() string {
	return "/import"
}

func (h *ImportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Import(r.Context(), r.Body)
	if err != nil {
		writeResponse(w, r, err)
		return
	}
}
