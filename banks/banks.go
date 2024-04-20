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

	resp, err := request.NewRequest(params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&banksResult)
	if err != nil {
		return nil, err
	}

	return banksResult, nil
}

func GetOne(id string) ([]bank, error) {
	params := request.Request{
		URL:    "https://brasilapi.com.br/api/banks/v1" + id,
		Method: "GET",
	}

	var banksResult []bank

	resp, err := request.NewRequest(params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&banksResult)
	if err != nil {
		return nil, err
	}

	return banksResult, nil
}
