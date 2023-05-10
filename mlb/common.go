package sn_stats_models

type ImageUrl struct {
	Lg string `json:"lg"`
	Md string `json:"md"`
	Sm string `json:"sm"`
	Xs string `json:"xs"`
}
type WinningPitcher struct {
	ID                        int      `json:"id"`
	FirstName                 string   `json:"first_name"`
	LastName                  string   `json:"last_name"`
	ImageURL                  string   `json:"image_url"`
	Position                  string   `json:"position"`
	ShortPosition             string   `json:"short_position"`
	Number                    int      `json:"number"`
	InningsPitched            string   `json:"innings_pitched"`
	Hits                      int      `json:"hits"`
	Runs                      int      `json:"runs"`
	EarnedRuns                int      `json:"earned_runs"`
	HomeRunsAllowed           int      `json:"home_runs_allowed"`
	PitchCount                int      `json:"pitch_count"`
	Strikes                   int      `json:"strikes"`
	BattersFaced              int      `json:"batters_faced"`
	GroundBalls               int      `json:"ground_balls"`
	FlyBalls                  int      `json:"fly_balls"`
	WildPitches               int      `json:"wild_pitches"`
	EarnedRunAverage          string   `json:"earned_run_average"`
	Walks                     int      `json:"walks"`
	StrikeOuts                int      `json:"strike_outs"`
	Wins                      int      `json:"wins"`
	Losses                    int      `json:"losses"`
	Saves                     int      `json:"saves"`
	Home                      bool     `json:"home"`
	PitcherID                 int      `json:"pitcher_id"`
	PitcherFirstName          string   `json:"pitcher_first_name"`
	PitcherLastName           string   `json:"pitcher_last_name"`
	PitcherNumber             int      `json:"pitcher_number"`
	Uniform                   int      `json:"uniform"`
	Throwing                  string   `json:"throwing"`
	IsSeasonStat              bool     `json:"is_season_stat"`
	Holds                     int      `json:"holds"`
	GameCreditedWin           bool     `json:"game_credited_win"`
	BlownSaves                int      `json:"blown_saves"`
	ImageUrls                 ImageUrl `json:"image_urls"`
	DisplayID                 int      `json:"display_id"`
	DisplayName               string   `json:"display_name"`
	PlayerCodeID              int      `json:"player_code_id"`
	SequenceNumber            int      `json:"sequence_number"`
	OpponentBattingAverage    string   `json:"opponent_batting_average"`
	StrikeOutsSeason          int      `json:"strike_outs_season"`
	WalksHitsPerInningAverage string   `json:"walks_hits_per_inning_average"`
}
type LosingPitcher struct {
	ID                        int      `json:"id"`
	FirstName                 string   `json:"first_name"`
	LastName                  string   `json:"last_name"`
	ImageURL                  string   `json:"image_url"`
	Position                  string   `json:"position"`
	ShortPosition             string   `json:"short_position"`
	Number                    int      `json:"number"`
	InningsPitched            string   `json:"innings_pitched"`
	Hits                      int      `json:"hits"`
	Runs                      int      `json:"runs"`
	EarnedRuns                int      `json:"earned_runs"`
	HomeRunsAllowed           int      `json:"home_runs_allowed"`
	PitchCount                int      `json:"pitch_count"`
	Strikes                   int      `json:"strikes"`
	BattersFaced              int      `json:"batters_faced"`
	GroundBalls               int      `json:"ground_balls"`
	FlyBalls                  int      `json:"fly_balls"`
	WildPitches               int      `json:"wild_pitches"`
	EarnedRunAverage          string   `json:"earned_run_average"`
	Walks                     int      `json:"walks"`
	StrikeOuts                int      `json:"strike_outs"`
	Wins                      int      `json:"wins"`
	Losses                    int      `json:"losses"`
	Saves                     int      `json:"saves"`
	Home                      bool     `json:"home"`
	PitcherID                 int      `json:"pitcher_id"`
	PitcherFirstName          string   `json:"pitcher_first_name"`
	PitcherLastName           string   `json:"pitcher_last_name"`
	PitcherNumber             int      `json:"pitcher_number"`
	Uniform                   int      `json:"uniform"`
	Throwing                  string   `json:"throwing"`
	IsSeasonStat              bool     `json:"is_season_stat"`
	Holds                     int      `json:"holds"`
	GameCreditedLoss          bool     `json:"game_credited_loss"`
	BlownSaves                int      `json:"blown_saves"`
	ImageUrls                 ImageUrl `json:"image_urls"`
	DisplayID                 int      `json:"display_id"`
	DisplayName               string   `json:"display_name"`
	PlayerCodeID              int      `json:"player_code_id"`
	SequenceNumber            int      `json:"sequence_number"`
	OpponentBattingAverage    string   `json:"opponent_batting_average"`
	StrikeOutsSeason          int      `json:"strike_outs_season"`
	WalksHitsPerInningAverage string   `json:"walks_hits_per_inning_average"`
}
type SavingPitcher struct {
	ID                        int      `json:"id"`
	FirstName                 string   `json:"first_name"`
	LastName                  string   `json:"last_name"`
	ImageURL                  string   `json:"image_url"`
	Position                  string   `json:"position"`
	ShortPosition             string   `json:"short_position"`
	Number                    int      `json:"number"`
	InningsPitched            string   `json:"innings_pitched"`
	Hits                      int      `json:"hits"`
	Runs                      int      `json:"runs"`
	EarnedRuns                int      `json:"earned_runs"`
	HomeRunsAllowed           int      `json:"home_runs_allowed"`
	PitchCount                int      `json:"pitch_count"`
	Strikes                   int      `json:"strikes"`
	BattersFaced              int      `json:"batters_faced"`
	GroundBalls               int      `json:"ground_balls"`
	FlyBalls                  int      `json:"fly_balls"`
	WildPitches               int      `json:"wild_pitches"`
	EarnedRunAverage          string   `json:"earned_run_average"`
	Walks                     int      `json:"walks"`
	StrikeOuts                int      `json:"strike_outs"`
	Wins                      int      `json:"wins"`
	Losses                    int      `json:"losses"`
	Saves                     int      `json:"saves"`
	Home                      bool     `json:"home"`
	PitcherID                 int      `json:"pitcher_id"`
	PitcherFirstName          string   `json:"pitcher_first_name"`
	PitcherLastName           string   `json:"pitcher_last_name"`
	PitcherNumber             int      `json:"pitcher_number"`
	Uniform                   int      `json:"uniform"`
	Throwing                  string   `json:"throwing"`
	IsSeasonStat              bool     `json:"is_season_stat"`
	Holds                     int      `json:"holds"`
	GameCreditedSave          bool     `json:"game_credited_save"`
	BlownSaves                int      `json:"blown_saves"`
	ImageUrls                 ImageUrl `json:"image_urls"`
	DisplayID                 int      `json:"display_id"`
	DisplayName               string   `json:"display_name"`
	PlayerCodeID              int      `json:"player_code_id"`
	SequenceNumber            int      `json:"sequence_number"`
	OpponentBattingAverage    string   `json:"opponent_batting_average"`
	StrikeOutsSeason          int      `json:"strike_outs_season"`
	WalksHitsPerInningAverage string   `json:"walks_hits_per_inning_average"`
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
type CurrentLine struct {
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
type WebDetails struct {
	ID                   int            `json:"id"`
	SrGameUuid           string         `json:"sr_game_uuid"`
	LeagueShortName      string         `json:"league_short_name"`
	Status               string         `json:"status"`
	Timestamp            int            `json:"timestamp"`
	Type                 string         `json:"type"`
	IsAllStar            bool           `json:"is_all_star"`
	IsPreseason          bool           `json:"is_preseason"`
	Inning               int            `json:"inning"`
	InningStatus         string         `json:"inning_status"`
	Outs                 int            `json:"outs"`
	IsActive             bool           `json:"is_active"`
	RunnerOnFirstBase    bool           `json:"runner_on_first_base"`
	RunnerOnSecondBase   bool           `json:"runner_on_second_base"`
	RunnerOnThirdBase    bool           `json:"runner_on_third_base"`
	Attendance           int            `json:"attendance"`
	Location             string         `json:"location"`
	LocationImageURL     string         `json:"location_image_url"`
	WinningPitcher       WinningPitcher `json:"winning_pitcher"`
	LosingPitcher        LosingPitcher  `json:"losing_pitcher"`
	SavingPitcher        SavingPitcher  `json:"saving_pitcher"`
	CurrentBatter        int            `json:"current_batter"`
	LastPlay             string         `json:"last_play"`
	LocationImageMed     string         `json:"location_image_med"`
	LocationImageSml     string         `json:"location_image_sml"`
	ExtraInnings         bool           `json:"extra_innings"`
	SeriesMatchupsResult string         `json:"series_matchups_result"`
	SeasonSeries         []SeasonSeries `json:"season_series"`
	CurrentAtBat         CurrentAtBat   `json:"current_at_bat"`
	GameDuration         string         `json:"game_duration"`
	City                 string         `json:"city"`
	Country              string         `json:"country"`
	Conference           string         `json:"conference"`
}
type Details struct {
	ID                   int            `json:"id"`
	SrGameUuid           string         `json:"sr_game_uuid"`
	LeagueShortName      string         `json:"league_short_name"`
	Status               string         `json:"status"`
	Timestamp            int            `json:"timestamp"`
	Type                 string         `json:"type"`
	IsAllStar            bool           `json:"is_all_star"`
	IsPreseason          bool           `json:"is_preseason"`
	Inning               int            `json:"inning"`
	InningStatus         string         `json:"inning_status"`
	Outs                 int            `json:"outs"`
	IsActive             bool           `json:"is_active"`
	RunnerOnFirstBase    bool           `json:"runner_on_first_base"`
	RunnerOnSecondBase   bool           `json:"runner_on_second_base"`
	RunnerOnThirdBase    bool           `json:"runner_on_third_base"`
	Attendance           int            `json:"attendance"`
	Location             string         `json:"location"`
	LocationImageURL     string         `json:"location_image_url"`
	WinningPitcher       WinningPitcher `json:"winning_pitcher"`
	LosingPitcher        LosingPitcher  `json:"losing_pitcher"`
	SavingPitcher        SavingPitcher  `json:"saving_pitcher"`
	CurrentBatter        int            `json:"current_batter"`
	LastPlay             string         `json:"last_play"`
	LocationImageMed     string         `json:"location_image_med"`
	LocationImageSml     string         `json:"location_image_sml"`
	Weather              string         `json:"weather"`
	Wind                 string         `json:"wind"`
	Umpires              []Umpires      `json:"umpires"`
	ExtraInnings         bool           `json:"extra_innings"`
	SeriesMatchupsResult string         `json:"series_matchups_result"`
	SeasonSeries         []SeasonSeries `json:"season_series"`
	CurrentAtBat         CurrentAtBat   `json:"current_at_bat"`
	CipbpComplete        bool           `json:"cipbp_complete"`
	UnderReview          bool           `json:"under_review"`
	GameDuration         string         `json:"game_duration"`
	City                 string         `json:"city"`
	Country              string         `json:"country"`
	Conference           string         `json:"conference"`
	HomePlayoffSeed      int            `json:"home_playoff_seed"`
	VisitingPlayoffSeed  int            `json:"visiting_playoff_seed"`
	State                string         `json:"state"`
	OpeningLine          OpeningLine    `json:"opening_line"`
	CurrentLine          CurrentLine    `json:"current_line"`
}
type Division struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
}
type Conference struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
	ImageURL  string `json:"image_url"`
	Color     string `json:"color"`
}
type GameStats struct {
	Hits     int `json:"hits"`
	Errors   int `json:"errors"`
	HomeRuns int `json:"home_runs"`
	Steals   int `json:"steals"`
	Runs     int `json:"runs"`
}
type HomeRunLeader struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ImageURL  string `json:"image_url"`
	Total     int    `json:"total"`
}
type RunsBattedInLeader struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ImageURL  string `json:"image_url"`
	Total     int    `json:"total"`
}
type Batters struct {
	ID                 int      `json:"id"`
	FirstName          string   `json:"first_name"`
	LastName           string   `json:"last_name"`
	ImageURL           string   `json:"image_url,omitempty"`
	Position           string   `json:"position"`
	ShortPosition      string   `json:"short_position"`
	Number             int      `json:"number"`
	AtBats             int      `json:"at_bats"`
	Runs               int      `json:"runs"`
	HomeRuns           int      `json:"home_runs"`
	Hits               int      `json:"hits"`
	RunsBattedIn       int      `json:"runs_batted_in"`
	Walks              int      `json:"walks"`
	StrikeOuts         int      `json:"strike_outs"`
	LeftOnBase         int      `json:"left_on_base"`
	BattingAverage     string   `json:"batting_average"`
	StolenBases        int      `json:"stolen_bases"`
	Uniform            int      `json:"uniform"`
	Batting            string   `json:"batting"`
	SeasonHomeRuns     int      `json:"season_home_runs"`
	SeasonRunsBattedIn int      `json:"season_runs_batted_in"`
	Doubles            int      `json:"doubles"`
	Triples            int      `json:"triples"`
	OnBasePercentage   string   `json:"on_base_percentage"`
	IsSeasonStat       bool     `json:"is_season_stat"`
	StartedGame        bool     `json:"started_game"`
	SacrificeFlies     int      `json:"sacrifice_flies"`
	HitByPitch         int      `json:"hit_by_pitch"`
	ImageUrls          ImageUrl `json:"image_urls,omitempty"`
	DisplayID          int      `json:"display_id"`
	DisplayName        string   `json:"display_name"`
	PlayerCodeID       int      `json:"player_code_id"`
	BattingSlot        int      `json:"batting_slot"`
	SluggingPercentage string   `json:"slugging_percentage"`
	RunsBattedInSeason int      `json:"runs_batted_in_season"`
}
type Pitchers struct {
	ID                        int      `json:"id"`
	FirstName                 string   `json:"first_name"`
	LastName                  string   `json:"last_name"`
	ImageURL                  string   `json:"image_url,omitempty"`
	Position                  string   `json:"position"`
	ShortPosition             string   `json:"short_position"`
	Number                    int      `json:"number"`
	InningsPitched            string   `json:"innings_pitched"`
	Hits                      int      `json:"hits"`
	Runs                      int      `json:"runs"`
	EarnedRuns                int      `json:"earned_runs"`
	HomeRunsAllowed           int      `json:"home_runs_allowed"`
	PitchCount                int      `json:"pitch_count"`
	Strikes                   int      `json:"strikes"`
	BattersFaced              int      `json:"batters_faced"`
	GroundBalls               int      `json:"ground_balls"`
	FlyBalls                  int      `json:"fly_balls"`
	WildPitches               int      `json:"wild_pitches"`
	EarnedRunAverage          string   `json:"earned_run_average"`
	Walks                     int      `json:"walks"`
	StrikeOuts                int      `json:"strike_outs"`
	Wins                      int      `json:"wins"`
	Losses                    int      `json:"losses"`
	Saves                     int      `json:"saves"`
	Home                      bool     `json:"home"`
	PitcherID                 int      `json:"pitcher_id"`
	PitcherFirstName          string   `json:"pitcher_first_name"`
	PitcherLastName           string   `json:"pitcher_last_name"`
	PitcherNumber             int      `json:"pitcher_number"`
	Uniform                   int      `json:"uniform"`
	Throwing                  string   `json:"throwing"`
	IsSeasonStat              bool     `json:"is_season_stat"`
	Holds                     int      `json:"holds"`
	GameCreditedLoss          bool     `json:"game_credited_loss,omitempty"`
	BlownSaves                int      `json:"blown_saves"`
	ImageUrls                 ImageUrl `json:"image_urls,omitempty"`
	DisplayID                 int      `json:"display_id"`
	DisplayName               string   `json:"display_name"`
	PlayerCodeID              int      `json:"player_code_id"`
	SequenceNumber            int      `json:"sequence_number"`
	OpponentBattingAverage    string   `json:"opponent_batting_average"`
	StrikeOutsSeason          int      `json:"strike_outs_season"`
	WalksHitsPerInningAverage string   `json:"walks_hits_per_inning_average"`
	GameCreditedWin           bool     `json:"game_credited_win,omitempty"`
	GameCreditedSave          bool     `json:"game_credited_save,omitempty"`
}
type WinsLeader struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	ImageURL  string   `json:"image_url"`
	Total     int      `json:"total"`
	ImageUrls ImageUrl `json:"image_urls"`
}
type SavesLeader struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	ImageURL  string   `json:"image_url"`
	Total     int      `json:"total"`
	ImageUrls ImageUrl `json:"image_urls"`
}
type Pitching struct {
	WinningPitcher string   `json:"winning_pitcher"`
	BattersFaced   []string `json:"batters_faced"`
	WildPitches    []string `json:"wild_pitches"`
}
type StrikeoutsLeader struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	ImageURL  string   `json:"image_url"`
	Total     int      `json:"total"`
	ImageUrls ImageUrl `json:"image_urls"`
}
type VisitingTeam struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	ShortName          string             `json:"short_name"`
	City               string             `json:"city"`
	ImageURL           string             `json:"image_url"`
	Score              int                `json:"score"`
	ImageURL90         string             `json:"image_url_90"`
	Division           Division           `json:"division"`
	Conference         Conference         `json:"conference"`
	GameStats          GameStats          `json:"game_stats"`
	HomeRunLeader      HomeRunLeader      `json:"home_run_leader"`
	RunsBattedInLeader RunsBattedInLeader `json:"runs_batted_in_leader"`
	Batters            []Batters          `json:"batters"`
	Pitchers           []Pitchers         `json:"pitchers"`
	Color              string             `json:"color"`
	CurrentPitcher     int                `json:"current_pitcher"`
	AtBat              bool               `json:"at_bat"`
	WinsLeader         WinsLeader         `json:"wins_leader"`
	SavesLeader        SavesLeader        `json:"saves_leader"`
	Pitching           Pitching           `json:"pitching"`
	ImageURL59         string             `json:"image_url_59"`
	StrikeoutsLeader   StrikeoutsLeader   `json:"strikeouts_leader"`
	ImageURL25         string             `json:"image_url_25"`
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

type HomeTeam struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	ShortName          string             `json:"short_name"`
	City               string             `json:"city"`
	ImageURL           string             `json:"image_url"`
	Score              int                `json:"score"`
	ImageURL90         string             `json:"image_url_90"`
	Division           Division           `json:"division"`
	Conference         Conference         `json:"conference"`
	Injuries           []Injuries         `json:"injuries"`
	GameStats          GameStats          `json:"game_stats"`
	HomeRunLeader      HomeRunLeader      `json:"home_run_leader"`
	RunsBattedInLeader RunsBattedInLeader `json:"runs_batted_in_leader"`
	Batters            []Batters          `json:"batters"`
	Pitchers           []Pitchers         `json:"pitchers"`
	Color              string             `json:"color"`
	CurrentPitcher     int                `json:"current_pitcher"`
	AtBat              bool               `json:"at_bat"`
	WinsLeader         WinsLeader         `json:"wins_leader"`
	SavesLeader        SavesLeader        `json:"saves_leader"`
	Pitching           Pitching           `json:"pitching"`
	ImageURL59         string             `json:"image_url_59"`
	StrikeoutsLeader   StrikeoutsLeader   `json:"strikeouts_leader"`
	ImageURL25         string             `json:"image_url_25"`
}
