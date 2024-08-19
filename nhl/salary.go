package nhl

type NhlPlayerSalary struct {
	Data struct {
		Lastname       string `json:"lastname"`
		Position       string `json:"position"`
		Birthdate      string `json:"birthdate"`
		Firstname      string `json:"firstname"`
		PlayerId       int    `json:"player_id"`
		SlideRisk      string `json:"slide_risk"`
		NhlTeamId      string `json:"nhl_team_id"`
		StatsTeamId    int    `json:"stats_team_id"`
		WaiversExempt  int    `json:"waivers_exempt"`
		CareerEarnings struct {
			Salary struct {
				Estimate int `json:"estimate"`
			} `json:"salary"`
		} `json:"career_earnings"`
		FreeAgentType string `json:"free_agent_type"`
		StatsGlobalId int    `json:"stats_global_id"`
	} `json:"data"`
	Meta struct {
		Code         int    `json:"code"`
		ErrorMessage string `json:"error_message"`
	} `json:"meta"`
	Salary struct {
		Caphit struct {
			Final              int     `json:"final"`
			Initial            string  `json:"initial"`
			Retained           int     `json:"retained"`
			CeilingPercentage  string  `json:"ceiling_percentage"`
			RetainedPercentage float64 `json:"retained_percentage"`
		} `json:"caphit"`
		Salary       string `json:"salary"`
		BaseSalary   string `json:"base_salary"`
		SigningBonus string `json:"signing_bonus"`
	} `json:"salary"`
	ActiveContract struct {
		Type        string `json:"type"`
		Value       string `json:"value"`
		LastYear    int    `json:"last_year"`
		FirstYear   int    `json:"first_year"`
		SigningDate string `json:"signing_date"`
	} `json:"active_contract"`
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
