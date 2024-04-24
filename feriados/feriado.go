package feriados

import (
	"encoding/json"
	"strconv"

	"github.com/marcelospegiorin/brasilapi-go-sdk/utils/request"
)

type feriado struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func Get(ano int) ([]feriado, error) {
	anoReq := strconv.Itoa(ano)

	params := request.Request{
		URL:    "https://brasilapi.com.br/api/feriados/v1/" + anoReq,
		Method: "GET",
	}

	var feriadosResult []feriado

	req, err := request.NewRequest(params)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&feriadosResult)
	if err != nil {
		return nil, err
	}

	return feriadosResult, nil
}
