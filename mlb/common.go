package mlb

type ImageUrls struct {
	Lg string `json:"lg"`
	Md string `json:"md"`
	Sm string `json:"sm"`
	Xs string `json:"xs"`
}

type Story struct {
	Type     string `json:"type"`
	Headline string `json:"headline"`
	Content  string `json:"content"`
}
type Umpires struct {
	Position string `json:"position"`
	FullName string `json:"full_name"`
}
type SeasonSeries struct {
	Date              string `json:"date"`
	Location          string `json:"location"`
	VisitingTeamScore int    `json:"visiting_team_score"`
	HomeTeamScore     int    `json:"home_team_score"`
	ID                int    `json:"id"`
	VisitingTeamID    int    `json:"visiting_team_id"`
	HomeTeamID        int    `json:"home_team_id"`
}
type CurrentAtBat struct {
	Batter  Player  `json:"batter"`
	Pitcher Player  `json:"pitcher"`
	Pitch   []Pitch `json:"pitch"`
}
type Pitch struct {
	PitchType     string `json:"pitch_type"`
	PitchTypeText string `json:"pitch_type_text"`
	PitchTypeDesc string `json:"pitch_type_desc"`
	PitchVelocity int    `json:"pitch_velocity"`
	PitchX        int    `json:"pitch_x"`
	PitchY        string `json:"pitch_y"`
	Strikes       int    `json:"strikes"`
	Balls         int    `json:"balls"`
	CurrentCount  string `json:"current_count"`
}
type CurrentInningPbp struct {
	EventID       int       `json:"event_id"`
	EventCode     int       `json:"event_code"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	ImageURL      string    `json:"image_url,omitempty"`
	TeamShortName string    `json:"team_short_name"`
	TeamLogoURL   string    `json:"team_logo_url"`
	Description   string    `json:"description"`
	PlayType      string    `json:"play_type"`
	Pitches       []Pitch   `json:"pitches,omitempty"`
	SnEventID     string    `json:"sn_event_id"`
	ImageUrls     ImageUrls `json:"image_urls,omitempty"`
	PlayerID      int       `json:"player_id"`
}

type RunnersOnBase struct {
	First struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"first"`
	Second struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"second"`
	Third struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"third"`
}
type LastOut struct {
	EventID       int       `json:"event_id"`
	EventCode     int       `json:"event_code"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	ImageURL      string    `json:"image_url"`
	TeamShortName string    `json:"team_short_name"`
	TeamLogoURL   string    `json:"team_logo_url"`
	Description   string    `json:"description"`
	PlayType      string    `json:"play_type"`
	Pitches       []Pitch   `json:"pitches"`
	SnEventID     string    `json:"sn_event_id"`
	ImageUrls     ImageUrls `json:"image_urls"`
	PlayerID      int       `json:"player_id"`
}
type CurrentLine struct {
	FavID      int     `json:"fav_id"`
	Name       string  `json:"name"`
	FavPoints  float64 `json:"fav_points"`
	FavMoney   int     `json:"fav_money"`
	HomeMoney  int     `json:"home_money"`
	AwayMoney  int     `json:"away_money"`
	Total      int     `json:"total"`
	OverMoney  int     `json:"over_money"`
	UnderMoney int     `json:"under_money"`
	DrawMoney  int     `json:"draw_money"`
}
type OpeningLine struct {
	FavID         int     `json:"fav_id"`
	Name          string  `json:"name"`
	FavPoints     float64 `json:"fav_points"`
	FavMoney      int     `json:"fav_money"`
	UnderdogMoney int     `json:"underdog_money"`
	HomeMoney     int     `json:"home_money"`
	AwayMoney     int     `json:"away_money"`
	Total         int     `json:"total"`
	OverMoney     int     `json:"over_money"`
	UnderMoney    int     `json:"under_money"`
	DrawMoney     int     `json:"draw_money"`
}

type Division struct {
	ID             int    `json:"id"`
	SrDivisionUuid string `json:"sr_division_uuid"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	Rank           int    `json:"rank"`
}
type Conference struct {
	ID               int    `json:"id"`
	SrConferenceUuid string `json:"sr_conference_uuid"`
	Name             string `json:"name"`
	ShortName        string `json:"short_name"`
	Rank             int    `json:"rank"`
	ImageURL         string `json:"image_url"`
	Color            string `json:"color"`
}
type GameStats struct {
	Hits              int `json:"hits"`
	Errors            int `json:"errors"`
	HomeRuns          int `json:"home_runs"`
	Steals            int `json:"steals"`
	Walks             int `json:"walks"`
	StrikeOuts        int `json:"strike_outs"`
	Runs              int `json:"runs"`
	RunnersLeftOnBase int `json:"runners_left_on_base"`
}
type SeasonStats struct {
	RunsPerGame        float64 `json:"runs_per_game"`
	TeamBattingAverage float64 `json:"team_batting_average"`
	HomeRuns           int     `json:"home_runs"`
	EarnedRunAverage   float64 `json:"earned_run_average"`
	FieldPercent       float64 `json:"field_percent"`
	Wins               int     `json:"wins"`
	Losses             int     `json:"losses"`
	GamesBack          float64 `json:"games_back"`
	OnBasePercentage   float64 `json:"on_base_percentage"`
	Saves              int     `json:"saves"`
	Errors             int     `json:"errors"`
}

