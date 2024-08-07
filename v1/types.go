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
	apiKey           *string
}

type GameCategory struct {
	Name           string `json:"name"`
	CategoryID     int    `json:"category_id"`
	ParentCategory any    `json:"parent_category"` // Can be bool or int. Go pls implement unions :(
}

type Mod struct {
	Name                    string `json:"name"`
	Summary                 string `json:"summary"`
	Description             string `json:"description"`
	PictureURL              string `json:"picture_url"`
	ModDownloads            uint32 `json:"mod_downloads"`
	ModUniqueDownloads      uint32 `json:"mod_unique_downloads"`
	UID                     uint64 `json:"uid"`
	ModID                   uint32 `json:"mod_id"`
	GameID                  uint32 `json:"game_id"`
	AllowRating             bool   `json:"allow_rating"`
	DomainName              string `json:"domain_name"`
	CategoryID              uint16 `json:"category_id"`
	Version                 string `json:"version"`
	EndorsementCount        uint32 `json:"endorsement_count"`
	CreatedTimestamp        uint32 `json:"created_timestamp"`
	CreatedTime             string `json:"created_time"`
	UpdatedTimestamp        uint32 `json:"updated_timestamp"`
	UpdatedTime             string `json:"updated_time"`
	Author                  string `json:"author"`
	UploadedBy              string `json:"uploaded_by"`
	UploadedUsersProfileURL string `json:"uploaded_users_profile_url"`
	ContainsAdultContent    bool   `json:"contains_adult_content"`
	Status                  string `json:"status"`
	Available               bool   `json:"available"`
	User                    struct {
		MemberID      uint32 `json:"member_id"`
		MemberGroupID uint32 `json:"member_group_id"`
		Name          string `json:"name"`
	} `json:"user"`
	Endorsement struct {
		EndorseStatus string  `json:"endorse_status"`
		Timestamp     *string `json:"timestamp"`
		Version       *string `json:"version"`
	} `json:"endorsement"`
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

type Status string

const (
	Endorsed  Status = "Endorsed"
	Abstained Status = "Abstained"
)

type Endorsement struct {
	ModID      int     `json:"mod_id"`
	DomainName string  `json:"domain_name"`
	Date       string  `json:"date"`
	Version    *string `json:"version"`
	Status     Status  `json:"status"`
}
