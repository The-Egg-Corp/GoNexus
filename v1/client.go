package v1

import (
	"errors"

	"github.com/the-egg-corp/gonexus/util"
)

type Client struct {
	apiKey string
}

func NewNexusClient(key string) (*Client, error) {
	if key == "" {
		return nil, errors.New("error creating client: cannot use empty string as api key")
	}

	user, err := ValidateUser(key)
	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, errors.New("error creating client: invalid api key provided")
	}

	return &Client{apiKey: key}, nil
}

func jsonGetRequest[T interface{}](endpoint string, client Client) (*T, error) {
	res, err := util.GetRequest(endpoint, client.apiKey)
	if err != nil {
		return nil, err
	}

	return util.ParseJsonBody[T](*res)
}

func jsonPostRequest[T interface{}](endpoint string, client Client) (*T, error) {
	res, err := util.PostRequest(endpoint, client.apiKey)
	if err != nil {
		return nil, err
	}

	return util.ParseJsonBody[T](*res)
}

// func ValidateKey(key *string) error {
// 	if key == nil {
// 		return errors.New("could not validate nil key. ensure one was specified")
// 	}

// 	//

// 	return nil
// }
