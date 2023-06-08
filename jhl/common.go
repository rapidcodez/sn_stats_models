package jhl

type Division struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
}
type Conference struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Rank      int    `json:"rank"`
}
type SeasonStats struct {
	Wins           int     `json:"wins"`
	Losses         int     `json:"losses"`
	GamesPlayed    int     `json:"games_played"`
	OvertimeLosses int     `json:"overtime_losses"`
	Points         int     `json:"points"`
	WinPercentage  float64 `json:"win_percentage"`
	ShootoutLosses int     `json:"shootout_losses"`
	OvertimeWins   int     `json:"overtime_wins"`
}
type GameStats struct {
	Shots              int `json:"shots"`
	PenaltiesInMinutes int `json:"penalties_in_minutes"`
}
type Skaters struct {
	ID               int    `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	ImageURL         string `json:"image_url"`
	Position         string `json:"position"`
	ShortPosition    string `json:"short_position"`
	Goals            int    `json:"goals"`
	Assists          int    `json:"assists"`
	Shots            int    `json:"shots"`
	PenaltyInMinutes int    `json:"penalty_in_minutes"`
	PlusMinusRatio   int    `json:"plus_minus_ratio"`
	ShotsOnGoal      int    `json:"shots_on_goal"`
}
type Goalies struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ImageURL       string `json:"image_url"`
	Position       string `json:"position"`
	ShotsAgainst   int    `json:"shots_against"`
	GoalsAgainst   int    `json:"goals_against"`
	Saves          int    `json:"saves"`
	Seconds        int    `json:"seconds"`
	Minutes        int    `json:"minutes"`
	SavePercentage string `json:"save_percentage,omitempty"`
}

type TeamMobile struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	City      string `json:"city"`
	Score     int    `json:"score"`
}
type TeamWeb struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	ShortName   string      `json:"short_name"`
	City        string      `json:"city"`
	ImageURL    string      `json:"image_url"`
	Score       int         `json:"score"`
	ImageURL90  string      `json:"image_url_90"`
	Division    Division    `json:"division"`
	Conference  Conference  `json:"conference"`
	SeasonStats SeasonStats `json:"season_stats"`
	GameStats   GameStats   `json:"game_stats"`
	Skaters     []Skaters   `json:"skaters"`
	Goalies     []Goalies   `json:"goalies"`
	Color       string      `json:"color"`
	ImageURL25  string      `json:"image_url_25"`
	ImageURL59  string      `json:"image_url_59"`
}

type AssistingPlayers struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ImageURL  string `json:"image_url"`
}
type Goals struct {
	Minutes          int                `json:"minutes"`
	Seconds          int                `json:"seconds"`
	TeamID           int                `json:"team_id"`
	PlayerID         int                `json:"player_id"`
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	GoalNumber       int                `json:"goal_number"`
	AssistingPlayers []AssistingPlayers `json:"assisting_players"`
}
type Periods struct {
	VisitingTeamScore int     `json:"visiting_team_score"`
	HomeTeamScore     int     `json:"home_team_score"`
	Goals             []Goals `json:"goals,omitempty"`
}
