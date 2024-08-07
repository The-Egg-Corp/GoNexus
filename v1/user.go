package v1

import "github.com/the-egg-corp/gonexus/util"

func ValidateUser(key string) (User, error) {
	return util.JsonGetRequest[User]("v1/users/validate", key)
}

func (c Client) ValidateUser() (User, error) {
	return ValidateUser(c.apiKey)
}

func (c Client) GetEndorsements() ([]Endorsement, error) {
	return util.JsonGetRequest[[]Endorsement]("v1/user/endorsements", c.apiKey)
}
