// This package is for interacting with version 1.0 of the [NexusMods API].
//
// [NexusMods API]: https://app.swaggerhub.com/apis-docs/NexusMods/nexus-mods_public_api_params_in_form_data/1.0
package v1

import (
	"fmt"

	"github.com/the-egg-corp/gonexus/util"
)

// TODO: Support Single Sign-On (SSO) for application use. The current client will only work for personal keys afaik?

// The client that is used to send requests to the Nexus API. Since it holds the API key, it will usually be required by methods so they can return valid data -
// though this is almost always done internally and you should rarely ever need to pass the client around except for when creating a new service.
//
// IMPORTANT:
// While in most cases a personal API key is fine, it is recommended you learn about the rate-limits employed by NexusMods and read their
// [Acceptable Use Policy]. For use in a public facing a application, you may be required to register your app and use the Single Sign-On (SSO) system instead.
//
// [Acceptable Use Policy]: https://help.nexusmods.com/article/114-api-acceptable-use-policy
type NexusClient struct {
	apiKey string
}

// Creates a new [NexusClient] using the input API key. Before a client is instantiated, a request is made to validate the API key (see [VALIDATE_USER_ENDPOINT]).
// In the case that it is invalid/empty, a nil client is returned with an appropriate error.
// Otherwise, everything succeeded and the returned client is assumed to be valid.
func NewNexusClient(key string) (*NexusClient, error) {
	user, err := SendValidateUserRequest(key)
	if err != nil {
		return nil, err
	}

	// Couldn't get a user back from the API given the provided key.
	if user == nil || user.Name == "" {
		return nil, fmt.Errorf("error creating client: invalid api key provided")
	}

	return &NexusClient{apiKey: key}, nil
}

func jsonGetRequest[T interface{}](endpoint string, client *NexusClient) (*T, error) {
	if client == nil {
		return nil, fmt.Errorf("error sending GET request to: %s. initialized client is nil", endpoint)
	}

	res, err := util.GetRequest(endpoint, client.apiKey)
	if err != nil {
		return nil, err
	}

	return util.ParseJsonBody[T](*res)
}

func jsonPostRequest[T interface{}](endpoint string, client *NexusClient) (*T, error) {
	if client == nil {
		return nil, fmt.Errorf("error sending POST request to %s. initialized client is nil", endpoint)
	}

	res, err := util.PostRequest(endpoint, client.apiKey)
	if err != nil {
		return nil, err
	}

	return util.ParseJsonBody[T](*res)
}
