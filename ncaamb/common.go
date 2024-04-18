package ncaamb

type Half struct {
	Type     string `json:"type"`
	Number   int    `json:"number"`
	Sequence int    `json:"sequence"`
	Score    int    `json:"score"`
}

type Venue struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Capacity         int    `json:"capacity"`
	Address          string `json:"address"`
	City             string `json:"city"`
	State            string `json:"state"`
	Zip              string `json:"zip"`
	Country          string `json:"country"`
	LocationImageURL string `json:"location_image_url"`
	LocationImageMed string `json:"location_image_med"`
	LocationImageSml string `json:"location_image_sml"`
}

type Division struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type Conference struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
	ImageURL  string `json:"image_url"`
}

type GameStats struct {
	FieldGoalsMade       int     `json:"field_goals_made"`
	FieldGoalsAttempted  int     `json:"field_goals_att"`
	FieldGoalsPct        float32 `json:"field_goals_pct"`
	ThreePointsMade      int     `json:"three_points_made"`
	ThreePointsAttempted int     `json:"three_points_att"`
	ThreePointsPct       float32 `json:"three_points_pct"`
	FreeThrowsMade       int     `json:"free_throws_made"`
	FreeThrowsAttempted  int     `json:"free_throws_att"`
	FreeThrowsPct        float32 `json:"free_throws_pct"`
	OffensiveRebounds    int     `json:"offensive_rebounds"`
	DefensiveRebounds    int     `json:"defensive_rebounds"`
	Assists              int     `json:"assists"`
	Steals               int     `json:"steals"`
	Blocks               int     `json:"blocks"`
	PersonalFouls        int     `json:"personal_fouls"`
	Ejections            int     `json:"ejections"`
	Points               int     `json:"points"`
	TeamRebounds         int     `json:"team_rebounds"`
	FlagrantFouls        int     `json:"flagrant_fouls"`
	Turnovers            int     `json:"turnovers"`
	PlayerTechFouls      int     `json:"player_tech_fouls"`
	TeamTechFouls        int     `json:"team_tech_fouls"`
	CoachTechFouls       int     `json:"coach_tech_fouls"`
	PointsInPaint        int     `json:"points-in-paint"`
	BiggestLead          int     `json:"biggest_lead"`
	SecondChancePoints   int     `json:"second_chance_pts"`
	TeamTurnovers        int     `json:"team_turnovers"`
	PointsOffTurnovers   int     `json:"points_off_turnovers"`
}
