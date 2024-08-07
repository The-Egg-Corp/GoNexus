package v1

import "github.com/the-egg-corp/gonexus/util"

func (c Client) ValidateUser() (User, error) {
	return util.JsonGetRequest[User]("v1/users/validate", c.APIKey)
}
