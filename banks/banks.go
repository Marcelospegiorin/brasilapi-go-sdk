package banks

import (
	"encoding/json"
	"strconv"

	"github.com/Marcelospegiorin/brasilapi-go-sdk/utils/request"
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

func GetOne(idReq int) (bank, error) {

	id := strconv.Itoa(idReq)

	params := request.Request{
		URL:    "https://brasilapi.com.br/api/banks/v1/" + id,
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