type SeasonLeader struct {
	ID           int       `json:"id"`
	SrPlayerUUID string    `json:"sr_player_uuid"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	ImageURL     string    `json:"image_url"`
	Total        string    `json:"total"`
	ImageUrls    ImageUrls `json:"image_urls"`
}

type Player struct {
	ID            int    `json:"id"`
	SrPlayerUUID  string `json:"sr_player_uuid"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ImageURL      string `json:"image_url"`
	Position      string `json:"position"`
	ShortPosition string `json:"short_position"`
	Number        int    `json:"number"`

	AtBats             int       `json:"at_bats,omitempty"`
	Runs               int       `json:"runs,omitempty"`
	HomeRuns           int       `json:"home_runs,omitempty"`
	Hits               int       `json:"hits,omitempty"`
	RunsBattedIn       int       `json:"runs_batted_in,omitempty"`
	Walks              int       `json:"walks,omitempty"`
	StrikeOuts         int       `json:"strike_outs,omitempty"`
	LeftOnBase         int       `json:"left_on_base,omitempty"`
	BattingAverage     string    `json:"batting_average,omitempty"`
	StolenBases        int       `json:"stolen_bases,omitempty"`
	Uniform            int       `json:"uniform,omitempty"`
	Batting            string    `json:"batting,omitempty"`
	SeasonHomeRuns     int       `json:"season_home_runs,omitempty"`
	SeasonRunsBattedIn int       `json:"season_runs_batted_in,omitempty"`
	Doubles            int       `json:"doubles,omitempty"`
	Triples            int       `json:"triples,omitempty"`
	OnBasePercentage   string    `json:"on_base_percentage,omitempty"`
	IsSeasonStat       bool      `json:"is_season_stat,omitempty"`
	StartedGame        bool      `json:"started_game,omitempty"`
	SacrificeFlies     int       `json:"sacrifice_flies,omitempty"`
	HitByPitch         int       `json:"hit_by_pitch,omitempty"`
	ImageUrls          ImageUrls `json:"image_urls,omitempty"`
	DisplayID          int       `json:"display_id,omitempty"`
	DisplayName        string    `json:"display_name,omitempty"`
	PlayerCodeID       int       `json:"player_code_id,omitempty"`
	BattingSlot        int       `json:"batting_slot,omitempty"`
	SluggingPercentage string    `json:"slugging_percentage,omitempty"`
	RunsBattedInSeason int       `json:"runs_batted_in_season,omitempty"`

	//pitchers
	InningsPitched            string `json:"innings_pitched,omitempty"`
	EarnedRuns                int    `json:"earned_runs,omitempty"`
	HomeRunsAllowed           int    `json:"home_runs_allowed,omitempty"`
	PitchCount                int    `json:"pitch_count,omitempty"`
	Strikes                   int    `json:"strikes,omitempty"`
	BattersFaced              int    `json:"batters_faced,omitempty"`
	GroundBalls               int    `json:"ground_balls,omitempty"`
	FlyBalls                  int    `json:"fly_balls,omitempty"`
	WildPitches               int    `json:"wild_pitches,omitempty"`
	EarnedRunAverage          string `json:"earned_run_average,omitempty"`
	Wins                      int    `json:"wins,omitempty"`
	Losses                    int    `json:"losses,omitempty"`
	Saves                     int    `json:"saves,omitempty"`
	Home                      bool   `json:"home,omitempty"`
	PitcherID                 int    `json:"pitcher_id,omitempty"`
	PitcherFirstName          string `json:"pitcher_first_name,omitempty"`
	PitcherLastName           string `json:"pitcher_last_name,omitempty"`
	Throwing                  string `json:"throwing,omitempty"`
	Holds                     int    `json:"holds,omitempty"`
	GameCreditedWin           bool   `json:"game_credited_win,omitempty,omitempty"`
	BlownSaves                int    `json:"blown_saves,omitempty"`
	SequenceNumber            int    `json:"sequence_number,omitempty"`
	OpponentBattingAverage    string `json:"opponent_batting_average,omitempty"`
	StrikeOutsSeason          int    `json:"strike_outs_season,omitempty"`
	WalksHitsPerInningAverage string `json:"walks_hits_per_inning_average,omitempty"`
	GameCreditedLoss          bool   `json:"game_credited_loss,omitempty,omitempty"`
}

