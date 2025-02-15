package v1

import "fmt"

type ModService struct {
	client *Client
}

// Initializes a new GameService, reponsible for handling mod-related requests.
func NewModService(client *Client) *ModService {
	// TODO: Maybe accept game here instead of in each method? Assess if the coupling would be justified.
	return &ModService{client: client}
}

// ------------------------------------------------------------------------------------------------------------
// TODO: Make these more versatile by accepting game name and mod id directly instead of structs?

func (ms *ModService) GetModByID(game Game, id uint32) (*Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%v", game.DomainName, id)

	// TODO: Implement a way to detect 404/not found to return nil (with error) instead of an empty struct.
	return jsonGetRequest[Mod](endpoint, ms.client)
}

func (ms *ModService) GetModChangelogs(game Game, mod Mod) (*Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%v/changelogs", game.DomainName, mod.ModID)
	return jsonGetRequest[Mod](endpoint, ms.client)
}

// Retrieves 10 trending mods for the given game.
func (ms *ModService) Trending(game Game) (*[]Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/trending", game.DomainName)
	return jsonGetRequest[[]Mod](endpoint, ms.client)
}

// Retrieves the 10 latest added mods for the given game.
func (ms *ModService) LatestAdded(game Game) (*[]Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/latest_added", game.DomainName)
	return jsonGetRequest[[]Mod](endpoint, ms.client)
}

// Retrieves the 10 latest updated mods for the given game.
func (ms *ModService) LatestUpdated(game Game) (*[]Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/latest_updated", game.DomainName)
	return jsonGetRequest[[]Mod](endpoint, ms.client)
}

// Retrieves a list of mods that were updated in the given period, with timestamps of their last update.
// This list is cached by NexusMods for 5 minutes.
func (ms *ModService) UpdatedWithinPeriod(game Game, period ModUpdatePeriod) (*[]ModUpdateInfo, error) {
	endpoint := fmt.Sprintf("v1/games/%s/updated?period=%s", game.DomainName, period)
	return jsonGetRequest[[]ModUpdateInfo](endpoint, ms.client)
}

// Sends a POST request, indicating the current user has endorsed (liked) this mod.
func (ms *ModService) EndorseMod(game Game, mod Mod) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/endorse", game.DomainName, mod.ModID)
	return jsonPostRequest[EndorsementEvent](endpoint, ms.client)
}

func (ms *ModService) AbstainEndorsement(game Game, mod Mod) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/abstain", game.DomainName, mod.ModID)
	return jsonPostRequest[EndorsementEvent](endpoint, ms.client)
}

func (ms *ModService) GetModFiles(game Game, mod Mod) (*ModFilesResponse, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/files", game.DomainName, mod.ModID)
	return jsonPostRequest[ModFilesResponse](endpoint, ms.client)
}

func (ms *ModService) GetModFileByID(game Game, mod Mod, fileId uint16) (*ModFile, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/files/%d", game.DomainName, mod.ModID, fileId)
	return jsonPostRequest[ModFile](endpoint, ms.client)
}

// ------------------------------------------------------------------------------------------------------------
