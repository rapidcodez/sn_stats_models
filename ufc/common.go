package ufc

import "time"

type DetailsBase struct {
	ID                  int              `json:"id"`
	LeagueShortName     string           `json:"league_short_name"`
	SrSportEventID      string           `json:"sr_sport_event_id"`
	SrCompetitionID     string           `json:"sr_competition_id"`
	SrSeasonID          string           `json:"sr_season_id"`
	Status              string           `json:"status"`
	Timestamp           int              `json:"timestamp"`
	Venue               Venue            `json:"venue"`
	StartTime           time.Time        `json:"start_time"`
	StartTimeConfirmed  bool             `json:"start_time_confirmed"`
	SportID             string           `json:"sport_id"`
	SportName           string           `json:"sport_name"`
	CategoryID          string           `json:"category_id"`
	CategoryName        string           `json:"category_name"`
	CompetitionID       string           `json:"competition_id"`
	CompetitionName     string           `json:"competition_name"`
	CompetitionParentID string           `json:"competition_parent_id"`
	SeasonID            string           `json:"season_id"`
	SeasonName          string           `json:"season_name"`
	SeasonStartDate     string           `json:"season_start_date"`
	SeasonEndDate       string           `json:"season_end_date"`
	SeasonYear          string           `json:"season_year"`
	CoverageIsLive      bool             `json:"coverage_is_live"`
	CoverageType        string           `json:"coverage_type"`
	SportEventStatus    SportEventStatus `json:"sport_event_status"`
}

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
