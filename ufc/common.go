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

type SportEventStatus struct {
	Status           string `json:"status"`
	MatchStatus      string `json:"match_status"`
	WinnerId         string `json:"winner_id"`
	FinalRound       int    `json:"final_round"`
	FinalRoundLength string `json:"final_round_length"`
	Method           string `json:"method"`
	Winner           string `json:"winner"`
	ScheduledLength  int    `json:"scheduled_length"`
	WeightClass      string `json:"weight_class"`
	TitleFight       bool   `json:"title_fight"`
	MainEvent        bool   `json:"main_event"`
}
