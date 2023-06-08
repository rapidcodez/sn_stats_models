package nba

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam TeamMobile    `json:"visiting_team"`
	HomeTeam     TeamMobile    `json:"home_team"`
}

type DetailsMobile struct {
	LeagueShortName    string `json:"league_short_name"`
	ID                 int    `json:"id"`
	SrGameUuid         string `json:"sr_game_uuid"`
	Quarter            int    `json:"quarter"`
	IsActive           bool   `json:"is_active"`
	Clock              string `json:"clock"`
	Timestamp          int    `json:"timestamp"`
	Status             string `json:"status"`
	Overtime           int    `json:"overtime"`
	HomeSeriesWins     int    `json:"home_series_wins"`
	VisitingSeriesWins int    `json:"visiting_series_wins"`
}

type TeamMobile struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	City      string `json:"city"`
	Score     int    `json:"score"`
}
