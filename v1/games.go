package v1

import (
	"errors"
	"fmt"

	"github.com/the-egg-corp/gonexus/util"
)

func (c Client) GetAllGames() ([]Game, error) {
	return util.JsonGetRequest[[]Game]("v1/games", apiKey)
}

func (c Client) GetGame(name string) (Game, error) {
	game, err := util.JsonGetRequest[Game](fmt.Sprint("v1/games/", name), apiKey)
	if game.ID == 0 && game.Name == "" {
		return game, errors.New("no game found with name: " + name)
	}

	return game, err
}
