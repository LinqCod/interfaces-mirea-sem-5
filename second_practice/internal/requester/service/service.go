package requester

import (
	"errors"
	"fmt"
	requester "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/config"
	"io/ioutil"
	"log"
	"net/http"
)

type Service struct {
	config requester.Config
	client *http.Client
}

func NewService(config requester.Config) *Service {
	return &Service{
		config: config,
		client: &http.Client{},
	}
}

func (r Service) doRequest(req *http.Request) ([]byte, error) {
	res, err := r.client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while sending request. %v", err))
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	result, _ := ioutil.ReadAll(res.Body)

	return result, nil
}

// GetHistoricalCurrenciesByPeriod
// Метод получения курса валют %currencies к %source за определенный период
// временной промежуток задается переменными %startDate и %endDate формата "YYYY-MM-DD"
func (r Service) GetHistoricalCurrenciesByPeriod(startDate, endDate, currencies, source string) ([]byte, error) {
	log.Printf("\n---исторические данные о курсах %s к %s (%s:%s)---", currencies, source, startDate, endDate)

	reqHistoricalPeriodCurrencies, err := http.NewRequest("GET",
		fmt.Sprintf("%s/timeframe?start_date=%s&end_date=%s&currencies=%s&source=%s",
			r.config.Base, startDate, endDate, currencies, source), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while creating historical period currencies request. %v", err))
	}
	reqHistoricalPeriodCurrencies.Header.Set("apikey", r.config.ApiKey)

	result, err := r.doRequest(reqHistoricalPeriodCurrencies)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetHistoricalCurrencies
// Метод получения курса валют %currencies к %source за конкретную дату
// Дата задается переменной %date формата "YYYY-MM-DD"
func (r Service) GetHistoricalCurrencies(date, currencies, source string) ([]byte, error) {
	log.Printf("\n---исторические курсы валют от %s, для %s к %s---", date, currencies, source)

	reqHistoricalCurrencies, err := http.NewRequest("GET",
		fmt.Sprintf("%s/historical?date=%s&currencies=%s&source=%s",
			r.config.Base, date, currencies, source), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while creating historical currencies request. %v", err))
	}
	reqHistoricalCurrencies.Header.Set("apikey", r.config.ApiKey)

	result, err := r.doRequest(reqHistoricalCurrencies)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSupportedCurrencies
// Метод получения списка поддерживаемых API Currency Layer валют
func (r Service) GetSupportedCurrencies() ([]byte, error) {
	log.Println("\n---список поддерживаемых API Currency Layer валют---")

	reqSupportedCurrencies, err := http.NewRequest("GET",
		fmt.Sprintf("%s/list", r.config.Base), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while creating supported currencies request. %v", err))
	}
	reqSupportedCurrencies.Header.Set("apikey", r.config.ApiKey)

	result, err := r.doRequest(reqSupportedCurrencies)
	if err != nil {
		return nil, err
	}

	return result, nil
}
