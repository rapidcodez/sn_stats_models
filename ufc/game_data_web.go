package ufc

import "time"

type GameDataWeb struct {
	Details      DetailsWeb `json:"details"`
	VisitingTeam Competitor `json:"visiting_team"`
	HomeTeam     Competitor `json:"home_team"`
}

type DetailsWeb struct {
	ID                  int       `json:"id"`
	LeagueShortName     string    `json:"league_short_name"`
	SrSportEventID      string    `json:"sr_sport_event_id"`
	SrCompetitionID     string    `json:"sr_competition_id"`
	Round               int       `json:"round"`
	Clock               string    `json:"clock"`
	TitleFight          bool      `json:"title_fight"`
	MainEvent           bool      `json:"main_event"`
	SrSeasonID          string    `json:"sr_season_id"`
	Status              string    `json:"status"`
	WeightClass         string    `json:"weight_class"`
	Method              string    `json:"method"`
	Timestamp           int       `json:"timestamp"`
	Venue               Venue     `json:"venue"`
	StartTime           time.Time `json:"start_time"`
	StartTimeConfirmed  bool      `json:"start_time_confirmed"`
	SportID             string    `json:"sport_id"`
	SportName           string    `json:"sport_name"`
	CategoryID          string    `json:"category_id"`
	CategoryName        string    `json:"category_name"`
	CompetitionID       string    `json:"competition_id"`
	CompetitionName     string    `json:"competition_name"`
	CompetitionParentID string    `json:"competition_parent_id"`
	SeasonID            string    `json:"season_id"`
	SeasonName          string    `json:"season_name"`
	SeasonStartDate     string    `json:"season_start_date"`
	SeasonEndDate       string    `json:"season_end_date"`
	SeasonYear          string    `json:"season_year"`
	CoverageIsLive      bool      `json:"coverage_is_live"`
	CoverageType        string    `json:"coverage_type"`
	FinalWinner         *string   `json:"finalWinner"`
}
