package handler

import (
	"context"
	"net/http"
)

type ExportHandler struct {
	Service ExportService
}

type ExportService interface {
	Export(ctx context.Context) (jsonEncoded []byte, _ error)
}

func (h *ExportHandler) Method() string {
	return http.MethodGet
}

func (h *ExportHandler) Path() string {
	return "/export"
}

func (h *ExportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	export, err := h.Service.Export(r.Context())
	if err != nil {
		writeResponse(w, r, err)
		return
	}
	w.Header().Set("Content-Disposition", `attachment; filename="stocks_data.json"`)
	w.Write(export)
}
