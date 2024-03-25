package ufc

type Venue struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	CityName         string `json:"city_name"`
	CountryName      string `json:"country_name"`
	CountryCode      string `json:"country_code"`
	Timezone         string `json:"timezone"`
	State            string `json:"state"`
	LocationImageURL string `json:"location_image_url"`
	LocationImageMed string `json:"location_image_med"`
	LocationImageSml string `json:"location_image_sml"`
}

type Competitor struct {
	Id               int                  `json:"id"`
	CompetitorId     string               `json:"team_id"`          // To be sourced from config when it supports competitor data, serialized as team_id for backwards compat reasons.
	SrCompetitorUUID string               `json:"sr_competitor_id"` // Vendor-specific id.
	Name             string               `json:"name"`
	Abbreviation     string               `json:"abbreviation"`
	Gender           string               `json:"gender"`
	Qualifier        string               `json:"qualifier"`
	AgeGroup         string               `json:"age_group"`
	Country          string               `json:"country"`
	CountryCode      string               `json:"country_code"`
	Virtual          bool                 `json:"virtual"`
	Statistics       CompetitorStatistics `json:"statistics"`
	IsWinner         bool                 `json:"is_winner"`
}

type CompetitorStatistics struct {
	Control                     string  `json:"control"`
	Knockdowns                  int     `json:"knockdowns"`
	SignificantStrikePercentage float32 `json:"significant_strike_percentage"`
	SignificantStrikes          int     `json:"significant_strikes"`
	SignificantStrikesAttempted int     `json:"significant_strikes_attempted"`
	SubmissionAttempts          int     `json:"submission_attempts"`
	TakedownPercentage          float32 `json:"takedown_percentage"`
	Takedowns                   int     `json:"takedowns"`
	TakedownsAttempted          int     `json:"takedowns_attempted"`
	TotalStrikePercentage       float32 `json:"total_strike_percentage"`
	TotalStrikes                int     `json:"total_strikes"`
	TotalStrikesAttempted       int     `json:"total_strikes_attempted"`
}
