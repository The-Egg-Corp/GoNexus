package v1

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

type User struct {
	UserID      int    `json:"user_id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	IsPremium   bool   `json:"is_premium"`
	IsSupporter bool   `json:"is_supporter"`
	Email       string `json:"email"`
	ProfileURL  string `json:"profile_url"`
}
