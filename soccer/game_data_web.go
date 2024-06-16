package soccer

import "github.com/digitalmedia/sn_stats_models/common"

type GameDataWeb struct {
	Details      DetailsWeb     `json:"details"`
	VisitingTeam TeamWeb        `json:"visiting_team"`
	HomeTeam     TeamWeb        `json:"home_team"`
	Periods      []SoccerPeriod `json:"periods"`
}

type DetailsWeb struct {
	LeagueShortName    string                 `json:"league_short_name"`
	ID                 string                 `json:"id"`
	SrGameUuid         string                 `json:"sr_game_uuid"`
	Period             int                    `json:"period"`
	IsActive           bool                   `json:"is_active"`
	Clock              string                 `json:"clock"`          // "match_clock":"45:00"
	Time               int                    `json:"time"`           // "match_time":45,
	StoppageClock      string                 `json:"stoppage_clock"` // "stoppage_time_clock":"45:00"
	StoppageTime       int                    `json:"stoppage_time"`  // "stoppage_time":45,
	Type               string                 `json:"type"`
	Timestamp          int                    `json:"timestamp"`
	Status             string                 `json:"status"`
	IsHalfTime         bool                   `json:"ishalftime"`
	HasAggregateScores bool                   `json:"hasAggregateScores"`
	AggregateWinner    string                 `json:"aggregate_winner"`
	Location           string                 `json:"location"`
	LocationImageURL   string                 `json:"location_image_url"`
	ExtraTime          bool                   `json:"extra_time"`
	Overtime           int                    `json:"overtime"`
	PenaltyKicks       bool                   `json:"penalty_kicks"`
	Broadcast          []common.GameBroadcast `json:"broadcast"`
	Last5Meetings      []common.RecentGames   `json:"last_5_meetings"`
	LocationDetails    SoccerLocationDetails  `json:"location_details"`
	Tbd                bool                   `json:"tbd"`
	Attendance         int                    `json:"attendance"`
	SimMode            bool                   `json:"sim_mode"`
	StatusLabel        int                    `json:"status_label"`
}

type SoccerLocationDetails struct {
	Venue   string `json:"venue"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type TeamWeb struct {
	ID               string                  `json:"id"`
	SrTeamUUID       string                  `json:"sr_team_uuid"`
	Name             string                  `json:"name"`
	ShortName        string                  `json:"short_name"`
	City             string                  `json:"city"`
	Country          string                  `json:"country"`
	ImageURL         string                  `json:"image_url"`
	Color            string                  `json:"color"`
	Score            int                     `json:"score"`
	PenaltyScore     int                     `json:"penalty_score"`
	AggregateScore   *int                    `json:"aggregate_score"` // Pointer so during unmarshalling it can be null when the json is missing the field, will handle the GO defaulting to zero values case
	ImageURL90       string                  `json:"image_url_90"`
	ImageURL25       string                  `json:"image_url_25"`
	ImageURL59       string                  `json:"image_url_59"`
	Division         SoccerDivision          `json:"division"`
	Conference       SoccerConference        `json:"conference"`
	Linescore        SoccerLineScore         `json:"linescore"`
	GameStats        GameStatistics          `json:"game_stats"`
	Goalkeepers      []SoccerGoalKeeper      `json:"goalkeepers"`
	Defenders        []SoccerDefender        `json:"defenders"`
	Midfielders      []SoccerMidfielder      `json:"midfielders"`
	Forwards         []SoccerForward         `json:"forwards"`
	SeasonStats      SoccerSeasonStatistics  `json:"season_stats"`
	LeagueRank       int                     `json:"league_rank"`
	TeamStats        SoccerTeamStatistics    `json:"team_stats"`
	BenchPlayers     []SoccerBenchPlayer     `json:"bench_players"`
	BenchGoalKeepers []SoccerBenchGoalKeeper `json:"bench_goalkeepers"`
	TeamStanding     SoccerTeamStanding      `json:"team_standings"`
	BoxScoreTotal    SoccerBoxScoreTotal     `json:"boxscore_totals"`
}

type SoccerConference struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
	ImageURL  string `json:"image_url"`
	Color     string `json:"color"`
}

type SoccerDivision struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
	ImageURL  string `json:"image_url"`
	Color     string `json:"color"`
}

type SoccerCard struct {
	Type           string `json:"type"`
	TeamID         string `json:"team_id"`
	PlayerID       int    `json:"player_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Minutes        int    `json:"minutes"`
	Seconds        int    `json:"seconds"`
	AdditionalMins int    `json:"additional_mins"`
}

