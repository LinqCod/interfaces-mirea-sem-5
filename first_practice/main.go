package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ApiKey = "JcL71lRK2latTwVVNxNaeBV0J7yGD6cl"

func main() {
	url := "https://api.apilayer.com/currency_data/convert?to=USD&from=RUB&amount=120"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("error while creating request. %v", err)
	}
	req.Header.Set("apikey", ApiKey)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("error while sending request. %v", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
