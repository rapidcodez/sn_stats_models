package ufc

import "time"

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam Competitor    `json:"visiting_team"`
	HomeTeam     Competitor    `json:"home_team"`
}

type DetailsMobile struct {
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
