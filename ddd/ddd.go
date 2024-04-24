package ddd

import (
	"encoding/json"
	"strconv"

	"github.com/marcelospegiorin/brasilapi-go-sdk/utils/request"
)

type ddd struct {
	State  string   `json:"state"`
	Cities []string `json:"cities"`
}

func Get(id int) (ddd, error) {
	dddReq := strconv.Itoa(id)

	params := request.Request{
		URL:    "https://brasilapi.com.br/api/ddd/v1/" + dddReq,
		Method: "GET",
	}

	var dddResult ddd

	req, err := request.NewRequest(params)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&dddResult)
	if err != nil {
		panic(err)
	}

	return dddResult, nil
}
