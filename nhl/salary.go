package nhl

type NhlPlayerSalary struct {
	StatsGlobalId  int    `json:"stats_global_id"`
	StatsTeamId    int    `json:"stats_team_id"`
	PlayerId       int    `json:"player_id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	FreeAgentType  int    `json:"free_agent_type"`
	Birthday       string `json:"birthday"`
	WaiversExempt  int    `json:"waivers_exempt"`
	SlideRisk      int    `json:"slide_risk"`
	TeamId         int    `json:"team_id"`
	Position       int    `json:"position"`
	CareerEarnings struct {
		Salary struct {
			Estimate int `json:"estimate"`
		} `json:"salary"`
	} `json:"career_earnings"`
	Salary struct {
		Salary       int `json:"salary"`
		SigningBonus int `json:"signing_bonus"`
		BaseSalary   int `json:"base_salary"`
		Caphit       struct {
			Initial            int     `json:"initial"`
			Retained           int     `json:"retained"`
			RetainedPercentage float64 `json:"retained_percentage"`
			Final              int     `json:"final"`
			CeilingPercentage  float64 `json:"ceiling_percentage"`
		} `json:"caphit"`
	} `json:"salary"`
	ActiveContract struct {
		FirstYear   int    `json:"first_year"`
		LastYear    int    `json:"last_year"`
		SigningDate string `json:"signing_date"`
		Type        int    `json:"type"`
		Value       int    `json:"value"`
	} `json:"active_contract"`
	Clauses map[string]struct {
		Nmc int `json:"nmc"`
		Ntc int `json:"ntc"`
	} `json:"clauses"`
}

type NhlTeamSalary struct {
	StatsTeamId               int     `json:"stats_team_id"`
	TeamId                    int     `json:"team_id"`
	Name                      string  `json:"name"`
	Logo                      string  `json:"logo"`
	LinkId                    string  `json:"link_id"`
	Abbreviation              string  `json:"abbreviation"`
	Active                    int     `json:"active"`
	ActiveRoster              int     `json:"active_roster"`
	StandardPlayerContracts   int     `json:"standard_player_contracts"`
	Injuries                  int     `json:"injuries"`
	CurrentCaphit             int     `json:"current_caphit"`
	CaphitSeasonAverage       int     `json:"caphit_season_average"`
	ProjectedLtirUsed         int     `json:"projected_ltir_used"`
	ProjectedCapspace         int     `json:"projected_capspace"`
	ProjectedCaphit           int     `json:"projected_caphit"`
	ProjectedCaphitPercentage float64 `json:"projected_caphit_percentage"`
	CurrentCapspace           int     `json:"current_capspace"`
	TradeDeadlinePassed       int     `json:"trade_deadline_passed"`
	TradeDeadlineCapspace     int     `json:"trade_deadline_capspace"`
	Postseason                int     `json:"postseason"`
	SalaryCap                 struct {
		Season  int `json:"season"`
		Ceiling int `json:"ceiling"`
		Floor   int `json:"floor"`
	} `json:"salary_cap"`
}
