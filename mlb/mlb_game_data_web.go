package mlb

type GameDataWeb struct {
	Details      DetailsWeb   `json:"details"`
	VisitingTeam VisitingTeam `json:"visiting_team"`
	HomeTeam     HomeTeam     `json:"home_team"`
}
