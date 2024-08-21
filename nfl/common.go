package nfl

import (
	"github.com/digitalmedia/sn_stats_models/common"
	"time"
)

type QuarterPlay struct {
	VisitingTeamScore int    `json:"visiting_team_score"`
	HomeTeamScore     int    `json:"home_team_score"`
	TeamId            string `json:"team_id"`
	Time              string `json:"time"`
	Play              string `json:"play"`
}

type Quarter struct {
	Plays             []QuarterPlay `json:"plays"`
	VisitingTeamScore int           `json:"visiting_team_score"`
	HomeTeamScore     int           `json:"home_team_score"`
}

type TeamMobile struct {
	Id         int    `json:"id"`
	TeamID     string `json:"team_id"`
	Name       string `json:"name"`
	SrTeamUUID string `json:"sr_team_uuid"`
	ShortName  string `json:"short_name"`
	HasBall    bool   `json:"has_ball"`
	City       string `json:"city"`
	Score      int    `json:"score"`
}

type SeasonStats struct {
	Wins                          int     `json:"wins"`
	Losses                        int     `json:"losses"`
	Ties                          int     `json:"ties"`
	PassingYardsPerGame           float64 `json:"passing_yards_per_game"`
	NetPassingYardsPerGame        float64 `json:"net_passing_yards_per_game"`
	RushingYardsPerGame           float64 `json:"rushing_yards_per_game"`
	OffenseYardsPerGame           float64 `json:"offense_yards_per_game"`
	DefenseYardsPerGame           float64 `json:"defense_yards_per_game"`
	Tackles                       int     `json:"tackles"`
	Interceptions                 int     `json:"interceptions"`
	FumblesRecovered              int     `json:"fumbles_recovered"`
	PointsScored                  int     `json:"points_scored"`
	PenaltiesPerGame              float64 `json:"penalties_per_game"`
	Touchdowns                    int     `json:"touchdowns"`
	ReturnYards                   int     `json:"return_yards"`
	Sacks                         int     `json:"sacks"`
	YardsPerGame                  float64 `json:"yards_per_game"`
	YardsAllowedPerGame           float64 `json:"yards_allowed_per_game"`
	ReturnYardsPerGame            float64 `json:"return_yards_per_game"`
	InterceptionsThrown           int     `json:"interceptions_thrown"`
	FumblesLost                   int     `json:"fumbles_lost"`
	PointsScoredPerGame           float64 `json:"points_scored_per_game"`
	PointsAllowedPerGame          float64 `json:"points_allowed_per_game"`
	SacksAllowed                  int     `json:"sacks_allowed"`
	PuntingSingles                int     `json:"punting_singles"`
	FieldGoalsSingles             int     `json:"field_goals_singles"`
	KickoffsSingles               int     `json:"kickoffs_singles"`
	ThirdDownEfficiencyPercentage float64 `json:"third_down_efficiency_percentage"`
	Takeaways                     int     `json:"takeaways"`
	ScoringEfficiencyPercentage   float64 `json:"scoring_efficiency_percentage"`
	HundredYardGamesRushing       int     `json:"hundred_yard_games_rushing"`
	HundredYardGamesReceiving     int     `json:"hundred_yard_games_receiving"`
	ThreeHundredYardGamesPassing  int     `json:"three_hundred_yard_games_passing"`
}

type TeamStandings struct {
	Wins                   int     `json:"wins"`
	Losses                 int     `json:"losses"`
	Ties                   int     `json:"ties"`
	WinPercentage          float64 `json:"win_percentage"`
	Points                 int     `json:"points"`
	PointsAgainst          int     `json:"points_against"`
	HomeRecord             string  `json:"home_record"`
	RoadRecord             string  `json:"road_record"`
	DivisionRecord         string  `json:"division_record"`
	ConferenceRecord       string  `json:"conference_record"`
	PointsFor              int     `json:"points_for"`
	HomeWins               int     `json:"home_wins"`
	HomeLosses             int     `json:"home_losses"`
	HomeTies               int     `json:"home_ties"`
	AwayWins               int     `json:"away_wins"`
	AwayLosses             int     `json:"away_losses"`
	AwayTies               int     `json:"away_ties"`
	Last5Wins              int     `json:"last_5_wins"`
	Last5Losses            int     `json:"last_5_losses"`
	Last5Ties              int     `json:"last_5_ties"`
	StrengthOfScheduleRank int     `json:"strength_of_schedule_rank"`
}

