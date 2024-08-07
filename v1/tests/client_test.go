package v1

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	nexus "github.com/the-egg-corp/gonexus/v1"
)

var NexusClient = NewNexusClient()

func NewNexusClient() nexus.Client {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	return nexus.Client{
		APIKey: os.Getenv("NEXUS_KEY"),
	}
}
