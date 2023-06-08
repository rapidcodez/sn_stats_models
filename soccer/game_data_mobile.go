package soccer

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam TeamMobile    `json:"visiting_team"`
	HomeTeam     TeamMobile    `json:"home_team"`
}

type DetailsMobile struct {
	LeagueShortName    string `json:"league_short_name"`
	ID                 int    `json:"id"`
	SrGameUuid         string `json:"sr_game_uuid"`
	Period             int    `json:"period"`
	IsActive           bool   `json:"is_active"`
	Clock              string `json:"clock"`          // "match_clock":"45:00"
	Time               int    `json:"time"`           // "match_time":45,
	StoppageClock      string `json:"stoppage_clock"` // "stoppage_time_clock":"45:00"
	StoppageTime       int    `json:"stoppage_time"`  // "stoppage_time":45,
	Timestamp          int    `json:"timestamp"`
	Status             string `json:"status"`
	IsHalfTime         bool   `json:"ishalftime"`
	HasAggregateScores bool   `json:"hasAggregateScores"`
	Overtime           int    `json:"overtime"`
	PenaltyKicks       bool   `json:"penalty_kicks"`
}

type TeamMobile struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	City           string `json:"city"`
	Score          int    `json:"score"`
	PenaltyScore   int    `json:"penalty_score"`
	AggregateScore *int   `json:"aggregate_score"`
}
