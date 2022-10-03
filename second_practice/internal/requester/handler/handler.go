package Handler

import (
	"fmt"
	requester "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/service"
	"net/http"
)

type Handler struct {
	service *requester.Service
}

func NewHandler(service *requester.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (r Handler) GetHistoricalCurrenciesByPeriod(rw http.ResponseWriter, req *http.Request) {
	startDate := req.URL.Query().Get("startDate")
	endDate := req.URL.Query().Get("endDate")
	currencies := req.URL.Query().Get("currencies")
	source := req.URL.Query().Get("source")

	result, err := r.service.GetHistoricalCurrenciesByPeriod(startDate, endDate, currencies, source)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%v", err)
	}

	fmt.Fprintf(rw, "%s", result)
}

func (r Handler) GetHistoricalCurrencies(rw http.ResponseWriter, req *http.Request) {
	date := req.URL.Query().Get("date")
	currencies := req.URL.Query().Get("currencies")
	source := req.URL.Query().Get("source")

	result, err := r.service.GetHistoricalCurrencies(date, currencies, source)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%v", err)
	}

	fmt.Fprintf(rw, "%s", result)
}

func (r Handler) GetSupportedCurrencies(rw http.ResponseWriter, req *http.Request) {
	result, err := r.service.GetSupportedCurrencies()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%v", err)
	}

	fmt.Fprintf(rw, "%s", result)
}
