package v1

type ModService struct {
	client *Client
}

// NewGameService initializes a new GameService.
func NewModService(client *Client) *ModService {
	return &ModService{client: client}
}
