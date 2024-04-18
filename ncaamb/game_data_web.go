package ncaamb

type GameDataWeb struct {
	Details      DetailsWeb `json:"details"`
	VisitingTeam TeamWeb    `json:"visiting_team"`
	HomeTeam     TeamWeb    `json:"home_team"`
}

type DetailsWeb struct {
	LeagueShortName  string `json:"league_short_name"`
	SrGameUuid       string `json:"sr_game_uuid"`
	Half             int    `json:"half"`
	Clock            string `json:"clock"`
	Timestamp        int    `json:"timestamp"`
	Status           string `json:"status"`
	Coverage         string `json:"coverage"`
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
	SrTeamUUID string     `json:"sr_team_uuid"`
	Name       string     `json:"name"`
	ShortName  string     `json:"short_name"`
	City       string     `json:"city"`
	ImageURL   string     `json:"image_url"`
	ImageURL90 string     `json:"image_url_90"`
	ImageURL59 string     `json:"image_url_59"`
	ImageURL25 string     `json:"image_url_25"`
	Score      int        `json:"score"`
	Division   Division   `json:"division"`
	Conference Conference `json:"conference"`
	Halfs      []Half     `json:"halfs"`
	GameStats  GameStats  `json:"game_stats"`
}
