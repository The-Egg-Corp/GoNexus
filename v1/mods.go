package v1

import (
	"fmt"
)

func (game Game) getMultipleMods(endpoint string, client Client) (*[]Mod, error) {
	return jsonGetRequest[[]Mod](fmt.Sprintf("v1/mods/%s/%s", game.DomainName, endpoint), client)
}

func (game Game) GetMod(id uint32) (Mod, error) {
	return Mod{}, nil
}

func (game Game) Updated() {

}

// Retrieves the 10 latest mods that were added for this game.
func (game Game) LatestAdded(client Client) (*[]Mod, error) {
	return game.getMultipleMods("latest_updated", client)
}

func (game Game) LatestUpdated(client Client) (*[]Mod, error) {
	return game.getMultipleMods("latest_updated", client)
}

// Retrieves 10 mods that are trending for this game.
func (game Game) Trending(client Client) (*[]Mod, error) {
	return game.getMultipleMods("trending", client)
}

// Alias for github.com/the-egg-corp/gonexus/v1/#Mod.ContainsAdultContent.
func (mod Mod) IsNSFW() bool {
	return mod.ContainsAdultContent
}

// Sends a POST request, indicating the current user has endorsed (liked) this mod.
func (game Game) EndorseMod(mod Mod, client Client) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/endorse", game.DomainName, mod.ModID)
	return jsonPostRequest[EndorsementEvent](endpoint, client)
}

func (game Game) AbstainEndorsement(mod Mod, client Client) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/abstain", game.DomainName, mod.ModID)
	return jsonPostRequest[EndorsementEvent](endpoint, client)
}