type Batting struct {
	Doubles                 []string `json:"doubles"`
	HomeRuns                []string `json:"home_runs"`
	RunsBattedIn            []string `json:"runs_batted_in"`
	TwoOutRBI               []string `json:"two_out_RBI"`
	GroundedIntoDoublePlays []string `json:"grounded_into_double_plays"`
	TeamRisp                string   `json:"team_risp"`
	TeamLob                 int      `json:"team_lob"`
}
type Pitching struct {
	BattersFaced      []string `json:"batters_faced"`
	GroundFlyBalls    []string `json:"ground_fly_balls"`
	HitByPitch        []string `json:"hit_by_pitch"`
	CsSsFbIps         []string `json:"cs_ss_fb_ips"`
	PitchCountStrikes []string `json:"pitch_count_strikes"`
}

type SeasonStatsStr struct {
	RunsPerGame        string `json:"runs_per_game"`
	TeamBattingAverage string `json:"team_batting_average"`
	HomeRuns           int    `json:"home_runs"`
	EarnedRunAverage   string `json:"earned_run_average"`
	Wins               int    `json:"wins"`
	Losses             int    `json:"losses"`
	FieldPercent       string `json:"field_percent"`
	OnBasePercentage   string `json:"on_base_percentage"`
	Saves              int    `json:"saves"`
	Errors             int    `json:"errors"`
	GamesBehind        string `json:"games_behind"`
}
type TeamStandings struct {
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	GamesBehind   string  `json:"games_behind"`
	WcGamesBehind string  `json:"wc_games_behind"`
	WinPercentage float64 `json:"win_percentage"`
	Last10Record  string  `json:"last_10_record"`
	HomeRecord    string  `json:"home_record"`
	EastRecord    string  `json:"east_record"`
	CentralRecord string  `json:"central_record"`
	WestRecord    string  `json:"west_record"`
	RoadRecord    string  `json:"road_record"`
	RunsScored    int     `json:"runs_scored"`
	RunsAllowed   int     `json:"runs_allowed"`
	StreakGames   string  `json:"streak_games"`
	HomeWins      int     `json:"home_wins"`
	HomeLosses    int     `json:"home_losses"`
	AwayWins      int     `json:"away_wins"`
	AwayLosses    int     `json:"away_losses"`
	Last10Wins    int     `json:"last_10_wins"`
	Last10Losses  int     `json:"last_10_losses"`
}

type Injuries struct {
	PlayerID  int    `json:"player_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
	Status    string `json:"status"`
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"`
}

type StartingPitcher struct {
	ID               int       `json:"id"`
	SrPlayerUUID     string    `json:"sr_player_uuid"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	ImageURL         string    `json:"image_url"`
	Position         string    `json:"position"`
	ShortPosition    string    `json:"short_position"`
	Number           int       `json:"number"`
	HomeRunsAllowed  int       `json:"home_runs_allowed"`
	EarnedRunAverage string    `json:"earned_run_average"`
	Wins             int       `json:"wins"`
	Losses           int       `json:"losses"`
	Throwing         string    `json:"throwing"`
	ImageUrls        ImageUrls `json:"image_urls"`
	InningsPitched   string    `json:"innings_pitched"`
	StrikeOuts       int       `json:"strike_outs"`
	GamesPlayed      int       `json:"games_played"`
	Whip             float64   `json:"whip"`
}

type Fielding struct {
	DoublePlays []string `json:"double_plays"`
	Errors      []string `json:"errors"`
}
type Plays struct {
	EventID     int    `json:"event_id"`
	EventCode   int    `json:"event_code"`
	Description string `json:"description"`
	IsOut       string `json:"is_out"`
}
type Inning struct {
	Inning             int     `json:"inning"`
	InningName         string  `json:"inning_name"`
	VisitingTeamScore  int     `json:"visiting_team_score,omitempty"`
	VisitingTeamHits   int     `json:"visiting_team_hits,omitempty"`
	VisitingTeamErrors int     `json:"visiting_team_errors,omitempty"`
	Plays              []Plays `json:"plays"`
	HomeTeamScore      int     `json:"home_team_score,omitempty"`
	HomeTeamHits       int     `json:"home_team_hits,omitempty"`
	HomeTeamErrors     int     `json:"home_team_errors,omitempty"`
}
