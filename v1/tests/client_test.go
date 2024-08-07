package v1

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	v1 "github.com/the-egg-corp/gonexus/v1"
)

var NexusClient, InitClientError = NewNexusClient()

func NewNexusClient() (*v1.Client, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	return v1.NewNexusClient(os.Getenv("NEXUS_KEY"))
}

func TestNewNexusClient(t *testing.T) {
	if NexusClient == nil {
		t.Fatal(InitClientError)
	}
}