type SoccerGoal struct {
	PlayerID       int    `json:"player_id"`
	TeamID         string `json:"team_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Minutes        int    `json:"minutes"`
	Seconds        int    `json:"seconds"`
	AdditionalMins int    `json:"additional_mins"`
	PenaltyKick    bool   `json:"penalty_kick"`
	OwnGoal        bool   `json:"own_goal"`
	ShootoutGoal   bool   `json:"shootout_goal"`
}

type SoccerSubstitution struct {
	TeamID         string                  `json:"team_id"`
	Minutes        int                     `json:"minutes"`
	Seconds        int                     `json:"seconds"`
	AdditionalMins int                     `json:"additional_mins"`
	In             SoccerSubstitutedPlayer `json:"in"`
	Out            SoccerSubstitutedPlayer `json:"out"`
}
type SoccerSubstitutedPlayer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SoccerPeriod struct {
	VisitingTeamScore int                  `json:"visiting_team_score"`
	HomeTeamScore     int                  `json:"home_team_score"`
	Goals             []SoccerGoal         `json:"goals"`
	Cards             []SoccerCard         `json:"cards"`
	Substitutions     []SoccerSubstitution `json:"substitutions"`
	ShootOutAttemps   []SoccerSOAttempt    `json:"shootout_attempts"`
}

type SoccerSOAttempt struct {
	PlayerID              int    `json:"player_id"`
	TeamID                string `json:"team_id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Minutes               int    `json:"minutes"`
	Seconds               int    `json:"seconds"`
	AdditionalMins        int    `json:"additional_mins"`
	Description           string `json:"description"`
	ShotResultID          int    `json:"shot_result_id"`
	ShotResultDescription string `json:"shot_result_description"`
	ShotDescription       string `json:"shot_description"`
}

type SoccerLineScore struct {
	Score    int                 `json:"score"`
	Shots    int                 `json:"shots"`
	Halfs    []SoccerHalf        `json:"halfs"`
	Shootout *SoccerHalfShootOut `json:"shootout"`
}

type SoccerHalf struct {
	Half  int `json:"half"`
	Score int `json:"score"`
	Shots int `json:"shots"`
}

type SoccerHalfShootOut struct {
	Goals int `json:"goals"`
}

type SoccerSeasonStatistics struct {
	Wins                    int    `json:"wins"`
	Ties                    int    `json:"ties"`
	Losses                  int    `json:"losses"`
	Points                  int    `json:"points"`
	Pct                     int    `json:"pct"`
	Fouls                   int    `json:"fouls"`
	YellowCards             int    `json:"yellow_cards"`
	RedCards                int    `json:"red_cards"`
	PointsPerGame           string `json:"points_per_game"`
	PointsTeam              int    `json:"points_team"`
	GoalsPerGame            string `json:"goals_per_game"`
	GoalsPerGameRank        int    `json:"goals_per_game_rank"`
	GoalsAgainstPerGame     string `json:"goals_against_per_game"`
	GoalsAgainstPerGameRank int    `json:"goals_against_per_game_rank"`
	GoalDifferential        int    `json:"goal_differential"`
	GoalDifferentialRank    int    `json:"goal_differential_rank"`
	Shutouts                int    `json:"shutouts"`
	ShutoutsRank            int    `json:"shutouts_rank"`
	Minutes                 int    `json:"minutes"`
	Tackles                 int    `json:"tackles"`
}

type SoccerTeamStatistics struct {
	Fouls                   int    `json:"fouls"`
	YellowCards             int    `json:"yellow_cards"`
	RedCards                int    `json:"red_cards"`
	GoalsPerGame            string `json:"goals_per_game"`
	GoalsPerGameRank        int    `json:"goals_per_game_rank"`
	GoalsAgainstPerGame     string `json:"goals_against_per_game"`
	GoalsAgainstPerGameRank int    `json:"goals_against_per_game_rank"`
	GoalDifferential        int    `json:"goal_differential"`
	GoalDifferentialRank    int    `json:"goal_differential_rank"`
	Shutouts                int    `json:"shutouts"`
	ShutoutsRank            int    `json:"shutouts_rank"`
	Minutes                 int    `json:"minutes"`
	Tackles                 int    `json:"tackles"`
}

type SoccerBasePlayer struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ImageURL      string `json:"image_url"`
	Position      string `json:"position"`
	ShortPosition string `json:"short_position"`
	Number        int    `json:"number"`
	Minutes       int    `json:"minutes"`
	Type          string `json:"type"`
	Starter       bool   `json:"starter"`
	Played        bool   `json:"played"`
}

type SoccerDefender struct {
	SoccerBasePlayer
	Goals            int    `json:"goals"`
	Assists          int    `json:"assists"`
	Shots            int    `json:"shots"`
	ShotsOnGoal      int    `json:"shots_on_goal"`
	Crosses          int    `json:"crosses"`
	Tackles          int    `json:"tackles"`
	Blocks           int    `json:"blocks"`
	Touches          int    `json:"touches"`
	Passes           int    `json:"passes"`
	Fouls            int    `json:"fouls"`
	Offsides         int    `json:"offsides"`
	CornerKicks      int    `json:"corner_kicks"`
	YellowCards      int    `json:"yellow_cards"`
	RedCards         int    `json:"red_cards"`
	OwnGoals         int    `json:"own_goals"`
	PenaltyKick      bool   `json:"penalty_kick"`
	PassesPercentage string `json:"passes_percentage"`
}

