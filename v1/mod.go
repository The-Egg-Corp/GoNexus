package v1

type ModUpdatePeriod string

const (
	MOD_UPDATE_PERIOD_ONE_DAY   ModUpdatePeriod = "1d"
	MOD_UPDATE_PERIOD_ONE_WEEK  ModUpdatePeriod = "1w"
	MOD_UPDATE_PERIOD_ONE_MONTH ModUpdatePeriod = "1m"
)

type ModUpdateInfo struct {
	ModId uint32 `json:"mod_id"`
	// TODO: These could possibly be uint.
	LatestFileUpdate  int32 `json:"latest_file_update"`
	LatestModActivity int32
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

// Alias for [Mod.ContainsAdultContent].
func (mod Mod) IsNSFW() bool {
	return mod.ContainsAdultContent
}
