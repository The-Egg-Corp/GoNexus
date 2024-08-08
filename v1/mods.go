package v1

import (
	"fmt"

	"github.com/the-egg-corp/gonexus/util"
)

func (game Game) getMultipleMods(endpoint string) ([]Mod, error) {
	err := ValidateKey(&apiKey)
	if err != nil {
		return nil, err
	}

	return util.JsonGetRequest[[]Mod](fmt.Sprintf("v1/mods/%s/%s", game.DomainName, endpoint), apiKey)
}

func (game Game) GetMod(id uint32) (Mod, error) {
	return Mod{}, nil
}

func (game Game) Updated() {

}

// Retrieves the 10 latest mods that were added for this game.
func (game Game) LatestAdded() ([]Mod, error) {
	return game.getMultipleMods("latest_updated")
}

func (game Game) LatestUpdated() ([]Mod, error) {
	return game.getMultipleMods("latest_updated")
}

// Retrieves 10 mods that are trending for this game.
func (game Game) Trending() ([]Mod, error) {
	return game.getMultipleMods("trending")
}

// Alias for github.com/the-egg-corp/gonexus/v1/#Mod.ContainsAdultContent.
func (mod Mod) IsNSFW() bool {
	return mod.ContainsAdultContent
}

// Sends a POST request, indicating the current user has endorsed (liked) this mod.
func (game Game) EndorseMod(mod Mod) (EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/endorse", game.DomainName, mod.ModID)
	return util.JsonPostRequest[EndorsementEvent](endpoint, apiKey)
}

func (game Game) AbstainEndorsement(mod Mod) (EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/abstain", game.DomainName, mod.ModID)
	return util.JsonPostRequest[EndorsementEvent](endpoint, apiKey)
}
