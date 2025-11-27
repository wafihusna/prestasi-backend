package model

import "time"

type AchievementDetails struct {
	CompetitionName   string    `json:"competitionName"`
	CompetitionLevel  string    `json:"competitionLevel"`
	Rank              int       `json:"rank"`
	MedalType         string    `json:"medalType"`

	PublicationType   string    `json:"publicationType"`
	PublicationTitle  string    `json:"publicationTitle"`
	Authors           []string  `json:"authors"`
	Publisher         string    `json:"publisher"`
	ISSN              string    `json:"issn"`

	OrganizationName  string    `json:"organizationName"`
	Position          string    `json:"position"`
	Period            struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"period"`

	CertificationName    string    `json:"certificationName"`
	IssuedBy             string    `json:"issuedBy"`
	CertificationNumber  string    `json:"certificationNumber"`
	ValidUntil           time.Time `json:"validUntil"`

	EventDate         time.Time         `json:"eventDate"`
	Location          string            `json:"location"`
	Organizer         string            `json:"organizer"`
	Score             int               `json:"score"`
	CustomFields      map[string]any    `json:"customFields"`
}