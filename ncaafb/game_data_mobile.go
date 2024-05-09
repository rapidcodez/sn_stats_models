package ncaafb

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam TeamMobile    `json:"visiting_team"`
	HomeTeam     TeamMobile    `json:"home_team"`
}

type TeamMobile struct {
	Id         int    `json:"id"`
	TeamID     string `json:"team_id"`
	Name       string `json:"name"`
	SrTeamUUID string `json:"sr_team_uuid"`
	ShortName  string `json:"short_name"`
	HasBall    bool   `json:"has_ball"`
	City       string `json:"city"`
	Rank       string `json:"rank"`
	Score      int    `json:"score"`
}

type DetailsMobile struct {
	Id                          int    `json:"id"`
	SrGameUuid                  string `json:"sr_game_uuid"`
	LeagueShortName             string `json:"league_short_name"`
	Quarter                     int    `json:"quarter"`
	IsActive                    bool   `json:"is_active"`
	IsConferenceChampionship    bool   `json:"is_conference_championship"`
	ConferenceChampionshipTitle string `json:"conference_championship_title"`
	IsPlayoffGame               bool   `json:"is_playoff_game"`
	PlayoffGameType             string `json:"playoff_game_type"`
	PlayoffGameTitle            string `json:"playoff_game_title"`
	IsBowlGame                  bool   `json:"is_bowl_game"`
	BowlGameTitle               string `json:"bowl_game_title"`
	Timestamp                   int    `json:"timestamp"`
	Status                      string `json:"status"`
	Overtime                    int    `json:"overtime"`
	IsIntermission              bool   `json:"is_intermission"`
	Sequence                    int64  `json:"sequence"`
	ScoreSequence               int64  `json:"score_sequence"`
	Downs                       int    `json:"downs"`
	Distance                    int    `json:"distance"`
	Clock                       string `json:"clock"`
}
