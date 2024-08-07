package v1

import "errors"

type Client struct {
	apiKey string
}

func NewNexusClient(key string) (*Client, error) {
	user, err := ValidateUser(key)

	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, errors.New("invalid api key")
	}

	return &Client{apiKey: key}, nil
}

func ValidateKey(key *string) error {
	if key == nil {
		return errors.New("couldn't send request. no api key provided")
	}

	return nil
}
