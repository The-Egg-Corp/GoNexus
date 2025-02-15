package v1

type ModUpdatePeriod string

const (
	MOD_UPDATE_PERIOD_ONE_DAY   ModUpdatePeriod = "1d"
	MOD_UPDATE_PERIOD_ONE_WEEK  ModUpdatePeriod = "1w"
	MOD_UPDATE_PERIOD_ONE_MONTH ModUpdatePeriod = "1m"
)

type ModUpdateInfo struct {
	ModID uint32 `json:"mod_id"`
	// TODO: These could possibly be uint.
	LatestFileUpdate  int32 `json:"latest_file_update"`
	LatestModActivity int32 `json:"latest_mod_activity"`
}

type ModFilesResponse struct {
	Files       []ModFile `json:"files"`
	FileUpdates []any     `json:"file_updates"` // TODO: Figure out exactly what this does/returns.
}

type ModFile struct {
	Name                 string   `json:"name"`
	Version              string   `json:"version"`
	ModVersion           string   `json:"mod_version"`
	ID                   []uint32 `json:"id"` // Usually a pair digits, ex: [6279, 5113]. Yet to see negatives - uint32 should suffice for most mod files.
	UID                  int64    `json:"uid"`
	FileID               uint32   `json:"file_id"`
	FileName             string   `json:"file_name"`
	CategoryID           uint16   `json:"category_id"`
	CategoryName         string   `json:"category_name"`
	IsPrimary            bool     `json:"is_primary"`
	UploadedTimestamp    int64    `json:"uploaded_timestamp"`
	UploadedTime         string   `json:"uploaded_time"`
	ExternalVirusScanURL string   `json:"external_virus_scan_url"`
	ChangelogHTML        string   `json:"changelog_html"`
	ContentPreviewLink   string   `json:"content_preview_link"`
	Description          string   `json:"description"`
	Size                 uint64   `json:"size"` // Size is always same as size_kb I believe?
	SizeKB               uint64   `json:"size_kb"`
	SizeBytes            uint64   `json:"size_bytes"`
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
