package sn_stats_models

type MlbGameDataWeb struct {
	Details      Details      `json:"details"`
	VisitingTeam VisitingTeam `json:"visiting_team"`
	HomeTeam     HomeTeam     `json:"home_team"`
}
