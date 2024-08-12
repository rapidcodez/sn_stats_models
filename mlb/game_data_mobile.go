package mlb

type GameDataMobile struct {
	Details      DetailsMobile `json:"details"`
	VisitingTeam TeamMobile    `json:"visiting_team"`
	HomeTeam     TeamMobile    `json:"home_team"`
	Sequence     int64         `json:"sequence"`
}

type DetailsMobile struct {
	ID                 int          `json:"id"`
	IsTBD              bool         `json:"is_tbd"`
	SrGameUuid         string       `json:"sr_game_uuid"`
	LeagueShortName    string       `json:"league_short_name"`
	Status             string       `json:"status"`
	SeriesStatus       string       `json:"seriesStatus"`
	Timestamp          int          `json:"timestamp"`
	Inning             int          `json:"inning"`
	InningStatus       string       `json:"inning_status"`
	WinningPitcher     *Player      `json:"winning_pitcher,omitempty"`
	LosingPitcher      *Player      `json:"losing_pitcher,omitempty"`
	Outs               int          `json:"outs"`
	IsActive           bool         `json:"is_active"`
	LastScoreSequence  int64        `json:"score_sequence_last"`
	RunnerOnFirstBase  bool         `json:"runner_on_first_base"`
	RunnerOnSecondBase bool         `json:"runner_on_second_base"`
	RunnerOnThirdBase  bool         `json:"runner_on_third_base"`
	HomeSeriesWins     *int         `json:"home_series_wins"`
	VisitingSeriesWins *int         `json:"visiting_series_wins"`
	Round              *string      `json:"round"`
	SeriesWinner       *string      `json:"seriesWinner"`
	ScoreLogs          []ScoreLog   `json:"score_logs"`
	CurrentLine        *CurrentLine `json:"current_line,omitempty"`
}

type ScoreLog struct {
	ScoreSequence int    `json:"score_sequence"`
	ScoreHash     string `json:"score_hash"`
	HomeScore     int    `json:"home_score"`
	AwayScore     int    `json:"away_score"`
}

type TeamMobile struct {
	ID              int              `json:"id"`           // this will be deprecated
	TeamID          string           `json:"team_id"`      // this will be the config Team Id and will be the PK going forward
	SrTeamUUID      string           `json:"sr_team_uuid"` // this will be vendor SRspecific information. Each vendor will have their own reference
	Name            string           `json:"name"`
	ShortName       string           `json:"short_name"`
	StartingPitcher *StartingPitcher `json:"starting_pitcher,omitempty"`
	Score           int              `json:"score"`
	City            string           `json:"city"`
}
