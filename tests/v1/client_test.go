// TO RUN ALL TESTS: go test -timeout 30s github.com/the-egg-corp/gonexus/tests/v1 -v
package v1

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	v1 "github.com/the-egg-corp/gonexus/v1"
)

var NexusClient, InitClientError = NewNexusClient()

func NewNexusClient() (*v1.Client, error) {
	vars, err := godotenv.Read("../../.env")
	if err != nil {
		log.Fatalf("\nFailed to load required environment variables.\n%s", err)
	}

	key, found := vars["NEXUS_KEY"]
	if !found {
		log.Fatalf("\nCould not find required environment variable: NEXUS_KEY\n%s", err)
	}

	return v1.NewNexusClient(key)
}

func TestNewNexusClient(t *testing.T) {
	if NexusClient == nil {
		if InitClientError != nil {
			t.Fatalf("\nFailed to initialize the client.\n%s", InitClientError)
			return
		}

		t.Fatal("\nFailed to initialize the client but no error was provided.")
	}
}