type SeasonLeadersLeader struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Number        int    `json:"number"`
	Position      string `json:"position"`
	ShortPosition string `json:"short_position"`
	ImageUrl      string `json:"image_url"`
	ImageUrls     struct {
		Lg string `json:"lg"`
		Md string `json:"md"`
		Sm string `json:"sm"`
		Xs string `json:"xs"`
	} `json:"image_urls"`

	Yards      int    `json:"yards"`
	Average    string `json:"average"`
	Touchdowns int    `json:"touchdowns"`
	Attempts   int    `json:"attempts"`

	QbRating float64 `json:"qb_rating"`

	PassesAttempted int `json:"passes_attempted"`

	Interceptions        int    `json:"interceptions"`
	CompletionPercentage string `json:"completion_percentage"`
	LongestPass          int    `json:"longest_pass"`

	FieldGoalsAttempted  int     `json:"field_goals_attempted"`
	FieldGoalsMade       int     `json:"field_goals_made"`
	FieldGoalsPercentage float64 `json:"field_goals_percentage"`
	ExtraPointsMade      int     `json:"extra_points_made"`

	Tackles int     `json:"tackles"`
	Sacks   float64 `json:"sacks"`

	ForcedFumbles           int `json:"forced_fumbles"`
	Assists                 int `json:"assists"`
	TotalTackles            int `json:"total_tackles"`
	InterceptionsYards      int `json:"interceptions_yards"`
	InterceptionsTouchdowns int `json:"interceptions_touchdowns"`

	Receptions        int     `json:"receptions"`
	LongestReceiving  int     `json:"longest_receiving"`
	YardsPerReception float64 `json:"yards_per_reception"`
	Targets           int     `json:"targets"`

	YardsPerGame    float64 `json:"yards_per_game"`
	YardsPerAttempt float64 `json:"yards_per_attempt"`

	Completions  int `json:"completions"`
	GamesStarted int `json:"games_started"`
	RushingYards int `json:"rushing_yards"`
}

type SeasonPassingLeader struct {
}

type SeasonLeaders struct {
	PassingLeader        *SeasonLeadersLeader `json:"passing_leader,omitempty"`
	RushingLeader        *SeasonLeadersLeader `json:"rushing_leader,omitempty"`
	ReceivingLeader      *SeasonLeadersLeader `json:"receiving_leader,omitempty"`
	TacklesLeader        *SeasonLeadersLeader `json:"tackles_leader,omitempty"`
	SacksLeader          *SeasonLeadersLeader `json:"sacks_leader,omitempty"`
	InterceptionsLeader  *SeasonLeadersLeader `json:"interceptions_leader,omitempty"`
	ForcedFumblesLeader  *SeasonLeadersLeader `json:"forced_fumbles_leader,omitempty"`
	FieldGoalsMadeLeader *SeasonLeadersLeader `json:"field_goals_made_leader,omitempty"`
}

type TeamWebGameStats struct {
	FirstDowns                    int     `json:"first_downs"`
	RushingYards                  int     `json:"rushing_yards"`
	PassingYards                  int     `json:"passing_yards"`
	TotalYards                    int     `json:"total_yards"`
	PossessionMinutes             int     `json:"possession_minutes"`
	PossessionSeconds             int     `json:"possession_seconds"`
	Sacks                         float64 `json:"sacks"`
	InterceptionsThrown           int     `json:"interceptions_thrown"`
	FumblesLost                   int     `json:"fumbles_lost"`
	Turnovers                     int     `json:"turnovers"`
	PenaltyYards                  int     `json:"penalty_yards"`
	Tackles                       int     `json:"tackles"`
	FumblesRecovered              int     `json:"fumbles_recovered"`
	PointsScored                  int     `json:"points_scored"`
	Penalties                     int     `json:"penalties"`
	Touchdowns                    int     `json:"touchdowns"`
	ReturnYards                   int     `json:"return_yards"`
	PuntingSingles                int     `json:"punting_singles"`
	FieldGoalsSingles             int     `json:"field_goals_singles"`
	KickoffsSingles               int     `json:"kickoffs_singles"`
	YardsLost                     int     `json:"yards_lost"`
	ThirdDownEfficiencyPercentage float64 `json:"third_down_efficiency_percentage"`
	GameTotalPlays                int     `json:"game_total_plays"`
	GameTotalAverage              string  `json:"game_total_average"`
	PassingSacks                  float64 `json:"passing_sacks"`
}

