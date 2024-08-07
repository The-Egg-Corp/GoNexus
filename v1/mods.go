package v1

import (
	"fmt"

	"github.com/the-egg-corp/gonexus/util"
)

func (game Game) getMultipleMods(endpoint string) ([]Mod, error) {
	err := ValidateKey(game.apiKey)
	if err != nil {
		return nil, err
	}

	return util.JsonGetRequest[[]Mod](fmt.Sprintf("v1/mods/%s/%s", game.DomainName, endpoint), *game.apiKey)
}

func (game Game) GetMod(id uint32) (Mod, error) {
	return Mod{}, nil
}

// Retrieves the 10 latest mods that were added for this game.
func (game Game) LatestAddedMods() ([]Mod, error) {
	return game.getMultipleMods("latest_updated")
}

func (game Game) LatestUpdatedMods() ([]Mod, error) {
	return game.getMultipleMods("latest_updated")
}

// Retrieves 10 mods that are trending for this game.
func (game Game) TrendingMods() ([]Mod, error) {
	return game.getMultipleMods("trending")
}

// Alias for github.com/the-egg-corp/gonexus/v1/#Mod.ContainsAdultContent.
func (mod Mod) IsNSFW() bool {
	return mod.ContainsAdultContent
}

// Sends a POST request, indicating the current user has endorsed (liked) this mod.
func (mod Mod) Endorse() any {
	return EndorseMod(mod)
}

func EndorseMod(mod Mod) any {
	return nil
}
