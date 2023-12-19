package nhl

import "github.com/digitalmedia/sn_stats_models/common"

type GameDataWeb struct {
	Sequence     int64      `json:"sequence"`
	Details      DetailsWeb `json:"details"`
	VisitingTeam TeamWeb    `json:"visiting_team"`
	HomeTeam     TeamWeb    `json:"home_team"`
	Periods      []Periods  `json:"periods"`
}

type ImageUrl struct {
	Lg string `json:"lg"`
	Md string `json:"md"`
	Sm string `json:"sm"`
	Xs string `json:"xs"`
}
type Stars struct {
	ID               string   `json:"id"`
	TeamID           string   `json:"team_id"`
	FirstName        string   `json:"first_name"`
	LastName         string   `json:"last_name"`
	ImageURL         string   `json:"image_url"`
	Position         string   `json:"position"`
	ShortPosition    string   `json:"short_position"`
	Number           int      `json:"number"`
	Status           bool     `json:"status"`
	Goals            int      `json:"goals"`
	Assists          int      `json:"assists"`
	Points           int      `json:"points"`
	PlusMinusRatio   int      `json:"plus_minus_ratio"`
	PenaltyInMinutes int      `json:"penalty_in_minutes"`
	Minutes          int      `json:"minutes"`
	Seconds          int      `json:"seconds"`
	ShotsOnGoal      int      `json:"shots_on_goal"`
	TimeOnIce        string   `json:"time_on_ice"`
	Hits             int      `json:"hits"`
	Faceoffs         int      `json:"faceoffs"`
	ImageUrls        ImageUrl `json:"image_urls"`
}
type StartingGoalie struct {
	ID                      string   `json:"id"`
	SnPlayerId              int      `json:"sn_player_id"`
	FirstName               string   `json:"first_name"`
	LastName                string   `json:"last_name"`
	GameID                  int      `json:"game_id"`
	GameDate                string   `json:"game_date"`
	TeamID                  string   `json:"team_id"`
	TeamShortName           string   `json:"team_short_name"`
	Record                  string   `json:"record"`
	Gaa                     string   `json:"gaa"`
	Svpct                   string   `json:"svpct"`
	So                      int      `json:"so"`
	Status                  int      `json:"status"`
	StatusAttributionURL    string   `json:"status_attribution_url"`
	StatusAttributionSource string   `json:"status_attribution_source"`
	ImageURL                string   `json:"image_url"`
	Otl                     int      `json:"otl"`
	Number                  int      `json:"number"`
	ImageUrls               ImageUrl `json:"image_urls"`
}
type Story struct {
	Type     string `json:"type"`
	Headline string `json:"headline"`
	Content  string `json:"content"`
}

