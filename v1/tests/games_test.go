package v1

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/the-egg-corp/gonexus/util"
	nexus "github.com/the-egg-corp/gonexus/v1"
)

var Client = NewNexusClient()

func NewNexusClient() nexus.Client {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	return nexus.Client{
		APIKey: os.Getenv("API_KEY"),
	}
}

func TestAllGames(t *testing.T) {
	games, err := Client.GetGame("lethalcompany")
	if err != nil {
		t.Fatal()
	}

	util.PrettyPrint(games)
}
