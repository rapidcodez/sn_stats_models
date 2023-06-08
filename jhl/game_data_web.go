package jhl

type GameDataWeb struct {
	Details      DetailsWeb `json:"details"`
	VisitingTeam TeamWeb    `json:"visiting_team"`
	HomeTeam     TeamWeb    `json:"home_team"`
	Periods      []Periods  `json:"periods"`
}
type DetailsWeb struct {
	LeagueShortName string `json:"league_short_name"`
	Period          int    `json:"period"`
	IsActive        bool   `json:"is_active"`
	Overtime        int    `json:"overtime"`
	HasShootout     bool   `json:"has_shootout"`
	Clock           string `json:"clock"`
	ID              int    `json:"id"`
	SrGameUuid      string `json:"sr_game_uuid"`
	Timestamp       int    `json:"timestamp"`
	Status          string `json:"status"`
}
