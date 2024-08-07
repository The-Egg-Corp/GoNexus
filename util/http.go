package util

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const REQ_TIMEOUT = 15 * time.Second
const DOMAIN = "https://api.nexusmods.com/"

var client = resty.NewWithClient(&http.Client{Timeout: REQ_TIMEOUT})

func get(url string, apiKey string) ([]byte, error) {
	response, err := client.R().
		SetHeader("apiKey", apiKey).
		Get(url)

	if err != nil {
		return nil, err
	}

	// if !response.IsSuccess() {
	// 	fmt.Println(fmt.Sprint(response.StatusCode(), " ", url))
	// }

	return response.Body(), nil
}

func asJSON[T interface{}](res []byte, err error) (T, error) {
	var data T

	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(res), &data)
	return data, nil
}

func JsonGetRequest[T interface{}](endpoint string, apiKey string) (T, error) {
	return asJSON[T](get(DOMAIN+endpoint, apiKey))
}
