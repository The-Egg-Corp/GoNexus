package v1

import (
	"testing"

	"github.com/the-egg-corp/gonexus/util"
	v1 "github.com/the-egg-corp/gonexus/v1"
)

var GameService = v1.NewGameService(NexusClient)

func TestGetAllGames(t *testing.T) {
	games, err := GameService.GetAllGames()
	if err != nil {
		t.Fatal("error getting all games\n", err)
	}

	util.PrettyPrint(games)
}

func TestGetGame(t *testing.T) {
	game, err := GameService.GetGame("lethalcompany")
	if err != nil {
		t.Fatal("error getting game\n", err)
	}

	util.PrettyPrint(game)
}
