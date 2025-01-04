package v1

import (
	"errors"
)

func (c Client) GetAllGames() (*[]Game, error) {
	return jsonGetRequest[[]Game]("v1/games", c)
}

func (c Client) GetGame(name string) (*Game, error) {
	game, err := jsonGetRequest[Game]("v1/games/"+name, c)
	if err != nil {
		return nil, err
	}

	if game.ID == 0 && game.Name == "" {
		return game, errors.New("no game found with name: " + name)
	}

	return game, err
}
