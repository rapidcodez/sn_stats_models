package wnba

import (
	"time"

	"github.com/digitalmedia/sn_stats_models/common"
)

type GameDataWeb struct {
	Details      DetailsWeb `json:"details"`
	VisitingTeam TeamWeb    `json:"visiting_team"`
	HomeTeam     TeamWeb    `json:"home_team"`
	Quarters     []Quarters `json:"quarters"`
	Sequence     int64      `json:"sequence,omitempty"`
}
type Story struct {
	Type     string `json:"type"`
	Headline string `json:"headline"`
	Content  string `json:"content"`
}
type Last5Meetings struct {
	Date              string `json:"date"`
	Location          string `json:"location"`
	GameID            string `json:"game_id"`
	SRGameID          string `json:"sr_game_id"`
	HomeTeam          string `json:"home_team"`
	VisitingTeam      string `json:"visiting_team"`
	VisitingTeamScore int    `json:"visiting_team_score"`
	HomeTeamScore     int    `json:"home_team_score"`
}
type ImageUrls struct {
	Lg string `json:"lg"`
	Md string `json:"md"`
	Sm string `json:"sm"`
	Xs string `json:"xs"`
}
type Last9Events struct {
	PlayerID          string    `json:"player_id"`
	PlayerFirstName   string    `json:"player_first_name"`
	PlayerLastName    string    `json:"player_last_name"`
	Description       string    `json:"description"`
	Quarter           int       `json:"quarter"`
	Time              string    `json:"time"`
	TeamID            string    `json:"team_id"`
	PlayerDisplayName string    `json:"player_display_name"`
	EventID           string    `json:"event_id"`
	Event             string    `json:"event"`
	DetailID          int       `json:"detail_id"`
	Detail            string    `json:"detail"`
	ImageURL          string    `json:"image_url"`
	ID                int       `json:"id"`
	ImageUrls         ImageUrls `json:"image_urls"`
	Position          string    `json:"position"`
	Sequence          int64     `json:"sequence"`
}
type PlayersOnCourt struct {
	ID                     string    `json:"id"`
	FirstName              string    `json:"first_name"`
	LastName               string    `json:"last_name"`
	ImageURL               string    `json:"image_url"`
	Position               string    `json:"position"`
	ShortPosition          string    `json:"short_position"`
	Number                 int       `json:"number"`
	Mins                   string    `json:"mins"`
	Points                 int       `json:"points"`
	ReboundsTotal          int       `json:"rebounds_total"`
	Assists                int       `json:"assists"`
	Steals                 int       `json:"steals"`
	BlockedShots           int       `json:"blocked_shots"`
	OffensiveRebounds      int       `json:"offensive_rebounds"`
	DefensiveRebounds      int       `json:"defensive_rebounds"`
	FieldGoalsMade         int       `json:"field_goals_made"`
	FieldGoalsAttempted    int       `json:"field_goals_attempted"`
	ThreePointersMade      int       `json:"three_pointers_made"`
	ThreePointersAttempted int       `json:"three_pointers_attempted"`
	FreeThrowsMade         int       `json:"free_throws_made"`
	FreeThrowsAttempted    int       `json:"free_throws_attempted"`
	Turnovers              int       `json:"turnovers"`
	PersonalFouls          int       `json:"personal_fouls"`
	Games                  int       `json:"games"`
	Reason                 string    `json:"reason"`
	TeamID                 string    `json:"team_id"`
	ImageUrls              ImageUrls `json:"image_urls"`
}
type PlayersList struct {
	ID                     string    `json:"id"`
	FirstName              string    `json:"first_name"`
	LastName               string    `json:"last_name"`
	ImageURL               string    `json:"image_url"`
	Position               string    `json:"position"`
	ShortPosition          string    `json:"short_position"`
	Number                 int       `json:"number"`
	Mins                   string    `json:"mins"`
	Points                 int       `json:"points"`
	ReboundsTotal          int       `json:"rebounds_total"`
	Assists                int       `json:"assists"`
	Steals                 int       `json:"steals"`
	BlockedShots           int       `json:"blocked_shots"`
	OffensiveRebounds      int       `json:"offensive_rebounds"`
	DefensiveRebounds      int       `json:"defensive_rebounds"`
	FieldGoalsMade         int       `json:"field_goals_made"`
	FieldGoalsAttempted    int       `json:"field_goals_attempted"`
	ThreePointersMade      int       `json:"three_pointers_made"`
	ThreePointersAttempted int       `json:"three_pointers_attempted"`
	FreeThrowsMade         int       `json:"free_throws_made"`
	FreeThrowsAttempted    int       `json:"free_throws_attempted"`
	Turnovers              int       `json:"turnovers"`
	PersonalFouls          int       `json:"personal_fouls"`
	Games                  int       `json:"games"`
	Reason                 string    `json:"reason"`
	TeamID                 string    `json:"team_id"`
	SecondaryPosition      string    `json:"secondary_position"`
	ImageUrls              ImageUrls `json:"image_urls"`
}
type OpeningLine struct {
	FavID         int     `json:"fav_id"`
	Name          string  `json:"name"`
	FavPoints     float64 `json:"fav_points"`
	FavMoney      int     `json:"fav_money"`
	UnderdogMoney int     `json:"underdog_money"`
	HomeMoney     int     `json:"home_money"`
	AwayMoney     int     `json:"away_money"`
	Total         float64 `json:"total"`
	OverMoney     int     `json:"over_money"`
	UnderMoney    int     `json:"under_money"`
	DrawMoney     int     `json:"draw_money"`
}
type CurrentLine struct {
	FavID         int    `json:"fav_id"`
	Name          string `json:"name"`
	FavPoints     int    `json:"fav_points"`
	FavMoney      int    `json:"fav_money"`
	UnderdogMoney int    `json:"underdog_money"`
	HomeMoney     int    `json:"home_money"`
	AwayMoney     int    `json:"away_money"`
	Total         int    `json:"total"`
	OverMoney     int    `json:"over_money"`
	UnderMoney    int    `json:"under_money"`
	DrawMoney     int    `json:"draw_money"`
}
type DetailsWeb struct {
	LeagueShortName      string                 `json:"league_short_name"`
	ID                   string                 `json:"id"`
	SrGameId             string                 `json:"sr_game_id"`
	SrGameUuid           string                 `json:"sr_game_uuid"`
	Quarter              int                    `json:"quarter"`
	IsActive             bool                   `json:"is_active"`
	Clock                string                 `json:"clock"`
	Type                 string                 `json:"type"`
	Timestamp            int                    `json:"timestamp"`
	Datetime             time.Time              `json:"datetime"`
	Status               string                 `json:"status"`
	SrStatus             string                 `json:"sr_status"`
	Location             string                 `json:"location"`
	LocationImageURL     string                 `json:"location_image_url"`
	Broadcast            []common.GameBroadcast `json:"broadcast"`
	Overtime             int                    `json:"overtime"`
	HomeSeriesWins       int                    `json:"home_series_wins"`
	VisitingSeriesWins   int                    `json:"visiting_series_wins"`
	IsAllStar            bool                   `json:"is_all_star"`
	SeriesRound          int                    `json:"series_round"`
	LocationImageMed     string                 `json:"location_image_med"`
	LocationImageSml     string                 `json:"location_image_sml"`
	Story                Story                  `json:"story"`
	Last5Meetings        []Last5Meetings        `json:"last_5_meetings"`
	SeriesMatchupsResult string                 `json:"series_matchups_result"`
	Last9Events          []Last9Events          `json:"last_9_events"`
	PlayersOnCourt       []PlayersOnCourt       `json:"players_on_court"`
	PlayersList          []PlayersList          `json:"players_list"`
	Timeout              string                 `json:"timeout"`
	City                 string                 `json:"city"`
	Country              string                 `json:"country"`
	IfNecessary          bool                   `json:"if_necessary"`
	HomePlayoffSeed      int                    `json:"home_playoff_seed"`
	VisitingPlayoffSeed  int                    `json:"visiting_playoff_seed"`
	State                string                 `json:"state"`
	OpeningLine          OpeningLine            `json:"opening_line"`
	CurrentLine          CurrentLine            `json:"current_line"`
	SimMode              bool                   `json:"sim_mode"`
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
type Injuries struct {
	PlayerID  string `json:"player_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
	Status    string `json:"status"`
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"`
}
type SeasonStats struct {
	Wins                      int     `json:"wins"`
	Losses                    int     `json:"losses"`
	PointsPerGame             float64 `json:"points_per_game"`
	OffenseReboundsPerGame    float64 `json:"offense_rebounds_per_game"`
	DefenseReboundsPerGame    float64 `json:"defense_rebounds_per_game"`
	PointsAllowedPerGame      float64 `json:"points_allowed_per_game"`
	FieldGoalPercent          float64 `json:"field_goal_percent"`
	ThreePointShootingPercent float64 `json:"three_point_shooting_percent"`
	AssistsPerGame            float64 `json:"assists_per_game"`
	StealsPerGame             float64 `json:"steals_per_game"`
	BlocksPerGame             float64 `json:"blocks_per_game"`
	ThreePointFgPct           float64 `json:"three_point_fg_pct"`
	FreeThrowPct              float64 `json:"free_throw_pct"`
	TurnoversPerGame          float64 `json:"turnovers_per_game"`
	ReboundsPerGame           float64 `json:"rebounds_per_game"`
	GamesBack                 float64 `json:"games_back"`
}
type GameStats struct {
	FieldGoalsMade           int    `json:"field_goals_made"`
	FieldGoalsMissed         int    `json:"field_goals_missed"`
	FieldGoalsAttempted      int    `json:"field_goals_attempted"`
	FieldGoalsPct            string `json:"field_goals_pct"`
	Assists                  int    `json:"assists"`
	ReboundsTotal            int    `json:"rebounds_total"`
	ReboundsOffensive        int    `json:"rebounds_offensive"`
	ReboundsDefensive        int    `json:"rebounds_defensive"`
	PersonalFouls            int    `json:"personal_fouls"`
	TotalFouls               int    `json:"total_fouls"`
	FreeThrowsMade           int    `json:"free_throws_made"`
	FreeThrowsMissed         int    `json:"free_throws_missed"`
	FreeThrowsAttempted      int    `json:"free_throws_attempted"`
	FreeThrowsPct            string `json:"free_throws_pct"`
	ThreePointPct            string `json:"three_point_pct"`
	Turnovers                int    `json:"turnovers"`
	Blocks                   int    `json:"blocks"`
	Steals                   int    `json:"steals"`
	ThreePointGoalsMade      int    `json:"three_point_goals_made"`
	ThreePointGoalsAttempted int    `json:"three_point_goals_attempted"`
	ThreePointGoalsPct       string `json:"three_point_goals_pct"`
	Minutes                  int    `json:"minutes"`
}

type GameLeaderStat struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     int       `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type GameLeaders struct {
	PointsLeader       GameLeaderStat `json:"points_leader"`
	ReboundLeader      GameLeaderStat `json:"rebound_leader"`
	AssistLeader       GameLeaderStat `json:"assist_leader"`
	BlockedShotsLeader GameLeaderStat `json:"blocked_shots_leader"`
	StealsLeader       GameLeaderStat `json:"steals_leader"`
	TurnoversLeader    GameLeaderStat `json:"turnovers_leader"`
}
type PointsPerGameLeader struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     float64   `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type AssistsPerGameLeader struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     float64   `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type ReboundsPerGameLeader struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     float64   `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type BlocksPerGameLeader struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     float64   `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type StealsPerGameLeader struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     float64   `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type FieldGoalPercentageLeader struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     float64   `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type PerGameLeader struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ImageURL  string    `json:"image_url"`
	Total     float64   `json:"total"`
	ImageUrls ImageUrls `json:"image_urls"`
}
type SeasonLeaders struct {
	AssistsPerGameLeader      PerGameLeader `json:"assists_per_game_leader"`
	BlocksPerGameLeader       PerGameLeader `json:"blocks_per_game_leader"`
	FieldGoalPercentageLeader PerGameLeader `json:"field_goal_percentage_leader"`
	FreeThrowPCTLeader        PerGameLeader `json:"free_throw_percentage_leader"`
	PointsPerGameLeader       PerGameLeader `json:"points_per_game_leader"`
	ReboundsPerGameLeader     PerGameLeader `json:"rebounds_per_game_leader"`
	StealsPerGameLeader       PerGameLeader `json:"steals_per_game_leader"`
	ThreePointPCTLeader       PerGameLeader `json:"three_point_percentage_leader"`
}
type PlayerStats struct {
	ID                     string    `json:"id"`
	FirstName              string    `json:"first_name"`
	LastName               string    `json:"last_name"`
	ImageURL               string    `json:"image_url"`
	Position               string    `json:"position"`
	ShortPosition          string    `json:"short_position"`
	Number                 int       `json:"number"`
	Mins                   string    `json:"mins"`
	Points                 int       `json:"points"`
	ReboundsTotal          int       `json:"rebounds_total"`
	Assists                int       `json:"assists"`
	Steals                 int       `json:"steals"`
	BlockedShots           int       `json:"blocked_shots"`
	OffensiveRebounds      int       `json:"offensive_rebounds"`
	DefensiveRebounds      int       `json:"defensive_rebounds"`
	FieldGoalsMade         int       `json:"field_goals_made"`
	FieldGoalsAttempted    int       `json:"field_goals_attempted"`
	ThreePointersMade      int       `json:"three_pointers_made"`
	ThreePointersAttempted int       `json:"three_pointers_attempted"`
	FreeThrowsMade         int       `json:"free_throws_made"`
	FreeThrowsAttempted    int       `json:"free_throws_attempted"`
	Turnovers              int       `json:"turnovers"`
	PersonalFouls          int       `json:"personal_fouls"`
	Games                  int       `json:"games"`
	Reason                 string    `json:"reason"`
	TeamID                 string    `json:"team_id"`
	ImageUrls              ImageUrls `json:"image_urls"`
}
type BoxscoreTotals struct {
	Mins                   int `json:"mins"`
	Points                 int `json:"points"`
	ReboundsTotal          int `json:"rebounds_total"`
	Assists                int `json:"assists"`
	Steals                 int `json:"steals"`
	BlockedShots           int `json:"blocked_shots"`
	FieldGoalsMade         int `json:"field_goals_made"`
	FieldGoalsAttempted    int `json:"field_goals_attempted"`
	ThreePointersMade      int `json:"three_pointers_made"`
	ThreePointersAttempted int `json:"three_pointers_attempted"`
	FreeThrowsMade         int `json:"free_throws_made"`
	FreeThrowsAttempted    int `json:"free_throws_attempted"`
	OffensiveRebounds      int `json:"offensive_rebounds"`
	PersonalFouls          int `json:"personal_fouls"`
	Turnovers              int `json:"turnovers"`
	DefensiveRebounds      int `json:"defensive_rebounds"`
}
type TeamStandings struct {
	Wins                  int     `json:"wins"`
	Losses                int     `json:"losses"`
	GamesBehind           float64 `json:"games_behind"`
	ConferenceGamesBehind float64     `json:"conference_games_behind"`
	WinPercentage         string  `json:"win_percentage"`
	HomeRecord            string  `json:"home_record"`
	RoadRecord            string  `json:"road_record"`
	DivisionRecord        string  `json:"division_record"`
	ConferenceRecord      string  `json:"conference_record"`
	Last10Record          string  `json:"last10_record"`
	Streak                string  `json:"streak"`
}
type TeamWeb struct {
	ID               string         `json:"id"`
	SrTeamUUID       string         `json:"sr_team_uuid"`
	Name             string         `json:"name"`
	ShortName        string         `json:"short_name"`
	City             string         `json:"city"`
	ImageURL         string         `json:"image_url"`
	Score            int            `json:"score"`
	ImageURL90       string         `json:"image_url_90"`
	Division         Division       `json:"division"`
	Conference       Conference     `json:"conference"`
	Injuries         []Injuries     `json:"injuries"`
	SeasonStats      SeasonStats    `json:"season_stats"`
	GameStats        GameStats      `json:"game_stats"`
	GameLeaders      GameLeaders    `json:"game_leaders"`
	SeasonLeaders    SeasonLeaders  `json:"season_leaders"`
	Starters         []PlayerStats  `json:"starters"`
	Bench            []PlayerStats  `json:"bench"`
	Color            string         `json:"color"`
	BoxscoreTotals   BoxscoreTotals `json:"boxscore_totals"`
	FullTimeoutsLeft int            `json:"full_timeouts_left"`
	ImageURL25       string         `json:"image_url_25"`
	ImageURL59       string         `json:"image_url_59"`
	TeamStandings    TeamStandings  `json:"team_standings"`
}

type Quarters struct {
	Quarter           int `json:"quarter"`
	HomeTeamScore     int `json:"home_team_score"`
	VisitingTeamScore int `json:"visiting_team_score"`
	HomeTeamFouls     int `json:"home_team_fouls"`
	VisitingTeamFouls int `json:"visiting_team_fouls"`
}