type GameLeader struct {
	Id                   string `json:"id"`
	PassesAttempted      int    `json:"passes_attempted"`
	PassesCompleted      int    `json:"passes_completed"`
	Interceptions        int    `json:"interceptions"`
	Number               int    `json:"number"`
	Average              string `json:"average"`
	Touchdowns           int    `json:"touchdowns"`
	Yards                int    `json:"yards"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	ImageUrl             string `json:"image_url"`
	Position             string `json:"position"`
	CompletionPercentage string `json:"completion_percentage"`
	LongestPass          int    `json:"longest_pass"`
	RushesAttempted      int    `json:"rushes_attempted"`
	LongestRush          int    `json:"longest_rush"`
	Receptions           int    `json:"receptions"`
	LongestReceiving     int    `json:"longest_receiving"`
	Targets              int    `json:"targets"`
	Made                 int    `json:"made"`
	Attempted            int    `json:"attempted"`
	Percentage           string `json:"percentage"`
	LongestKick          int    `json:"longest_kick"`
	ExtraPointsMade      int    `json:"extra_points_made"`
	FieldGoalsMissed     int    `json:"field_goals_missed"`
	ExtraPointsAttempted int    `json:"extra_points_attempted"`

	ImageUrls struct {
		Lg string `json:"lg"`
		Md string `json:"md"`
		Sm string `json:"sm"`
		Xs string `json:"xs"`
	} `json:"image_urls"`
}

type GameLeaders struct {
	PassingLeader           GameLeader `json:"passing_leader"`
	RushingLeader           GameLeader `json:"rushing_leader"`
	ReceivingLeader         GameLeader `json:"receiving_leader"`
	InterceptionsLeader     GameLeader `json:"interceptions_leader"`
	FieldGoalsSinglesLeader GameLeader `json:"field_goals_singles_leader"`
	KickoffsSinglesLeader   GameLeader `json:"kickoffs_singles_leader"`
}

type TeamWebPlayersPlayer struct {
	Id                   string  `json:"id"`
	FirstName            string  `json:"first_name"`
	LastName             string  `json:"last_name"`
	Number               int     `json:"number"`
	Yards                int     `json:"yards"`
	PassesAttempted      int     `json:"passes_attempted"`
	PassesCompleted      int     `json:"passes_completed"`
	Interceptions        int     `json:"interceptions"`
	Average              string  `json:"average"`
	Touchdowns           int     `json:"touchdowns"`
	Position             string  `json:"position"`
	CompletionPercentage string  `json:"completion_percentage"`
	LongestPass          int     `json:"longest_pass"`
	RushesAttempted      int     `json:"rushes_attempted"`
	LongestRush          int     `json:"longest_rush"`
	Receptions           int     `json:"receptions"`
	LongestReceiving     int     `json:"longest_receiving"`
	Targets              int     `json:"targets"`
	Returns              int     `json:"returns"`
	LongestPuntReturn    int     `json:"longest_punt_return"`
	Made                 int     `json:"made"`
	Attempted            int     `json:"attempted"`
	Percentage           string  `json:"percentage"`
	LongestKick          int     `json:"longest_kick"`
	ExtraPointsMade      int     `json:"extra_points_made"`
	FieldGoalsMissed     int     `json:"field_goals_missed"`
	ExtraPointsAttempted int     `json:"extra_points_attempted"`
	Tackles              int     `json:"tackles"`
	Sacks                float64 `json:"sacks"`
	SackYards            float64 `json:"sack_yards"`
	PassesDefensed       int     `json:"passes_defensed"`
	SpTackles            int     `json:"sp_tackles"`

	ImageUrl  string `json:"image_url"`
	ImageUrls struct {
		Lg string `json:"lg"`
		Md string `json:"md"`
		Sm string `json:"sm"`
		Xs string `json:"xs"`
	} `json:"image_urls"`
}

type TeamWebPlayers struct {
	Passing       []TeamWebPlayersPlayer `json:"passing,omitempty"`
	Rushing       []TeamWebPlayersPlayer `json:"rushing,omitempty"`
	Receiving     []TeamWebPlayersPlayer `json:"receiving,omitempty"`
	Interceptions []TeamWebPlayersPlayer `json:"interceptions,omitempty"`
	PuntReturns   []TeamWebPlayersPlayer `json:"punt_returns,omitempty"`
	Kickings      []TeamWebPlayersPlayer `json:"kickings,omitempty"`
	Defense       []TeamWebPlayersPlayer `json:"defense,omitempty"`
	Fumbles       []TeamWebPlayersPlayer `json:"fumbles,omitempty"`
}

type TeamWebDivision struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
}

type TeamWeb struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	ShortName     string            `json:"short_name"`
	City          string            `json:"city"`
	ImageUrl      string            `json:"image_url"`
	Score         int               `json:"score"`
	ImageUrl90    string            `json:"image_url_90"`
	Division      TeamWebDivision   `json:"division"`
	Injuries      []common.Injuries `json:"injuries"`
	GameStats     TeamWebGameStats  `json:"game_stats"`
	SeasonStats   *SeasonStats      `json:"season_stats"`
	GameLeaders   GameLeaders       `json:"game_leaders"`
	SeasonLeaders *SeasonLeaders    `json:"season_leaders,omitempty"`
	Players       *TeamWebPlayers   `json:"players,omitempty"`
	Color         string            `json:"color"`
	TimeoutsLeft  int               `json:"timeouts_left"`
	ImageUrl25    string            `json:"image_url_25"`
	ImageUrl59    string            `json:"image_url_59"`
	TeamStandings *TeamStandings    `json:"team_standings,omitempty"`
}

type DetailsMobile struct {
	Id              int    `json:"id"`
	SrGameUuid      string `json:"sr_game_uuid"`
	Type            string `json:"type"`
	LeagueShortName string `json:"league_short_name"`
	Quarter         int    `json:"quarter"`
	IsActive        bool   `json:"is_active"`
	Timestamp       int    `json:"timestamp"`
	Status          string `json:"status"`
	Overtime        int    `json:"overtime"`
	IsIntermission  bool   `json:"is_intermission"`
	Sequence        int64  `json:"sequence"`
	ScoreSequence   int64  `json:"score_sequence"`
	Downs           int    `json:"downs"`
	Distance        int    `json:"distance"`
	Clock           string `json:"clock"`
}

type Last9EventsPlayer struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ImageUrls struct {
		Lg string `json:"lg"`
		Md string `json:"md"`
		Sm string `json:"sm"`
		Xs string `json:"xs"`
	} `json:"image_urls"`
}

type Last9Event struct {
	Id            string              `json:"id"`
	Time          string              `json:"time"`
	Datetime      time.Time           `json:"datetime"`
	Sequence      int                 `json:"sequence"`
	Description   string              `json:"description"`
	Quarter       int                 `json:"quarter"`
	Down          int                 `json:"down"`
	Distance      int                 `json:"distance"`
	Yardline      string              `json:"yardline"`
	EndYardline   string              `json:"end_yardline"`
	Possession    string              `json:"possession"`
	EndPossession string              `json:"end_possession"`
	PlayType      string              `json:"play_type"`
	PlayTypeId    int                 `json:"play_type_id"`
	Result        string              `json:"result"`
	Players       []Last9EventsPlayer `json:"players"`
}

type DetailsWebGameType struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type DetailsWeb struct {
	Id                   string                 `json:"id"`
	LeagueShortName      string                 `json:"league_short_name"`
	Quarter              int                    `json:"quarter"`
	IsActive             bool                   `json:"is_active"`
	Clock                string                 `json:"clock"`
	Type                 string                 `json:"type"`
	Timestamp            int                    `json:"timestamp"`
	Overtime             int                    `json:"overtime"`
	IsIntermission       bool                   `json:"is_intermission"`
	Broadcast            []common.GameBroadcast `json:"broadcast"`
	Status               string                 `json:"status"`
	SrStatus             string                 `json:"sr_status"`
	SimMode              bool                   `json:"sim_mode"`
	BallLocation         string                 `json:"ball_location"`
	TeamPossessionId     int                    `json:"team_possession_id"`
	Down                 int                    `json:"down"`
	Distance             int                    `json:"distance"`
	Location             string                 `json:"location"`
	SeriesMatchupsResult string                 `json:"series_matchups_result"`
	LocationImageUrl     string                 `json:"location_image_url"`
	Attendance           int                    `json:"attendance"`
	GameName             string                 `json:"game_name"`
	Tbd                  bool                   `json:"tbd"`
	LocationImageMed     string                 `json:"location_image_med"`
	LocationImageSml     string                 `json:"location_image_sml"`
	Story                struct {
		Type     string `json:"type"`
		Headline string `json:"headline"`
		Content  string `json:"content"`
	} `json:"story"`
	Week          int                  `json:"week"`
	City          string               `json:"city"`
	Country       string               `json:"country"`
	Last9Events   []Last9Event         `json:"last_9_events,omitempty"`
	Timeout       string               `json:"timeout"`
	Last5Meetings []common.RecentGames `json:"last_5_meetings"`
	State         string               `json:"state"`
	GameType      DetailsWebGameType   `json:"game_type"`
}
