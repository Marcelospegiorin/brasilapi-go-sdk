package banks

import (
	"brasilapi/utils/request"
	"encoding/json"
)

type bank struct {
	Isbp     string `json:"isbp"`
	Name     string `json:"name"`
	Code     int64  `json:"code"`
	FullName string `json:"fullname"`
}

func GetAll() ([]bank, error) {
	params := request.Request{
		URL:    "https://brasilapi.com.br/api/banks/v1",
		Method: "GET",
	}

	var banksResult []bank

	req, err := request.NewRequest(params)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&banksResult)
	if err != nil {
		return nil, err
	}

	return banksResult, nil
}

func GetOne(id int) (bank, error) {

	params := request.Request{
		URL:    "https://brasilapi.com.br/api/banks/v1/260",
		Method: "GET",
	}

	var bankResult bank

	req, err := request.NewRequest(params)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&bankResult)
	if err != nil {
		panic(err)
	}

	return bankResult, nil
}