type SoccerMidfielder struct {
	SoccerBasePlayer
	Goals            int    `json:"goals"`
	Assists          int    `json:"assists"`
	Shots            int    `json:"shots"`
	ShotsOnGoal      int    `json:"shots_on_goal"`
	Crosses          int    `json:"crosses"`
	Tackles          int    `json:"tackles"`
	Blocks           int    `json:"blocks"`
	Touches          int    `json:"touches"`
	Passes           int    `json:"passes"`
	Fouls            int    `json:"fouls"`
	Offsides         int    `json:"offsides"`
	CornerKicks      int    `json:"corner_kicks"`
	YellowCards      int    `json:"yellow_cards"`
	RedCards         int    `json:"red_cards"`
	OwnGoals         int    `json:"own_goals"`
	PenaltyKick      bool   `json:"penalty_kick"`
	PassesPercentage string `json:"passes_percentage"`
}

type SoccerBenchGoalKeeper struct {
	SoccerBasePlayer
}

type SoccerBenchPlayer struct {
	SoccerBasePlayer
}

type SoccerGoalKeeper struct {
	SoccerBasePlayer
	GoalsAgainst int `json:"goals_against"`
	Saves        int `json:"saves"`
	ShotsAgainst int `json:"shots_against"`
	ShotsOnGoal  int `json:"shots_on_goal"`
}

type SoccerForward struct {
	SoccerBasePlayer
	Goals            int    `json:"goals"`
	Assists          int    `json:"assists"`
	Shots            int    `json:"shots"`
	ShotsOnGoal      int    `json:"shots_on_goal"`
	Crosses          int    `json:"crosses"`
	Tackles          int    `json:"tackles"`
	Blocks           int    `json:"blocks"`
	Touches          int    `json:"touches"`
	Passes           int    `json:"passes"`
	Fouls            int    `json:"fouls"`
	Offsides         int    `json:"offsides"`
	CornerKicks      int    `json:"corner_kicks"`
	YellowCards      int    `json:"yellow_cards"`
	RedCards         int    `json:"red_cards"`
	OwnGoals         int    `json:"own_goals"`
	PenaltyKick      bool   `json:"penalty_kick"`
	PassesPercentage string `json:"passes_percentage"`
}

type SoccerTeamStanding struct {
	Wins              int    `json:"wins"`
	Losses            int    `json:"losses"`
	Draws             int    `json:"draws"`
	Points            int    `json:"points"`
	GamesPlayed       int    `json:"games_played"`
	GoalsFor          int    `json:"goals_for"`
	GoalsAgainst      int    `json:"goals_against"`
	GoalDifferential  int    `json:"goal_differential"`
	Last10Wins        int    `json:"last_10_wins"`
	Last10Losses      int    `json:"last_10_losses"`
	Last10Ties        int    `json:"last_10_ties"`
	HomeWins          int    `json:"home_wins"`
	HomeLosses        int    `json:"home_losses"`
	HomeTies          int    `json:"home_ties"`
	VisitingWins      int    `json:"visiting_wins"`
	VisitingLosses    int    `json:"visiting_losses"`
	VisitingTies      int    `json:"visiting_ties"`
	Streak            string `json:"streak"`
	WinningPercentage string `json:"winning_percentage"`
}

type SoccerBoxScoreTotal struct {
	Goals       int    `json:"goals"`
	Assists     int    `json:"assists"`
	Shots       int    `json:"shots"`
	ShotsOnGoal int    `json:"shots_on_goal"`
	Touches     int    `json:"touches"`
	Passes      int    `json:"passes"`
	Percentage  string `json:"percentage"`
	Tackles     int    `json:"tackles"`
	Fouls       int    `json:"fouls"`
}

type GameStatistics struct {
	PossessionPercentage int    `json:"possession_percentage"`
	CardsGiven           int    `json:"cards_given"`
	Goals                int    `json:"goals"`
	Assists              int    `json:"assists"`
	Shots                int    `json:"shots"`
	ShotsOnGoal          int    `json:"shots_on_goal"`
	ShotsBlocked         int    `json:"shots_blocked"`
	ShotsOffTarget       int    `json:"shots_off_target"`
	ShotsOnTarget        int    `json:"shots_on_target"`
	ShotsSaved           int    `json:"shots_saved"`
	ShotsTotal           int    `json:"shots_total"`
	CornerKicks          int    `json:"corner_kicks"`
	FreeKicks            int    `json:"free_kicks"`
	GoalKicks            int    `json:"goal_kicks"`
	Injuries             int    `json:"injuries"`
	Offsides             int    `json:"offsides"`
	Fouls                int    `json:"fouls"`
	Crosses              int    `json:"crosses"`
	YellowCards          int    `json:"yellow_cards"`
	RedCards             int    `json:"red_cards"`
	Substitutions        int    `json:"substitutions"`
	YellowRedCards       int    `json:"yellow_red_cards"`
	ThrowIns             int    `json:"throw_ins"`
	Saves                int    `json:"saves"`
	Minutes              int    `json:"minutes"`
	PassesPercentage     string `json:"passes_percentage"`
	Tackles              int    `json:"tackles"`
}
