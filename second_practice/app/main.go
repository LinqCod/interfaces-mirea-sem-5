package main

import (
	"fmt"
	requesterConfig "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/config"
	requesterHandler "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/handler"
	requesterService "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/service"
	serverConfig "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/server/config"
	"log"
	"net/http"
)

func main() {
	reqConfig, err := requesterConfig.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	rService := requesterService.NewService(reqConfig)
	rHandler := requesterHandler.NewHandler(rService)

	http.HandleFunc("/api/historicalCurrenciesByPeriod", rHandler.GetHistoricalCurrenciesByPeriod)
	http.HandleFunc("/api/historicalCurrencies", rHandler.GetHistoricalCurrencies)
	http.HandleFunc("/api/supportedCurrencies", rHandler.GetSupportedCurrencies)

	sConfig, err := serverConfig.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	fmt.Printf("Starting server on port: %v", sConfig.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", sConfig.Port), nil))
}
