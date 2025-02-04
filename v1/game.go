package v1

import "fmt"

type Game struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	ForumURL         string         `json:"forum_url"`
	NexusURL         string         `json:"nexusmods_url"`
	Genre            string         `json:"genre"`
	FileCount        uint32         `json:"file_count"`
	Downloads        uint32         `json:"downloads"` // Surely no game will ever get 4B dls.
	DomainName       string         `json:"domain_name"`
	ApprovedDate     int            `json:"approved_date"`
	FileViews        int            `json:"file_views"`
	Authors          int            `json:"authors"`
	FileEndorsements int            `json:"file_endorsements"`
	Mods             int            `json:"mods"`
	Categories       []GameCategory `json:"categories"`
}

type GameCategory struct {
	Name           string `json:"name"`
	CategoryID     int    `json:"category_id"`
	ParentCategory any    `json:"parent_category"` // Can be bool or int. Go pls implement unions :(
}

// --------------------------------------------------------------------------------
// TODO: Replace these with a mod service?

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

// Sends a POST request, indicating the current user has endorsed (liked) this mod.
func (game Game) EndorseMod(mod Mod, client Client) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/endorse", game.DomainName, mod.ModID)
	return jsonPostRequest[EndorsementEvent](endpoint, client)
}

func (game Game) AbstainEndorsement(mod Mod, client Client) (*EndorsementEvent, error) {
	endpoint := fmt.Sprintf("v1/games/%s/mods/%d/abstain", game.DomainName, mod.ModID)
	return jsonPostRequest[EndorsementEvent](endpoint, client)
}

// --------------------------------------------------------------------------------
