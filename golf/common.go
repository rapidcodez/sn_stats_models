package golf

type Tournament struct {
	ID                 string   `json:"id"`
	Course             string   `json:"course"`
	CurrentChampions   []Player `json:"current_champions"`
	DefendingChampion  Player   `json:"defending_champion"`
	DefendingChampions []Player `json:"defending_champions"`
	EndDate            string   `json:"end_date"`
	End                int      `json:"end"`
	Location           string   `json:"location"`
	Name               string   `json:"name"`
	Players            []Player `json:"players"`
	Prize              float64  `json:"prize"`
	Result             Result   `json:"result,omitempty"`
	ShortName          string   `json:"short_name"`
	Start              int      `json:"start"`
	StartDate          string   `json:"start_date"`
	Status             string   `json:"status"`
	Par                int      `json:"par,omitempty"`
	Yards              int      `json:"yards,omitempty"`
}

type Player struct {
	ID              string `json:"id"`
	FirstName       string `json:"first_name"`
	FlagURL         string `json:"flag_url"`
	LastName        string `json:"last_name"`
	MadeCut         bool   `json:"made_cut,omitempty"`
	Playoff         bool   `json:"playoff,omitempty"`
	Rank            string `json:"rank,omitempty"`
	Round           int    `json:"round,omitempty"`
	RoundHoles      int    `json:"round_holes,omitempty"`
	RoundScore      int    `json:"round_score,omitempty"`
	Rounds          []int  `json:"rounds,omitempty"`
	TeeTime         int    `json:"tee_time,omitempty"`
	TotalScore      int    `json:"total_score,omitempty"`
	TotalStrokes    int    `json:"total_strokes,omitempty"`
	TotalStrokesWeb int    `json:"total_strokes_web,omitempty"`
	Withdrawn       bool   `json:"withdrawn,omitempty"`
}

type Result struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Purse     float64 `json:"purse"`
	Score     int     `json:"score"`
	Strokes   int     `json:"strokes"`
}
