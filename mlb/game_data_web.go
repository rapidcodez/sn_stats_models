package mlb

import "github.com/digitalmedia/sn_stats_models/common"

type GameDataWeb struct {
	Details      DetailsWeb `json:"details"`
	VisitingTeam TeamWeb    `json:"visiting_team"`
	HomeTeam     TeamWeb    `json:"home_team"`
	Innings      []Inning   `json:"innings"`
}

type DetailsWeb struct {
	ID                   int                     `json:"id"`
	SrGameUuid           string                  `json:"sr_game_uuid"`
	LeagueShortName      string                  `json:"league_short_name"`
	Status               string                  `json:"status"`
	Timestamp            int                     `json:"timestamp"`
	Type                 string                  `json:"type"`
	IsAllStar            bool                    `json:"is_all_star"`
	IsPreseason          bool                    `json:"is_preseason"`
	Inning               int                     `json:"inning"`
	InningStatus         string                  `json:"inning_status"`
	Outs                 int                     `json:"outs"`
	IsActive             bool                    `json:"is_active"`
	RunnerOnFirstBase    bool                    `json:"runner_on_first_base"`
	RunnerOnSecondBase   bool                    `json:"runner_on_second_base"`
	RunnerOnThirdBase    bool                    `json:"runner_on_third_base"`
	Attendance           int                     `json:"attendance"`
	Location             string                  `json:"location"`
	LocationImageURL     string                  `json:"location_image_url"`
	WinningPitcher       *Player                 `json:"winning_pitcher,omitempty"`
	LosingPitcher        *Player                 `json:"losing_pitcher,omitempty"`
	SavingPitcher        *Player                 `json:"saving_pitcher,omitempty"`
	CurrentBatter        int                     `json:"current_batter"`
	LastPlay             string                  `json:"last_play"`
	LocationImageMed     string                  `json:"location_image_med"`
	LocationImageSml     string                  `json:"location_image_sml"`
	Weather              string                  `json:"weather"`
	Wind                 string                  `json:"wind"`
	Umpires              *[]Umpires              `json:"umpires,omitempty"`
	ExtraInnings         bool                    `json:"extra_innings"`
	SeriesMatchupsResult string                  `json:"series_matchups_result"`
	SeasonSeries         *[]SeasonSeries         `json:"season_series,omitempty"`
	CurrentAtBat         *CurrentAtBat           `json:"current_at_bat,omitempty"`
	CipbpComplete        bool                    `json:"cipbp_complete"`
	UnderReview          bool                    `json:"under_review"`
	GameDuration         string                  `json:"game_duration"`
	City                 string                  `json:"city"`
	Country              string                  `json:"country"`
	Conference           string                  `json:"conference"`
	HomePlayoffSeed      int                     `json:"home_playoff_seed"`
	VisitingPlayoffSeed  int                     `json:"visiting_playoff_seed"`
	State                string                  `json:"state"`
	OpeningLine          *OpeningLine            `json:"opening_line,omitempty"`
	CurrentLine          *CurrentLine            `json:"current_line,omitempty"`
	Broadcast            *[]common.GameBroadcast `json:"broadcast,omitempty"`
	Story                *Story                  `json:"story"`
	CurrentInningPbp     *[]CurrentInningPbp     `json:"current_inning_pbp,omitempty"`
	CipbpInning          string                  `json:"cipbp_inning"`
	RunnersOnBase        *RunnersOnBase          `json:"runners_on_base,omitempty"`
	LastOut              *LastOut                `json:"last_out,omitempty"`
	Code                 int                     `json:"code"`
}

type TeamWeb struct {
	ID                   int              `json:"id"`
	SrTeamUUID           string           `json:"sr_team_uuid"`
	Name                 string           `json:"name"`
	ShortName            string           `json:"short_name"`
	City                 string           `json:"city"`
	ImageURL             string           `json:"image_url"`
	Score                int              `json:"score"`
	ImageURL90           string           `json:"image_url_90"`
	Division             *Division        `json:"division"`
	Conference           *Conference      `json:"conference"`
	GameStats            *GameStats       `json:"game_stats"`
	SeasonStats          *SeasonStats     `json:"season_stats"`
	HomeRunLeader        *SeasonLeader    `json:"home_run_leader"`
	BattingAverageLeader *SeasonLeader    `json:"batting_average_leader"`
	RunsBattedInLeader   *SeasonLeader    `json:"runs_batted_in_leader"`
	Homeruns             *[]SeasonLeader  `json:"homeruns"`
	Batters              *[]Player        `json:"batters"`
	Pitchers             *[]Player        `json:"pitchers"`
	StartingPitcher      *StartingPitcher `json:"starting_pitcher,omitempty"`
	Color                string           `json:"color"`
	CurrentPitcher       int              `json:"current_pitcher"`
	AtBat                bool             `json:"at_bat"`
	WinsLeader           *SeasonLeader    `json:"wins_leader"`
	SavesLeader          *SeasonLeader    `json:"saves_leader"`
	Batting              *Batting         `json:"batting"`
	Pitching             *Pitching        `json:"pitching"`
	Fielding             *Fielding        `json:"fielding"`
	ImageURL59           string           `json:"image_url_59"`
	StrikeoutsLeader     *SeasonLeader    `json:"strikeouts_leader"`
	SeasonStatsStr       *SeasonStatsStr  `json:"season_stats_str"`
	ImageURL25           string           `json:"image_url_25"`
	TeamStandings        *TeamStandings   `json:"team_standings"`
	Injuries             *[]Injuries      `json:"injuries"`
}
