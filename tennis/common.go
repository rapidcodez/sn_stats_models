package tennis

type TennisTournamentMatchPlayer struct {
	Id             string `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	CountryFlag    string `json:"country_flag"`
	LineScore      []int  `json:"line_score"`
	TieBreakScore  []int  `json:"tie_break_score"`
	Winner         bool   `json:"winner"`
	TournamentRank int    `json:"tournament_rank"`
	Seed           int    `json:"seed,omitempty"`
}

type TennisTournamentMatch struct {
	Id                  int                         `json:"id"`
	Round               string                      `json:"round"`
	Start               int                         `json:"start"`
	Status              string                      `json:"status"`
	PlayerOne           TennisTournamentMatchPlayer `json:"player_one"`
	PlayerTwo           TennisTournamentMatchPlayer `json:"player_two"`
	LoserWithdraw       string                      `json:"loser_withdraw"`
	LoserWithdrawReason string                      `json:"loser_withdraw_reason"`
}

type TennisTournament struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Year      string `json:"year"`

	CompetitionId     string                            `json:"competition_id"`
	Status            string                            `json:"status"`
	Prize             int                               `json:"prize"`
	Country           string                            `json:"country"`
	City              string                            `json:"city"`
	State             string                            `json:"state"`
	Venue             string                            `json:"venue"`
	Surface           string                            `json:"surface"`
	DefendingChampion TennisTournamentDefendingChampion `json:"defending_champion"`
	Matches           []TennisTournamentMatch           `json:"matches"`
}

type TennisTournamentDefendingChampion struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FlagUrl   string `json:"flag_url"`
}
