package v1

import (
	"errors"

	"github.com/the-egg-corp/gonexus/util"
)

const VALIDATE_USER_ENDPOINT = "v1/users/validate"

type User struct {
	UserID      int    `json:"user_id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	IsPremium   bool   `json:"is_premium"`
	IsSupporter bool   `json:"is_supporter"`
	Email       string `json:"email"`
	ProfileURL  string `json:"profile_url"`
}

// Sends a request to [VALIDATE_USER_ENDPOINT] to check the validity of the given key.
//
// This request does not count towards hourly rate limits.
func SendValidateUserRequest(key string) (*User, error) {
	if key == "" {
		return nil, errors.New("error validating user: specified key is empty")
	}

	res, err := util.GetRequest(VALIDATE_USER_ENDPOINT, key)
	if err != nil {
		return nil, err
	}

	return util.ParseJsonBody[User](*res)
}

// Checks whether the api key of this client is valid.
func (c NexusClient) ValidateUser() (*User, error) {
	return SendValidateUserRequest(c.apiKey)
}

// Retrieves endorsements for current user (this client).
func (c NexusClient) GetEndorsements() (*[]Endorsement, error) {
	return jsonGetRequest[[]Endorsement]("v1/user/endorsements", &c)
}

// Fetches all mods being tracked by the current user.
func (c NexusClient) GetTrackedMods() (*[]Mod, error) {
	return jsonGetRequest[[]Mod]("v1/user/tracked_mods", &c)
}
