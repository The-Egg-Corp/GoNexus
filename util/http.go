package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type ReqMethod string

const (
	POST ReqMethod = "POST"
	GET  ReqMethod = "GET"
)

const API_BASE_URL = "https://api.nexusmods.com/"
const REQ_TIMEOUT = 15 * time.Second

var client = resty.NewWithClient(&http.Client{Timeout: REQ_TIMEOUT})

func fetchWithKey(url string, method ReqMethod, apiKey string) (res *resty.Response, err error) {
	req := client.R()
	req.SetHeader("apiKey", apiKey)

	if method == "POST" {
		res, err = req.Post(url)
	} else {
		res, err = req.Get(url)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

// After completing an HTTP request, this function attempts to unmarshal the body (assumed to be JSON) of the response into the given interface.
// In the case that the body is empty or there was an error unmarshalling, the result will be nil with the original error provided.
func ParseJsonBody[T interface{}](res resty.Response) (result *T, err error) {
	body := res.Body()
	if len(body) < 1 {
		return nil, fmt.Errorf("\nno body in response:\n\n%v", res)
	}

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, fmt.Errorf("\nerror unmarshalling json:\n\n%v", err)
	}

	return result, err
}

func GetRequest(endpoint string, apiKey string) (*resty.Response, error) {
	res, err := fetchWithKey(API_BASE_URL+endpoint, GET, apiKey)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func PostRequest(endpoint string, apiKey string) (*resty.Response, error) {
	res, err := fetchWithKey(API_BASE_URL+endpoint, POST, apiKey)
	if err != nil {
		return nil, err
	}

	return res, nil
}
