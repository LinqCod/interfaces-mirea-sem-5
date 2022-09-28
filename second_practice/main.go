package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ApiKey = "JcL71lRK2latTwVVNxNaeBV0J7yGD6cl"
	Base   = "https://api.apilayer.com/currency_data"
)

type requester struct {
	client *http.Client
}

func newRequester() *requester {
	return &requester{
		client: &http.Client{},
	}
}

func (r requester) doRequest(req *http.Request) {
	res, err := r.client.Do(req)
	if err != nil {
		log.Fatalf("error while sending request. %v", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func main() {
	r := newRequester()

	//TEST - 1
	fmt.Println("---список поддерживаемых API Currency Layer валют---")
	reqSupportedCurrencies, err := http.NewRequest("GET", Base+"/list", nil)
	if err != nil {
		log.Fatalf("error while creating supported currencies request. %v", err)
	}
	reqSupportedCurrencies.Header.Set("apikey", ApiKey)
	r.doRequest(reqSupportedCurrencies)

	//TEST - 2
	fmt.Println("---исторические курсы валют от 22 февраля 2018 года, для евро, фунтов стерлингов и иен к доллару США---")
	reqHistoricalCurrencies, err := http.NewRequest("GET", Base+"/historical?date=2018-02-22&currencies=EUR,GBP,JPY&source=USD", nil)
	if err != nil {
		log.Fatalf("error while creating historical currencies request. %v", err)
	}
	reqHistoricalCurrencies.Header.Set("apikey", ApiKey)
	r.doRequest(reqHistoricalCurrencies)

	//TEST - 3
	fmt.Println("---исторические данные о курсе евро к доллару США (2016-02-25:2017-02-21)---")
	reqHistoricalPeriodCurrencies, err := http.NewRequest("GET", Base+"/timeframe?start_date=2016-02-25&end_date=2017-02-21&currencies=EUR&source=USD", nil)
	if err != nil {
		log.Fatalf("error while creating historical period currencies request. %v", err)
	}
	reqHistoricalPeriodCurrencies.Header.Set("apikey", ApiKey)
	r.doRequest(reqHistoricalPeriodCurrencies)
}
