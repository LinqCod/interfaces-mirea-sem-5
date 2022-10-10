package test

import (
	"encoding/json"
	requesterConfig "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/config"
	requester "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/service"
	"testing"
)

type currencySuccessChecker struct {
	Success bool `json:"success"`
}

type historicalCurrenciesByPeriodTest struct {
	startDate  string
	endDate    string
	currencies string
	source     string
	success    bool
}

var historicalCurrenciesByPeriodTests = []historicalCurrenciesByPeriodTest{
	{
		startDate:  "2016-02-25",
		endDate:    "2016-02-26",
		currencies: "EUR,RUB",
		source:     "USD",
		success:    true,
	},
	{
		startDate:  "2013-03-17",
		endDate:    "2013-03-19",
		currencies: "RUB",
		source:     "EUR",
		success:    true,
	},
}

func TestGetHistoricalCurrenciesByPeriod(t *testing.T) {
	reqConfig := requesterConfig.Config{
		ApiKey: "JcL71lRK2latTwVVNxNaeBV0J7yGD6cl",
		Base:   "https://api.apilayer.com/currency_data",
	}
	rService := requester.NewService(reqConfig)

	var successChecker currencySuccessChecker
	for _, test := range historicalCurrenciesByPeriodTests {
		output, err := rService.GetHistoricalCurrenciesByPeriod(test.startDate, test.endDate, test.currencies, test.source)
		if err != nil {
			t.Errorf("cannot get currencies history: %v", err)
		}

		_ = json.Unmarshal(output, &successChecker)
		if successChecker.Success != test.success {
			t.Errorf("Output success %v not equal to expected success status %v", successChecker.Success, test.success)
		}
	}
}

func TestGetSupportedCurrencies(t *testing.T) {
	reqConfig := requesterConfig.Config{
		ApiKey: "JcL71lRK2latTwVVNxNaeBV0J7yGD6cl",
		Base:   "https://api.apilayer.com/currency_data",
	}
	rService := requester.NewService(reqConfig)

	output, err := rService.GetSupportedCurrencies()
	if err != nil {
		t.Errorf("cannot get currencies history: %v", err)
	}

	var successChecker currencySuccessChecker
	_ = json.Unmarshal(output, &successChecker)
	if successChecker.Success != true {
		t.Errorf("Output success should be true, but got: %v", successChecker.Success)
	}
}
