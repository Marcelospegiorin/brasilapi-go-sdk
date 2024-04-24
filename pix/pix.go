package pix

import (
	"encoding/json"
	"time"

	"github.com/marcelospegiorin/brasilapi-go-sdk/utils/request"
)

type pix struct {
	Ispb                   string    `json:"ispb"`
	Nome                   string    `json:"nome"`
	NomeReduzido           string    `json:"nome_reduzido"`
	ModalidadeParticipacao string    `json:"modalidade_participacao"`
	TipoParticipacao       string    `json:"tipo_participacao"`
	InicioOperacao         time.Time `json:"inicio_operacao"`
}

func GetAll() ([]pix, error) {
	params := request.Request{
		URL:    "https://brasilapi.com.br/api/pix/v1/participants",
		Method: "GET",
	}

	req, err := request.NewRequest(params)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	var pixResult []pix

	err = json.NewDecoder(req.Body).Decode(&pixResult)
	if err != nil {
		return nil, err
	}

	return pixResult, nil
}
