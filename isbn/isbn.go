package isbn

import (
	"encoding/json"
	"strconv"

	"github.com/Marcelospegiorin/brasilapi-go-sdk/utils/request"
)

type isbn struct {
	Isbn       string   `json:"isbn"`
	Title      string   `json:"title"`
	Subtitle   any      `json:"subtitle"`
	Authors    []string `json:"authors"`
	Publisher  string   `json:"publisher"`
	Synopsis   string   `json:"synopsis"`
	Dimensions struct {
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
		Unit   string  `json:"unit"`
	} `json:"dimensions"`
	Year        int      `json:"year"`
	Format      string   `json:"format"`
	PageCount   int      `json:"page_count"`
	Subjects    []string `json:"subjects"`
	Location    string   `json:"location"`
	RetailPrice any      `json:"retail_price"`
	CoverURL    any      `json:"cover_url"`
	Provider    string   `json:"provider"`
}

func Get(isbnCode int) (isbn, error) {
	id := strconv.Itoa(isbnCode)

	params := request.Request{
		URL:    "https://brasilapi.com.br/api/isbn/v1/" + id,
		Method: "GET",
	}

	req, err := request.NewRequest(params)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	var isbnResult isbn

	err = json.NewDecoder(req.Body).Decode(&isbnResult)
	if err != nil {
		panic(err)
	}

	return isbnResult, nil
}
