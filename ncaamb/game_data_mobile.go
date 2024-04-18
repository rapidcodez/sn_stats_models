package ncaamb

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam TeamMobile    `json:"visiting_team"`
	HomeTeam     TeamMobile    `json:"home_team"`
}

type DetailsMobile struct {
	ID                 int    `json:"id"`
	LeagueShortName    string `json:"league_short_name"`
	SrGameUuid         string `json:"sr_game_uuid"`
	Half               int    `json:"half"`
	Clock              string `json:"clock"`
	Timestamp          int    `json:"timestamp"`
	Status             string `json:"status"`
	GameTitle          string `json:"game_title,omitempty"`
	Coverage           string `json:"coverage"`
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
	HomeSeriesWins     *int   `json:"home_series_wins"`
	VisitingSeriesWins *int   `json:"visiting_series_wins"`
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
