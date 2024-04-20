package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	URL     string
	Method  string
	Header  string
	Client  *http.Client
	Payload interface{}
}

func NewRequest(params Request) (*http.Response, error) {
	payloadData, err := json.Marshal(params.Payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(params.Method, params.URL, bytes.NewBuffer(payloadData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
