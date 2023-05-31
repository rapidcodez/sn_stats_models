package mlb

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam TeamMobile    `json:"visiting_team"`
	HomeTeam     TeamMobile    `json:"home_team"`
}

type DetailsMobile struct {
	ID                 int    `json:"id"`
	SrGameUuid         string `json:"sr_game_uuid"`
	LeagueShortName    string `json:"league_short_name"`
	Status             string `json:"status"`
	Timestamp          int    `json:"timestamp"`
	Inning             int    `json:"inning"`
	InningStatus       string `json:"inning_status"`
	Outs               int    `json:"outs"`
	IsActive           bool   `json:"is_active"`
	RunnerOnFirstBase  bool   `json:"runner_on_first_base"`
	RunnerOnSecondBase bool   `json:"runner_on_second_base"`
	RunnerOnThirdBase  bool   `json:"runner_on_third_base"`
}

type TeamMobile struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Score     int    `json:"score"`
	City      string `json:"city"`
}
