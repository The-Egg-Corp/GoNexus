package v1

import (
	"testing"

	v1 "github.com/the-egg-corp/gonexus/v1"
)

const testGame = "lethalcompany"

var GameService = v1.NewGameService(NexusClient)

func TestGetAllGames(t *testing.T) {
	games, err := GameService.GetAllGames()
	if err != nil {
		t.Fatal("failed to get all games\n", err)
	}

	if games == nil {
		t.Fatalf("failed to get game: %s. response was nil but no error was provided", testGame)
	}

	//util.PrettyPrint(games)
}

func TestGetGame(t *testing.T) {
	game, err := GameService.GetGame(testGame)
	if err != nil {
		t.Fatalf("failed to get game: %s\n%v", testGame, err)
	}

	if game == nil {
		t.Fatalf("failed to get game: %s. response was nil but no error was provided", testGame)
	}

	//util.PrettyPrint(game)
}

func TestGetGame_NonExistentGameShouldError(t *testing.T) {
	_, err := GameService.GetGame("someBogusGame")
	if err == nil {
		t.Fatalf("bogus game did not error with 404 (Not Found)")
	}
}
