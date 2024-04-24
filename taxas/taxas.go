package taxas

import (
	"encoding/json"

	"github.com/Marcelospegiorin/brasilapi-go-sdk/utils/request"
)

type taxa struct {
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}

func GetAll() ([]taxa, error) {
	params := request.Request{
		URL:    "https://brasilapi.com.br/api/taxas/v1",
		Method: "GET",
	}

	req, err := request.NewRequest(params)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	var taxasResult []taxa

	err = json.NewDecoder(req.Body).Decode(&taxasResult)
	if err != nil {
		return nil, err
	}

	return taxasResult, nil
}

func GetOne(sigla string) (taxa, error) {
	params := request.Request{
		URL:    "https://brasilapi.com.br/api/taxas/v1/" + sigla,
		Method: "GET",
	}

	req, err := request.NewRequest(params)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	var taxasResult taxa

	err = json.NewDecoder(req.Body).Decode(&taxasResult)
	if err != nil {
		panic(err)
	}

	return taxasResult, nil
}
