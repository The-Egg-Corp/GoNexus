package v1

import "fmt"

type ModService struct {
	client *Client
}

// Initializes a new [ModService] that is reponsible for handling mod-related requests like
// endorsing or abstaining, getting changelogs, getting files for specific mods etc.
func NewModService(client *Client) *ModService {
	// TODO: Maybe accept game here instead of in each method? Assess if the coupling would be justified.
	return &ModService{client: client}
}

func (ms *ModService) GetModByID(gameName string, modId uint32) (*Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%v", gameName, modId)
	return jsonGetRequest[Mod](endpoint, ms.client)
}

func (ms *ModService) GetModChangelogs(gameName string, modId uint32) (*ModChangelogsResponse, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%v/changelogs", gameName, modId)
	return jsonGetRequest[ModChangelogsResponse](endpoint, ms.client)
}

// Retrieves 10 trending mods for the given game.
func (ms *ModService) Trending(gameName string) (*[]Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/trending", gameName)
	return jsonGetRequest[[]Mod](endpoint, ms.client)
}

// Retrieves the 10 latest added mods for the given game.
func (ms *ModService) LatestAdded(gameName string) (*[]Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/latest_added", gameName)
	return jsonGetRequest[[]Mod](endpoint, ms.client)
}

// Retrieves the 10 latest updated mods for the given game.
func (ms *ModService) LatestUpdated(gameName string) (*[]Mod, error) {
	endpoint := fmt.Sprintf("v1/games/%s/latest_updated", gameName)
	return jsonGetRequest[[]Mod](endpoint, ms.client)
}

// Retrieves a list of mods that were updated in the given period, with timestamps of their last update.
// This list is cached by NexusMods for 5 minutes.
func (ms *ModService) UpdatedWithinPeriod(gameName string, period ModUpdatePeriod) (*[]ModUpdateInfo, error) {
	endpoint := fmt.Sprintf("v1/games/%s/updated?period=%s", gameName, period)
	return jsonGetRequest[[]ModUpdateInfo](endpoint, ms.client)
}

// Sends a POST request, indicating the current user has endorsed (liked) this mod.
func (ms *ModService) EndorseMod(gameName string, modId uint32) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/endorse", gameName, modId)
	return jsonPostRequest[EndorsementEvent](endpoint, ms.client)
}

func (ms *ModService) AbstainEndorsement(gameName string, modId uint32) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/abstain", gameName, modId)
	return jsonPostRequest[EndorsementEvent](endpoint, ms.client)
}

func (ms *ModService) GetModFiles(gameName string, modId uint32) (*ModFilesResponse, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/files", gameName, modId)
	return jsonPostRequest[ModFilesResponse](endpoint, ms.client)
}

func (ms *ModService) GetModFileByID(gameName string, modId uint32, fileId uint32) (*ModFile, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/files/%d", gameName, modId, fileId)
	return jsonPostRequest[ModFile](endpoint, ms.client)
}
