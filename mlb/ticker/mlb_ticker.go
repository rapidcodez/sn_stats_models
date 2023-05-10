package sn_stats_models

type MlbTicker struct {
	Status string `json:"status,omitempty"`
	Code   int    `json:"code"`
	Data   Data   `json:"data,omitempty"`
}
type ImageUrls struct {
	Lg string `json:"lg,omitempty"`
	Md string `json:"md,omitempty"`
	Sm string `json:"sm,omitempty"`
	Xs string `json:"xs,omitempty"`
}
type StartingPitcher struct {
	ID               int        `json:"id,omitempty"`
	FirstName        string     `json:"first_name,omitempty"`
	LastName         string     `json:"last_name,omitempty"`
	ImageURL         string     `json:"image_url,omitempty"`
	Position         string     `json:"position,omitempty"`
	ShortPosition    string     `json:"short_position,omitempty"`
	Number           int        `json:"number,omitempty"`
	InningsPitched   string     `json:"innings_pitched,omitempty"`
	HomeRunsAllowed  int        `json:"home_runs_allowed,omitempty"`
	EarnedRunAverage string     `json:"earned_run_average,omitempty"`
	StrikeOuts       int        `json:"strike_outs,omitempty"`
	Wins             int        `json:"wins,omitempty"`
	Losses           int        `json:"losses,omitempty"`
	Throwing         string     `json:"throwing,omitempty"`
	GamesPlayed      int        `json:"games_played,omitempty"`
	Whip             float64    `json:"whip,omitempty"`
	ImageUrls        *ImageUrls `json:"image_urls,omitempty"`
}
type VisitingTeam struct {
	ID              int              `json:"id,omitempty"`
	Score           int              `json:"score,omitempty"`
	ImgURL90        string           `json:"img_url_90,omitempty"`
	ImgURL25        string           `json:"img_url_25,omitempty"`
	ImgURL59        string           `json:"img_url_59,omitempty"`
	ShortName       string           `json:"short_name,omitempty"`
	City            string           `json:"city,omitempty"`
	Name            string           `json:"name,omitempty"`
	StartingPitcher *StartingPitcher `json:"starting_pitcher,omitempty"`
	Wins            int              `json:"wins,omitempty"`
	Losses          int              `json:"losses,omitempty"`
	GamesBehind     float64          `json:"games_behind,omitempty"`
}
type HomeTeam struct {
	ID              int              `json:"id,omitempty"`
	Score           int              `json:"score,omitempty"`
	ImgURL90        string           `json:"img_url_90,omitempty"`
	ImgURL25        string           `json:"img_url_25,omitempty"`
	ImgURL59        string           `json:"img_url_59,omitempty"`
	ShortName       string           `json:"short_name,omitempty"`
	City            string           `json:"city,omitempty"`
	Name            string           `json:"name,omitempty"`
	StartingPitcher *StartingPitcher `json:"starting_pitcher,omitempty"`
	Wins            int              `json:"wins,omitempty"`
	Losses          int              `json:"losses,omitempty"`
	GamesBehind     int              `json:"games_behind,omitempty"`
}
type TickerGame struct {
	GameStatus           string         `json:"game_status"`
	Timestamp            int            `json:"timestamp"`
	ID                   int            `json:"id"`
	Code                 int            `json:"code"`
	Period               int            `json:"period"`
	Clock                string         `json:"clock"`
	Type                 string         `json:"type"`
	League               string         `json:"league"`
	Sport                string         `json:"sport"`
	Active               bool           `json:"active"`
	Tba                  bool           `json:"tba"`
	Env                  string         `json:"env"`
	SeriesMatchupsResult string         `json:"series_matchups_result"`
	Location             string         `json:"location"`
	City                 string         `json:"city"`
	Country              string         `json:"country"`
	LocationImageURL     string         `json:"location_image_url"`
	VisitingTeam         *VisitingTeam  `json:"visiting_team"`
	HomeTeam             *HomeTeam      `json:"home_team"`
	Broadcast            *[]interface{} `json:"broadcast"`
	OpeningLine          interface{}    `json:"opening_line"`
	CurrentLine          interface{}    `json:"current_line"`
	PeriodStatus         string         `json:"period_status"`
	Outs                 int            `json:"outs"`
	RunnerOnFirstBase    bool           `json:"runner_on_first_base"`
	RunnerOnSecondBase   bool           `json:"runner_on_second_base"`
	RunnerOnThirdBase    bool           `json:"runner_on_third_base"`
}
type Settings struct {
	Sport             string `json:"sport,omitempty"`
	League            string `json:"league,omitempty"`
	Time              int    `json:"time,omitempty"`
	TurnoverType      string `json:"turnover_type,omitempty"`
	TurnoverHour      int    `json:"turnover_hour,omitempty"`
	TurnoverWeeklyDay string `json:"turnover_weekly_day,omitempty"`
	TickerVisible     bool   `json:"ticker_visible,omitempty"`
	GlobalTicker      bool   `json:"global_ticker,omitempty"`
	SportTicker       bool   `json:"sport_ticker,omitempty"`
}
type Data struct {
	Service        bool         `json:"service,omitempty"`
	ServiceMessage string       `json:"service_message,omitempty"`
	Ts             int          `json:"ts,omitempty"`
	Games          []TickerGame `json:"games,omitempty"`
	Settings       *[]Settings  `json:"settings,omitempty"`
}
