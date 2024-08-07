package v1

import (
	"testing"

	"github.com/the-egg-corp/gonexus/util"
)

func TestAllGames(t *testing.T) {
	games, err := NexusClient.GetGame("lethalcompany")
	if err != nil {
		t.Fatal()
	}

	util.PrettyPrint(games)
}
