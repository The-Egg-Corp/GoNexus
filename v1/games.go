package v1

import (
	"fmt"

	"github.com/the-egg-corp/gonexus/util"
)

type Client struct {
	APIKey string
}

func (c Client) GetAllGames() ([]Game, error) {
	return util.JsonGetRequest[[]Game]("v1/games", c.APIKey)
}

func (c Client) GetGame(name string) (Game, error) {
	return util.JsonGetRequest[Game](fmt.Sprint("v1/games/", name), c.APIKey)
}
