package v1

type EndorsementStatus string

const (
	Endorsed  EndorsementStatus = "Endorsed"
	Abstained EndorsementStatus = "Abstained"
)

type Endorsement struct {
	ModID      int               `json:"mod_id"`
	DomainName string            `json:"domain_name"`
	Date       string            `json:"date"`
	Version    *string           `json:"version"`
	Status     EndorsementStatus `json:"status"`
}

type EndorsementEvent struct {
	Message string            `json:"message"`
	Status  EndorsementStatus `json:"status"`
}
