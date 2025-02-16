package v1

import (
	"fmt"

	"github.com/the-egg-corp/gonexus/util"
)

type Client struct {
	apiKey string
}

// NexusMods API Permalink:
// https://app.swaggerhub.com/apis-docs/NexusMods/nexus-mods_public_api_params_in_form_data/1.0

func NewNexusClient(key string) (*Client, error) {
	user, err := SendValidateUserRequest(key)
	if err != nil {
		return nil, err
	}

	// Couldn't get a user back from the API given the provided key.
	if user == nil || user.Name == "" {
		return nil, fmt.Errorf("error creating client: invalid api key provided")
	}

	return &Client{apiKey: key}, nil
}

func jsonGetRequest[T interface{}](endpoint string, client *Client) (*T, error) {
	if client == nil {
		return nil, fmt.Errorf("error sending GET request to: %s. initialized client is nil", endpoint)
	}

	res, err := util.GetRequest(endpoint, client.apiKey)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() == 404 {
		return nil, fmt.Errorf("error during GET request to: %s. resource not found (404)", endpoint)
	}

	return util.ParseJsonBody[T](*res)
}

func jsonPostRequest[T interface{}](endpoint string, client *Client) (*T, error) {
	if client == nil {
		return nil, fmt.Errorf("error sending POST request to %s. initialized client is nil", endpoint)
	}

	res, err := util.PostRequest(endpoint, client.apiKey)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() == 404 {
		return nil, fmt.Errorf("error during POST request to: %s. resource not found (404)", endpoint)
	}

	return util.ParseJsonBody[T](*res)
}
