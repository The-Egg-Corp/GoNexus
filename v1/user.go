package v1

import (
	"errors"

	"github.com/the-egg-corp/gonexus/util"
)

func ValidateUser(key string) (*User, error) {
	if key == "" {
		return nil, errors.New("error validating user: specified key is empty")
	}

	res, err := util.GetRequest("v1/users/validate", key)
	if err != nil {
		return nil, err
	}

	return util.ParseJsonBody[User](*res)
}

func (c Client) ValidateUser() (*User, error) {
	return ValidateUser(c.apiKey)
}

func (c Client) GetEndorsements() (*[]Endorsement, error) {
	return jsonGetRequest[[]Endorsement]("v1/user/endorsements", c)
}
