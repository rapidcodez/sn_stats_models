package ncaafb

type Quarter struct {
	Plays []struct {
		VisitingTeamScore int    `json:"visiting_team_score"`
		HomeTeamScore     int    `json:"home_team_score"`
		TeamId            int    `json:"team_id"`
		Time              string `json:"time"`
		Play              string `json:"play"`
	} `json:"plays"`
	VisitingTeamScore int `json:"visiting_team_score"`
	HomeTeamScore     int `json:"home_team_score"`
}
