package ncaamb

type GameDataWeb struct {
	Details      DetailsBase `json:"details"`
	VisitingTeam Competitor  `json:"visiting_team"`
	HomeTeam     Competitor  `json:"home_team"`
}
