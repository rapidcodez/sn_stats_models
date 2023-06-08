package jhl

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam TeamMobile    `json:"visiting_team"`
	HomeTeam     TeamMobile    `json:"home_team"`
}
type DetailsMobile struct {
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
