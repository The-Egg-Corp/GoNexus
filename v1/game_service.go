package v1

import (
	"errors"
)

type GameService struct {
	client *NexusClient
}

// Initializes a new [GameService] that is reponsible for handling game-related requests like getting all games or a specific game.
func NewGameService(client *NexusClient) *GameService {
	return &GameService{client: client}
}

func (gs *GameService) GetAllGames() (*[]Game, error) {
	return jsonGetRequest[[]Game]("v1/games", gs.client)
}

func (gs *GameService) GetGame(name string) (*Game, error) {
	game, err := jsonGetRequest[Game]("v1/games/"+name, gs.client)
	if err != nil {
		return nil, err
	}

	if game.ID == 0 && game.Name == "" {
		return game, errors.New("no game found with name: " + name)
	}

	return game, err
}
