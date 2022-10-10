package requester

import (
	"errors"
	"fmt"
	requester "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/config"
	"io/ioutil"
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
