package v1

import (
	"errors"
	"fmt"

	"github.com/the-egg-corp/gonexus/util"
)

func (c Client) GetAllGames() ([]Game, error) {
	return util.JsonGetRequest[[]Game]("v1/games", c.APIKey)
}

func (c Client) GetGame(name string) (Game, error) {
	res, err := util.JsonGetRequest[Game](fmt.Sprint("v1/games/", name), c.APIKey)
	if res.ID == 0 && res.Name == "" {
		return res, errors.New("no game found with name: " + name)
	}

	return res, err
}
