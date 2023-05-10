package nhl

type GameDataWeb struct {
	Sequence     int64        `json:"sequence"`
	Details      DetailsWeb   `json:"details"`
	VisitingTeam VisitingTeam `json:"visiting_team"`
	HomeTeam     HomeTeam     `json:"home_team"`
	Periods      []Periods    `json:"periods"`
}
