package v1

import (
	"testing"

	"github.com/the-egg-corp/gonexus/util"
	v1 "github.com/the-egg-corp/gonexus/v1"
)

var ModService = v1.NewModService(NexusClient)

func TestGetGameMod(t *testing.T) {
	game, err := GameService.GetGame("HogwartsLegacy")
	if err != nil {
		t.Fatal("error getting game\n", err)
	}

	mod, err := ModService.GetModByID(1863, *game)
	if err != nil {
		t.Fatal("error getting mod\n", err)
	}

	util.PrettyPrint(mod)
}
