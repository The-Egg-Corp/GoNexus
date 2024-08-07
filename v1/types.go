package v1

type Game struct {
	ID               string         `json:"id"`
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
	ParentCategory bool   `json:"parent_category"`
}