type WinningGoalies struct {
	ID             int      `json:"id"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	ImageURL       string   `json:"image_url"`
	Position       string   `json:"position"`
	ShortPosition  string   `json:"short_position"`
	Number         int      `json:"number"`
	Status         bool     `json:"status"`
	ShotsAgainst   int      `json:"shots_against"`
	GoalsAgainst   int      `json:"goals_against"`
	SavePercentage float64  `json:"save_percentage"`
	Minutes        int      `json:"minutes"`
	Seconds        int      `json:"seconds"`
	Saves          int      `json:"saves"`
	Decision       string   `json:"decision"`
	Record         string   `json:"record"`
	ImageUrls      ImageUrl `json:"image_urls"`
}
type LosingGoalies struct {
	ID             int      `json:"id"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	ImageURL       string   `json:"image_url"`
	Position       string   `json:"position"`
	ShortPosition  string   `json:"short_position"`
	Number         int      `json:"number"`
	Status         bool     `json:"status"`
	ShotsAgainst   int      `json:"shots_against"`
	GoalsAgainst   int      `json:"goals_against"`
	SavePercentage float64  `json:"save_percentage"`
	Minutes        int      `json:"minutes"`
	Seconds        int      `json:"seconds"`
	Saves          int      `json:"saves"`
	Decision       string   `json:"decision"`
	Record         string   `json:"record"`
	ImageUrls      ImageUrl `json:"image_urls"`
}
type Last9Events struct {
	PlayerID          string   `json:"player_id"`
	PlayerFirstName   string   `json:"player_first_name"`
	PlayerLastName    string   `json:"player_last_name"`
	Description       string   `json:"description"`
	Period            int      `json:"period"`
	Time              string   `json:"time"`
	TeamID            string   `json:"team_id"`
	PlayerDisplayName string   `json:"player_display_name"`
	EventID           string   `json:"event_id"`
	Event             string   `json:"event"`
	DetailID          int      `json:"detail_id"`
	Detail            string   `json:"detail"`
	SpecialShot       string   `json:"special_shot"`
	Strength          string   `json:"strength"`
	StrengthType      string   `json:"strength_type"`
	AwayScoreBefore   int      `json:"away_score_before"`
	AwayScoreAfter    int      `json:"away_score_after"`
	HomeScoreBefore   int      `json:"home_score_before"`
	HomeScoreAfter    int      `json:"home_score_after"`
	ImageURL          string   `json:"image_url"`
	ID                string   `json:"id"`
	ImageUrls         ImageUrl `json:"image_urls"`
	Sequence          int64    `json:"sequence"`
}
type GameWinningGoal struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
type DetailsWeb struct {
	LeagueShortName     string                 `json:"league_short_name"`
	ID                  int                    `json:"id"`
	SrGameUuid          string                 `json:"sr_game_uuid"`
	Period              int                    `json:"period"`
	IsActive            bool                   `json:"is_active"`
	Clock               string                 `json:"clock"`
	Type                string                 `json:"type"`
	Timestamp           int64                  `json:"timestamp"`
	Status              string                 `json:"status"`
	Location            string                 `json:"location"`
	LocationImageURL    string                 `json:"location_image_url"`
	Broadcast           []common.GameBroadcast `json:"broadcast"`
	Overtime            int                    `json:"overtime"`
	HasShootout         bool                   `json:"has_shootout"`
	HomeSeriesWins      int                    `json:"home_series_wins"`
	VisitingSeriesWins  int                    `json:"visiting_series_wins"`
	IsAllStar           bool                   `json:"is_all_star"`
	Stars               []Skaters              `json:"stars"`
	LocationImageMed    string                 `json:"location_image_med"`
	LocationImageSml    string                 `json:"location_image_sml"`
	StartingGoalies     []StartingGoalie       `json:"starting_goalies"`
	Story               Story                  `json:"story"`
	Attendance          int                    `json:"attendance"`
	Arena               string                 `json:"arena"`
	Referee1            string                 `json:"referee1"`
	Referee2            string                 `json:"referee2"`
	Linesman1           string                 `json:"linesman1"`
	Linesman2           string                 `json:"linesman2"`
	Last5Meetings       []RecentGames          `json:"last_5_meetings"`
	WinningGoalies      []WinningGoalies       `json:"winning_goalies"`
	LosingGoalies       []LosingGoalies        `json:"losing_goalies"`
	PpMin               int                    `json:"pp_min"`
	PpSec               int                    `json:"pp_sec"`
	Last9Events         []Last9Events          `json:"last_9_events"`
	GameWinningGoal     GameWinningGoal        `json:"game_winning_goal"`
	SeasonSeries        []RecentGames          `json:"season_series"`
	City                string                 `json:"city"`
	Country             string                 `json:"country"`
	GameNumber          int                    `json:"game_number"`
	Round               int                    `json:"round"`
	Conference          string                 `json:"conference"`
	IfNecessary         bool                   `json:"if_necessary"`
	HomePlayoffSeed     int                    `json:"home_playoff_seed"`
	VisitingPlayoffSeed int                    `json:"visiting_playoff_seed"`
	NhlID               int                    `json:"nhl_id"`
	State               string                 `json:"state"`
	OpeningLine         *OpeningLine           `json:"opening_line"`
	CurrentLine         *CurrentLine           `json:"current_line"`
	Tba                 bool                   `json:"tba"`
}

type RecentGames struct {
	Date              string `json:"date"`
	Location          string `json:"location"`
	GameID            string `json:"game_id"`
	HomeTeam          string `json:"home_team"`
	VisitingTeam      string `json:"visiting_team"`
	VisitingTeamScore int    `json:"visiting_team_score"`
	HomeTeamScore     int    `json:"home_team_score"`
}

type Division struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
}
type Conference struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
	ImageURL  string `json:"image_url"`
	Color     string `json:"color"`
}

