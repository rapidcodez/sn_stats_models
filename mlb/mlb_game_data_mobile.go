package snmodel

type MlbGameDataMobile struct {
	Details      WebDetails   `json:"details"`
	VisitingTeam VisitingTeam `json:"visiting_team"`
	HomeTeam     HomeTeam     `json:"home_team"`
}
