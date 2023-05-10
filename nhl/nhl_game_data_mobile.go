package sn_stats_models

type NhlGameDataMobile struct {
	Sequence     int64        `json:"sequence"`
	Details      Details      `json:"details"`
	VisitingTeam VisitingTeam `json:"visiting_team"`
	HomeTeam     HomeTeam     `json:"home_team"`
	Periods      []Periods    `json:"periods"`
}
