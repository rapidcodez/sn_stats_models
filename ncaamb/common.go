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
type DetailsMobile struct {
	ID                 int    `json:"id"`
	LeagueShortName    string `json:"league_short_name"`
	SrGameUuid         string `json:"sr_game_uuid"`
	Half               int    `json:"half"`
	Clock              string `json:"clock"`
	Timestamp          int    `json:"timestamp"`
	Status             string `json:"status"`
	IsActive           bool   `json:"is_active"`
	Overtime           int    `json:"overtime"`
	Sequence           int    `json:"sequence"`
	IsConferenceGame   bool   `json:"is_conference_game"`
	Timeout            string `json:"timeout"`
	TimesTied          int    `json:"times_tied"`
	LeadChanges        int    `json:"lead_changes"`
	EntryMode          string `json:"entry_mode"`
	Attendance         int    `json:"attendance"`
	Venue              Venue  `json:"venue"`
	HomeSeriesWins     int    `json:"home_series_wins"`
	VisitingSeriesWins int    `json:"visiting_series_wins"`
}

type TeamMobile struct {
	ID         int       `json:"id"`
	SrTeamUUID string    `json:"sr_team_uuid"`
	Name       string    `json:"name"`
	ShortName  string    `json:"short_name"`
	City       string    `json:"city"`
	Score      int       `json:"score"`
	Halfs      []Half    `json:"halfs"`
	GameStats  GameStats `json:"game_stats"`
}

type DetailsWeb struct {
	LeagueShortName  string `json:"league_short_name"`
	SrGameUuid       string `json:"sr_game_uuid"`
	Half             int    `json:"half"`
	Clock            string `json:"clock"`
	Timestamp        int    `json:"timestamp"`
	Status           string `json:"status"`
	IsActive         bool   `json:"is_active"`
	Overtime         int    `json:"overtime"`
	Sequence         int    `json:"sequence"`
	Type             string `json:"type"`
	IsConferenceGame bool   `json:"is_conference_game"`
	Timeout          string `json:"timeout"`
	TimesTied        int    `json:"times_tied"`
	LeadChanges      int    `json:"lead_changes"`
	EntryMode        string `json:"entry_mode"`
	Attendance       int    `json:"attendance"`
	Venue            Venue  `json:"venue"`
}

type TeamWeb struct {
	SrTeamUUID string    `json:"sr_team_uuid"`
	Name       string    `json:"name"`
	ShortName  string    `json:"short_name"`
	City       string    `json:"city"`
	ImageURL   string    `json:"image_url"`
	ImageURL90 string    `json:"image_url_90"`
	ImageURL59 string    `json:"image_url_59"`
	ImageURL25 string    `json:"image_url_25"`
	Score      int       `json:"score"`
	Halfs      []Half    `json:"halfs"`
	GameStats  GameStats `json:"game_stats"`
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
