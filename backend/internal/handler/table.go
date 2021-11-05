package handler

import (
	"context"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/service"
	"strconv"
	"strings"
)

type TableHandler struct {
	Service TableService
}

type TableService interface {
	TableFilter(ctx context.Context, r service.FilterRequest, page int64) (model.TableData, error)
}

func (h *TableHandler) Method() string {
	return http.MethodGet
}

func (h *TableHandler) Path() string {
	return "/table"
}

func filterRequestFromHttpRequest(r *http.Request) service.FilterRequest {
	rawCountries := r.FormValue("countries")
	countries := []string(nil)
	if rawCountries != "" {
		countries = strings.Split(rawCountries, ",")
	}
	return service.FilterRequest{
		SectorFilter:      r.FormValue("sector"),
		IndustryFilter:    r.FormValue("industry"),
		EmployeesFilter:   r.FormValue("employees"),
		TotalCash:         r.FormValue("total cash"),
		TotalCashPerShare: r.FormValue("total cash per share"),
		Ebitda:            r.FormValue("ebitda"),
		TotalDebt:         r.FormValue("total debt"),
		QuickRatio:        r.FormValue("quick ratio"),
		CurrentRatio:      r.FormValue("current ratio"),
		TotalRevenue:      r.FormValue("total revenue"),
		RevenuePerShare:   r.FormValue("revenue per share"),
		DebtToEquity:      r.FormValue("debt to equity"),
		ReturnOnAssets:    r.FormValue("roa"),
		ReturnOnEquity:    r.FormValue("roe"),
		CountriesFilter:   countries,
	}
}

func (h *TableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page := r.FormValue("page")
	var numberPage int64 = 0

	if page != "" {
		number, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			writeResponse(w, r, badRequest{err.Error()})
			return
		}
		if number < 0 {
			writeResponse(w, r, badRequest{err.Error()})
			return
		}
		numberPage = number
	}

	searchRequest := filterRequestFromHttpRequest(r)

	stocks, err := h.Service.TableFilter(r.Context(), searchRequest, numberPage)

	if err != nil {
		writeResponse(w, r, badRequest{err.Error()})
		return
	}

	writeResponse(w, r, stocks)
}
