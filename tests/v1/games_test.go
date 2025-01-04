package v1

import (
	"testing"

	"github.com/the-egg-corp/gonexus/util"
)

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
