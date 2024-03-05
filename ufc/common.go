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
	Id               int    `json:"id"`
	CompetitorId     string `json:"team_id"`          // To be sourced from config when it supports competitor data, serialized as team_id for backwards compat reasons.
	SrCompetitorUUID string `json:"sr_competitor_id"` // Vendor-specific id.
	Name             string `json:"name"`
	Abbreviation     string `json:"abbreviation"`
	Qualifier        string `json:"qualifier"`
	Gender           string `json:"gender"`
	// Additional data we pull from the competitor profile
	Nickname              string             `json:"nickname"`
	FightingOutOfLocation CompetitorLocation `json:"fighting_out_of_location"`
	BirthLocation         CompetitorLocation `json:"birth_location"`
	BirthDate             string             `json:"birth_date"`
	Reach                 string             `json:"reach"`
	Height                string             `json:"height"`
	Weight                string             `json:"weight"`
	Record                CompetitorRecord   `json:"record"`
}

type CompetitorLocation struct {
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type CompetitorRecord struct {
	Wins       int `json:"wins"`
	Draws      int `json:"draws"`
	Losses     int `json:"losses"`
	NoContests int `json:"no_contest"`
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
