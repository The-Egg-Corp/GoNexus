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

type ReqMethod string

const (
	Post ReqMethod = "POST"
	Get  ReqMethod = "GET"
)

func fetch(url string, method ReqMethod, apiKey string) ([]byte, error) {
	req := client.R().SetHeader("apiKey", apiKey)

	var res *resty.Response
	var err error

	if method == "POST" {
		res, err = req.Post(url)
	} else {
		res, err = req.Get(url)
	}

	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}

func asJSON[T interface{}](res []byte, err error) (T, error) {
	var data T
	if err != nil {
		return data, err
	}

	e := json.Unmarshal([]byte(res), &data)
	return data, e
}

func JsonGetRequest[T interface{}](endpoint string, apiKey string) (T, error) {
	return asJSON[T](fetch(DOMAIN+endpoint, Get, apiKey))
}

func JsonPostRequest[T interface{}](endpoint string, apiKey string) (T, error) {
	return asJSON[T](fetch(DOMAIN+endpoint, Post, apiKey))
}