type SeasonStats struct {
	Wins                int     `json:"wins"`
	Losses              int     `json:"losses"`
	OvertimeLosses      int     `json:"overtime_losses"`
	Points              int     `json:"points"`
	GoalsPerGame        float64 `json:"goals_per_game"`
	GoalsAgainstPerGame float64 `json:"goals_against_per_game"`
	TeamPoints          int     `json:"team_points"`
	PowerPlayPercentage float64 `json:"power_play_percentage"`
	PenaltyKill         float64 `json:"penalty_kill"`
}
type GameStats struct {
	Shots                  int    `json:"shots"`
	Hits                   int    `json:"hits"`
	Faceoffs               int    `json:"faceoffs"`
	PenaltiesInMinutes     int    `json:"penalties_in_minutes"`
	Blocks                 int    `json:"blocks"`
	TurnoverGiveaways      int    `json:"turnover_giveaways"`
	PowerplayGoals         int    `json:"powerplay_goals"`
	PowerplayOpportunities int    `json:"powerplay_opportunities"`
	TurnoverTakeaways      int    `json:"turnover_takeaways"`
	Powerplay              string `json:"powerplay"`
}
type Leader struct {
	ID        string   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	ImageURL  string   `json:"image_url"`
	Total     int      `json:"total"`
	ImageUrls ImageUrl `json:"image_urls"`
}
type Skaters struct {
	ID                string   `json:"id"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	ImageURL          string   `json:"image_url"`
	Position          string   `json:"position"`
	ShortPosition     string   `json:"short_position"`
	Number            int      `json:"number"`
	Status            bool     `json:"status"`
	Goals             int      `json:"goals"`
	Assists           int      `json:"assists"`
	Points            int      `json:"points"`
	PlusMinusRatio    int      `json:"plus_minus_ratio"`
	PenaltyInMinutes  int      `json:"penalty_in_minutes"`
	Minutes           int      `json:"minutes"`
	Seconds           int      `json:"seconds"`
	ShotsOnGoal       int      `json:"shots_on_goal"`
	TimeOnIce         string   `json:"time_on_ice"`
	Hits              int      `json:"hits"`
	Faceoffs          int      `json:"faceoffs"`
	FaceoffLost       int      `json:"faceoff_lost"`
	Blocks            int      `json:"blocks"`
	TurnoverGiveaways int      `json:"turnover_giveaways"`
	TurnoverTakeaways int      `json:"turnover_takeaways"`
	ImageUrls         ImageUrl `json:"image_urls"`
}
type Goalies struct {
	ID             string   `json:"id"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	ImageURL       string   `json:"image_url"`
	Position       string   `json:"position"`
	ShortPosition  string   `json:"short_position"`
	Number         int      `json:"number"`
	Status         bool     `json:"status"`
	ShotsAgainst   int      `json:"shots_against"`
	GoalsAgainst   int      `json:"goals_against"`
	SavePercentage float64  `json:"save_percentage"`
	Minutes        int      `json:"minutes"`
	Seconds        int      `json:"seconds"`
	Saves          int      `json:"saves"`
	ImageUrls      ImageUrl `json:"image_urls"`
}

