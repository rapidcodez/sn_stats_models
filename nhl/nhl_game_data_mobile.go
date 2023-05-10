package nhl

type GameDataMobile struct {
	Sequence     int64         `json:"sequence"`
	Details      DetailsMobile `json:"details"`
	VisitingTeam VisitingTeam  `json:"visiting_team"`
	HomeTeam     HomeTeam      `json:"home_team"`
}
