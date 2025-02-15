package v1

type ColourScheme struct {
	ID              uint16 `json:"id"`
	Name            string `json:"name"`
	PrimaryColour   string `json:"primary_colour"`
	SecondaryColour string `json:"secondary_colour"`
	DarkerColour    string `json:"darker_colour"`
}

func (c Client) GetAllColourSchemes() (*[]ColourScheme, error) {
	return jsonPostRequest[[]ColourScheme]("v1/colourschemes", &c)
}
