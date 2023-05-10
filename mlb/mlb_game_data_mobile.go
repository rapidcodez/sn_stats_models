package mlb

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam VisitingTeam  `json:"visiting_team"`
	HomeTeam     HomeTeam      `json:"home_team"`
}
