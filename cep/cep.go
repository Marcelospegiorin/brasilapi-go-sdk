package cep

import (
	"encoding/json"
	"strconv"

	"github.com/marcelospegiorin/brasilapi-go-sdk/utils/request"
)

type cep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func GetCep(cepReq int) (cep, error) {

	cepBusca := strconv.Itoa(cepReq)

	params := request.Request{
		URL:    "https://brasilapi.com.br/api/cep/v1/" + cepBusca,
		Method: "GET",
	}

	var cepResult cep

	req, err := request.NewRequest(params)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&cepResult)
	if err != nil {
		panic(err)
	}

	return cepResult, nil
}