type Lineup struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}
type TeamStandings struct {
	Wins                   int    `json:"wins"`
	Losses                 int    `json:"losses"`
	GamesPlayed            int    `json:"games_played"`
	OvertimeLosses         int    `json:"overtime_losses"`
	Points                 int    `json:"points"`
	WinPercentage          string `json:"win_percentage"`
	OvertimeWins           int    `json:"overtime_wins"`
	ShootoutLosses         int    `json:"shootout_losses"`
	Row                    int    `json:"row"`
	GoalsFor               int    `json:"goals_for"`
	GoalsAgainst           int    `json:"goals_against"`
	Last10Wins             int    `json:"last_10_wins"`
	Last10Losses           int    `json:"last_10_losses"`
	Last10OtLosses         int    `json:"last_10_ot_losses"`
	HomeWins               int    `json:"home_wins"`
	HomeLosses             int    `json:"home_losses"`
	HomeOtLosses           int    `json:"home_ot_losses"`
	VisitingWins           int    `json:"visiting_wins"`
	VisitingLosses         int    `json:"visiting_losses"`
	VisitingOtLosses       int    `json:"visiting_ot_losses"`
	Streak                 string `json:"streak"`
	ShootoutWins           int    `json:"shootout_wins"`
	Last10OtWins           int    `json:"last_10_ot_wins"`
	RegulationWins         int    `json:"regulation_wins"`
	Last10ShootoutWins     int    `json:"last_10_shootout_wins"`
	Last10ShootoutLosses   int    `json:"last_10_shootout_losses"`
	HomeOtWins             int    `json:"home_ot_wins"`
	HomeShootoutWins       int    `json:"home_shootout_wins"`
	HomeShootoutLosses     int    `json:"home_shootout_losses"`
	VisitingOtWins         int    `json:"visiting_ot_wins"`
	VisitingShootoutWins   int    `json:"visiting_shootout_wins"`
	VisitingShootoutLosses int    `json:"visiting_shootout_losses"`
}
type TeamWeb struct {
	ID            string        `json:"id"`
	SrTeamUUID    string        `json:"sr_team_uuid"`
	Name          string        `json:"name"`
	ShortName     string        `json:"short_name"`
	City          string        `json:"city"`
	ImageURL      string        `json:"image_url"`
	Score         int           `json:"score"`
	ImageURL90    string        `json:"image_url_90"`
	StrengthType  string        `json:"strength_type"`
	Division      Division      `json:"division"`
	Conference    Conference    `json:"conference"`
	Injuries      []Injuries    `json:"injuries"`
	SeasonStats   SeasonStats   `json:"season_stats"`
	GameStats     GameStats     `json:"game_stats"`
	GoalLeader    Leader        `json:"goal_leader"`
	AssistLeader  Leader        `json:"assist_leader"`
	WinLeader     Leader        `json:"win_leader"`
	Skaters       []Skaters     `json:"skaters"`
	Goalies       []Goalies     `json:"goalies"`
	Color         string        `json:"color"`
	PointLeader   Leader        `json:"point_leader"`
	Strength      int           `json:"strength"`
	Lineup        []Lineup      `json:"lineup"`
	ImageURL25    string        `json:"image_url_25"`
	ImageURL59    string        `json:"image_url_59"`
	TeamStandings TeamStandings `json:"team_standings"`
}
type Injuries struct {
	PlayerID             int    `json:"player_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Position             string `json:"position"`
	Status               string `json:"status"`
	Timestamp            int    `json:"timestamp"`
	Type                 string `json:"type"`
	ShortPosition        string `json:"short_position,omitempty"`
	Number               int    `json:"number,omitempty"`
	DisabilityListStatus string `json:"disability_list_status,omitempty"`
}
type HomeTeam struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	ShortName     string        `json:"short_name"`
	City          string        `json:"city"`
	ImageURL      string        `json:"image_url"`
	Score         int           `json:"score"`
	ImageURL90    string        `json:"image_url_90"`
	StrengthType  string        `json:"strength_type"`
	Division      Division      `json:"division"`
	Conference    Conference    `json:"conference"`
	Injuries      []Injuries    `json:"injuries"`
	SeasonStats   SeasonStats   `json:"season_stats"`
	GameStats     GameStats     `json:"game_stats"`
	GoalLeader    Leader        `json:"goal_leader"`
	AssistLeader  Leader        `json:"assist_leader"`
	WinLeader     Leader        `json:"win_leader"`
	Skaters       []Skaters     `json:"skaters"`
	Goalies       []Goalies     `json:"goalies"`
	Color         string        `json:"color"`
	PointLeader   Leader        `json:"point_leader"`
	Strength      int           `json:"strength"`
	Lineup        []Lineup      `json:"lineup"`
	ImageURL25    string        `json:"image_url_25"`
	ImageURL59    string        `json:"image_url_59"`
	TeamStandings TeamStandings `json:"team_standings"`
}
type AssistingPlayers struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ImageURL     string `json:"image_url"`
	AssistNumber int    `json:"assist_number"`
}
type PlayerImgUrls struct {
	Lg string `json:"lg"`
	Md string `json:"md"`
	Sm string `json:"sm"`
	Xs string `json:"xs"`
}
type Goals struct {
	Minutes          int                `json:"minutes"`
	Seconds          int                `json:"seconds"`
	TeamStrength     string             `json:"team_strength"`
	TeamID           int                `json:"team_id"`
	PlayerID         int                `json:"player_id"`
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	GoalNumber       int                `json:"goal_number"`
	AssistingPlayers []AssistingPlayers `json:"assisting_players"`
	TeamImgURL       string             `json:"team_img_url"`
	PlayerImgURL     string             `json:"player_img_url"`
	GoalStrength     string             `json:"goal_strength"`
	EmptyNet         string             `json:"empty_net"`
	PenaltyShot      string             `json:"penalty_shot"`
	PlayerImgUrls    PlayerImgUrls      `json:"player_img_urls"`
}
type Penalties struct {
	ID            int    `json:"id"`
	Description   string `json:"description"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PlayerID      int    `json:"player_id"`
	UniformNumber int    `json:"uniform_number"`
	TeamName      string `json:"team_name"`
	TeamAlias     string `json:"team_alias"`
	TeamCity      string `json:"team_city"`
	TeamCode      int    `json:"team_code"`
	Time          string `json:"time"`
	Length        int    `json:"length"`
	ServedBy      string `json:"served_by"`
	TeamImgURL    string `json:"team_img_url"`
}
type Periods struct {
	VisitingTeamScore int         `json:"visiting_team_score"`
	HomeTeamScore     int         `json:"home_team_score"`
	Goals             []Goals     `json:"goals"`
	Penalties         []Penalties `json:"penalties"`
}
