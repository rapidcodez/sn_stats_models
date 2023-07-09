package cfl

type Quarter struct {
	Plays []struct {
		VisitingTeamScore int    `json:"visiting_team_score"`
		HomeTeamScore     int    `json:"home_team_score"`
		TeamId            int    `json:"team_id"`
		Time              string `json:"time"`
		Play              string `json:"play"`
	} `json:"plays"`
	VisitingTeamScore int `json:"visiting_team_score"`
	HomeTeamScore     int `json:"home_team_score"`
}

type TeamMobile struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	City      string `json:"city"`
	Score     int    `json:"score"`
}

type TeamWeb struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	City       string `json:"city"`
	ImageUrl   string `json:"image_url"`
	Score      int    `json:"score"`
	ImageUrl90 string `json:"image_url_90"`
	Division   struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"short_name"`
		Rank      int    `json:"rank"`
	} `json:"division"`
	Injuries []struct {
		PlayerId  int    `json:"player_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Position  string `json:"position"`
		Status    string `json:"status"`
		Timestamp int    `json:"timestamp"`
		Type      string `json:"type"`
	} `json:"injuries"`
	GameStats struct {
		FirstDowns                    int    `json:"first_downs"`
		RushingYards                  int    `json:"rushing_yards"`
		PassingYards                  int    `json:"passing_yards"`
		TotalYards                    int    `json:"total_yards"`
		PossessionMinutes             int    `json:"possession_minutes"`
		PossessionSeconds             int    `json:"possession_seconds"`
		Sacks                         int    `json:"sacks"`
		InterceptionsThrown           int    `json:"interceptions_thrown"`
		FumblesLost                   int    `json:"fumbles_lost"`
		Turnovers                     int    `json:"turnovers"`
		PenaltyYards                  int    `json:"penalty_yards"`
		Tackles                       int    `json:"tackles"`
		FumblesRecovered              int    `json:"fumbles_recovered"`
		PointsScored                  int    `json:"points_scored"`
		Penalties                     int    `json:"penalties"`
		Touchdowns                    int    `json:"touchdowns"`
		ReturnYards                   int    `json:"return_yards"`
		PuntingSingles                int    `json:"punting_singles"`
		FieldGoalsSingles             int    `json:"field_goals_singles"`
		KickoffsSingles               int    `json:"kickoffs_singles"`
		YardsLost                     int    `json:"yards_lost"`
		ThirdDownEfficiencyPercentage int    `json:"third_down_efficiency_percentage"`
		GameTotalPlays                int    `json:"game_total_plays"`
		GameTotalAverage              string `json:"game_total_average"`
		PassingSacks                  int    `json:"passing_sacks"`
	} `json:"game_stats"`
	SeasonStats struct {
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
		ReturnYardsPerGame            int     `json:"return_yards_per_game"`
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
	} `json:"season_stats"`
	GameLeaders struct {
		PassingLeader struct {
			Id                   int    `json:"id"`
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
			ImageUrls            struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"passing_leader"`
		RushingLeader struct {
			Id              int    `json:"id"`
			RushesAttempted int    `json:"rushes_attempted"`
			LongestRush     int    `json:"longest_rush"`
			Number          int    `json:"number"`
			Average         string `json:"average"`
			Touchdowns      int    `json:"touchdowns"`
			Yards           int    `json:"yards"`
			FirstName       string `json:"first_name"`
			LastName        string `json:"last_name"`
			ImageUrl        string `json:"image_url"`
			Position        string `json:"position"`
			ImageUrls       struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"rushing_leader"`
		ReceivingLeader struct {
			Id               int    `json:"id"`
			Receptions       int    `json:"receptions"`
			LongestReceiving int    `json:"longest_receiving"`
			Number           int    `json:"number"`
			Average          string `json:"average"`
			Touchdowns       int    `json:"touchdowns"`
			Yards            int    `json:"yards"`
			FirstName        string `json:"first_name"`
			LastName         string `json:"last_name"`
			ImageUrl         string `json:"image_url"`
			Position         string `json:"position"`
			Targets          int    `json:"targets"`
			ImageUrls        struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"receiving_leader"`
		InterceptionsLeader struct {
			Id         int    `json:"id"`
			Number     int    `json:"number"`
			Touchdowns int    `json:"touchdowns"`
			Yards      int    `json:"yards"`
			FirstName  string `json:"first_name"`
			LastName   string `json:"last_name"`
			ImageUrl   string `json:"image_url"`
			Position   string `json:"position"`
			ImageUrls  struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"interceptions_leader"`
		FieldGoalsSinglesLeader struct {
			Id               int    `json:"id"`
			Made             int    `json:"made"`
			Attempted        int    `json:"attempted"`
			Percentage       string `json:"percentage"`
			LongestKick      int    `json:"longest_kick"`
			Number           int    `json:"number"`
			FirstName        string `json:"first_name"`
			LastName         string `json:"last_name"`
			ImageUrl         string `json:"image_url"`
			Position         string `json:"position"`
			ExtraPointsMade  int    `json:"extra_points_made"`
			FieldGoalsMissed int    `json:"field_goals_missed"`
			ImageUrls        struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
			ExtraPointsAttempted int `json:"extra_points_attempted"`
		} `json:"field_goals_singles_leader"`
		KickoffsSinglesLeader struct {
			Id               int    `json:"id"`
			Made             int    `json:"made"`
			Attempted        int    `json:"attempted"`
			Percentage       string `json:"percentage"`
			LongestKick      int    `json:"longest_kick"`
			Number           int    `json:"number"`
			FirstName        string `json:"first_name"`
			LastName         string `json:"last_name"`
			ImageUrl         string `json:"image_url"`
			Position         string `json:"position"`
			ExtraPointsMade  int    `json:"extra_points_made"`
			FieldGoalsMissed int    `json:"field_goals_missed"`
			ImageUrls        struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
			ExtraPointsAttempted int `json:"extra_points_attempted"`
		} `json:"kickoffs_singles_leader"`
	} `json:"game_leaders"`
	SeasonLeaders struct {
		PassingLeader struct {
			Id            int     `json:"id"`
			Interceptions int     `json:"interceptions"`
			Number        int     `json:"number"`
			Touchdowns    int     `json:"touchdowns"`
			Yards         int     `json:"yards"`
			FirstName     string  `json:"first_name"`
			LastName      string  `json:"last_name"`
			ImageUrl      string  `json:"image_url"`
			Position      string  `json:"position"`
			QbRating      float64 `json:"qb_rating"`
			ShortPosition string  `json:"short_position"`
			Attempts      int     `json:"attempts"`
			Completions   int     `json:"completions"`
			GamesStarted  int     `json:"games_started"`
			ImageUrls     struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
			RushingYards int `json:"rushing_yards"`
		} `json:"passing_leader"`
		RushingLeader struct {
			Id              int     `json:"id"`
			Number          int     `json:"number"`
			Touchdowns      int     `json:"touchdowns"`
			Yards           int     `json:"yards"`
			FirstName       string  `json:"first_name"`
			LastName        string  `json:"last_name"`
			ImageUrl        string  `json:"image_url"`
			Position        string  `json:"position"`
			ShortPosition   string  `json:"short_position"`
			Attempts        int     `json:"attempts"`
			YardsPerGame    float64 `json:"yards_per_game"`
			YardsPerAttempt float64 `json:"yards_per_attempt"`
			GamesStarted    int     `json:"games_started"`
			ImageUrls       struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"rushing_leader"`
		ReceivingLeader struct {
			Id                int     `json:"id"`
			Receptions        int     `json:"receptions"`
			Number            int     `json:"number"`
			Touchdowns        int     `json:"touchdowns"`
			Yards             int     `json:"yards"`
			FirstName         string  `json:"first_name"`
			LastName          string  `json:"last_name"`
			ImageUrl          string  `json:"image_url"`
			Position          string  `json:"position"`
			ShortPosition     string  `json:"short_position"`
			YardsPerGame      float64 `json:"yards_per_game"`
			YardsPerReception float64 `json:"yards_per_reception"`
			Targets           int     `json:"targets"`
			GamesStarted      int     `json:"games_started"`
			ImageUrls         struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"receiving_leader"`
		TacklesLeader struct {
			Id                      int     `json:"id"`
			Interceptions           int     `json:"interceptions"`
			Tackles                 int     `json:"tackles"`
			Sacks                   float64 `json:"sacks"`
			Number                  int     `json:"number"`
			ForcedFumbles           int     `json:"forced_fumbles"`
			FirstName               string  `json:"first_name"`
			LastName                string  `json:"last_name"`
			ImageUrl                string  `json:"image_url"`
			Position                string  `json:"position"`
			ShortPosition           string  `json:"short_position"`
			Assists                 int     `json:"assists"`
			GamesStarted            int     `json:"games_started"`
			TotalTackles            int     `json:"total_tackles"`
			InterceptionsYards      int     `json:"interceptions_yards"`
			InterceptionsTouchdowns int     `json:"interceptions_touchdowns"`
			ImageUrls               struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"tackles_leader"`
		SacksLeader struct {
			Id                      int    `json:"id"`
			Interceptions           int    `json:"interceptions"`
			Tackles                 int    `json:"tackles"`
			Sacks                   int    `json:"sacks"`
			Number                  int    `json:"number"`
			ForcedFumbles           int    `json:"forced_fumbles"`
			FirstName               string `json:"first_name"`
			LastName                string `json:"last_name"`
			ImageUrl                string `json:"image_url"`
			Position                string `json:"position"`
			ShortPosition           string `json:"short_position"`
			Assists                 int    `json:"assists"`
			GamesStarted            int    `json:"games_started"`
			TotalTackles            int    `json:"total_tackles"`
			InterceptionsYards      int    `json:"interceptions_yards"`
			InterceptionsTouchdowns int    `json:"interceptions_touchdowns"`
			ImageUrls               struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"sacks_leader"`
		InterceptionsLeader struct {
			Id                      int    `json:"id"`
			Interceptions           int    `json:"interceptions"`
			Tackles                 int    `json:"tackles"`
			Sacks                   int    `json:"sacks"`
			Number                  int    `json:"number"`
			ForcedFumbles           int    `json:"forced_fumbles"`
			FirstName               string `json:"first_name"`
			LastName                string `json:"last_name"`
			ImageUrl                string `json:"image_url"`
			Position                string `json:"position"`
			ShortPosition           string `json:"short_position"`
			Assists                 int    `json:"assists"`
			GamesStarted            int    `json:"games_started"`
			TotalTackles            int    `json:"total_tackles"`
			InterceptionsYards      int    `json:"interceptions_yards"`
			InterceptionsTouchdowns int    `json:"interceptions_touchdowns"`
			ImageUrls               struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"interceptions_leader"`
		ForcedFumblesLeader struct {
			Id                      int     `json:"id"`
			Interceptions           int     `json:"interceptions"`
			Tackles                 int     `json:"tackles"`
			Sacks                   float64 `json:"sacks"`
			Number                  int     `json:"number"`
			ForcedFumbles           int     `json:"forced_fumbles"`
			FirstName               string  `json:"first_name"`
			LastName                string  `json:"last_name"`
			ImageUrl                string  `json:"image_url"`
			Position                string  `json:"position"`
			ShortPosition           string  `json:"short_position"`
			Assists                 int     `json:"assists"`
			GamesStarted            int     `json:"games_started"`
			TotalTackles            int     `json:"total_tackles"`
			InterceptionsYards      int     `json:"interceptions_yards"`
			InterceptionsTouchdowns int     `json:"interceptions_touchdowns"`
			ImageUrls               struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"forced_fumbles_leader"`
		FieldGoalsMadeLeader struct {
			Id                   int     `json:"id"`
			Number               int     `json:"number"`
			FirstName            string  `json:"first_name"`
			LastName             string  `json:"last_name"`
			ImageUrl             string  `json:"image_url"`
			Position             string  `json:"position"`
			ShortPosition        string  `json:"short_position"`
			FieldGoalsAttempted  int     `json:"field_goals_attempted"`
			FieldGoalsMade       int     `json:"field_goals_made"`
			FieldGoalsPercentage float64 `json:"field_goals_percentage"`
			ExtraPointsMade      int     `json:"extra_points_made"`
			GamesStarted         int     `json:"games_started"`
			ImageUrls            struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"field_goals_made_leader"`
	} `json:"season_leaders"`
	Players struct {
		Passing []struct {
			Id                   int    `json:"id"`
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
			ImageUrls            struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"passing"`
		Rushing []struct {
			Id              int    `json:"id"`
			RushesAttempted int    `json:"rushes_attempted"`
			LongestRush     int    `json:"longest_rush"`
			Number          int    `json:"number"`
			Average         string `json:"average"`
			Touchdowns      int    `json:"touchdowns"`
			Yards           int    `json:"yards"`
			FirstName       string `json:"first_name"`
			LastName        string `json:"last_name"`
			ImageUrl        string `json:"image_url"`
			Position        string `json:"position"`
			ImageUrls       struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"rushing"`
		Receiving []struct {
			Id               int    `json:"id"`
			Receptions       int    `json:"receptions"`
			LongestReceiving int    `json:"longest_receiving"`
			Number           int    `json:"number"`
			Average          string `json:"average"`
			Touchdowns       int    `json:"touchdowns"`
			Yards            int    `json:"yards"`
			FirstName        string `json:"first_name"`
			LastName         string `json:"last_name"`
			ImageUrl         string `json:"image_url"`
			Position         string `json:"position"`
			Targets          int    `json:"targets"`
			ImageUrls        struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"receiving"`
		Interceptions []struct {
			Id         int    `json:"id"`
			Number     int    `json:"number"`
			Touchdowns int    `json:"touchdowns"`
			Yards      int    `json:"yards"`
			FirstName  string `json:"first_name"`
			LastName   string `json:"last_name"`
			ImageUrl   string `json:"image_url"`
			Position   string `json:"position"`
			ImageUrls  struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"interceptions"`
		PuntReturns []struct {
			Id                int    `json:"id"`
			Returns           int    `json:"returns"`
			LongestPuntReturn int    `json:"longest_punt_return"`
			Number            int    `json:"number"`
			Average           string `json:"average"`
			Touchdowns        int    `json:"touchdowns"`
			Yards             int    `json:"yards"`
			FirstName         string `json:"first_name"`
			LastName          string `json:"last_name"`
			ImageUrl          string `json:"image_url"`
			Position          string `json:"position"`
			ImageUrls         struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"punt_returns"`
		Kickings []struct {
			Id               int    `json:"id"`
			Made             int    `json:"made"`
			Attempted        int    `json:"attempted"`
			Percentage       string `json:"percentage"`
			LongestKick      int    `json:"longest_kick"`
			Number           int    `json:"number"`
			FirstName        string `json:"first_name"`
			LastName         string `json:"last_name"`
			ImageUrl         string `json:"image_url"`
			Position         string `json:"position"`
			ExtraPointsMade  int    `json:"extra_points_made"`
			FieldGoalsMissed int    `json:"field_goals_missed"`
			ImageUrls        struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
			ExtraPointsAttempted int `json:"extra_points_attempted"`
		} `json:"kickings"`
		Defense []struct {
			Id            int     `json:"id"`
			Interceptions int     `json:"interceptions"`
			Tackles       int     `json:"tackles"`
			Sacks         float64 `json:"sacks"`
			SackYards     int     `json:"sack_yards"`
			Number        int     `json:"number"`
			FirstName     string  `json:"first_name"`
			LastName      string  `json:"last_name"`
			ImageUrl      string  `json:"image_url"`
			Position      string  `json:"position"`
			ImageUrls     struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
			PassesDefensed int `json:"passes_defensed"`
			SpTackles      int `json:"sp_tackles"`
		} `json:"defense"`
		Fumbles []struct {
			Id        int    `json:"id"`
			Number    int    `json:"number"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			ImageUrl  string `json:"image_url"`
			Position  string `json:"position"`
			ImageUrls struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"fumbles"`
	} `json:"players"`
	Color         string `json:"color"`
	TimeoutsLeft  int    `json:"timeouts_left"`
	ImageUrl25    string `json:"image_url_25"`
	ImageUrl59    string `json:"image_url_59"`
	TeamStandings struct {
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
		StreakGames            int     `json:"streak_games"`
		Streak                 string  `json:"streak"`
	} `json:"team_standings"`
}

type DetailsMobile struct {
	Id              int    `json:"id"`
	SrGameUuid      string `json:"sr_game_uuid"`
	LeagueShortName string `json:"league_short_name"`
	Quarter         int    `json:"quarter"`
	IsActive        bool   `json:"is_active"`
	Timestamp       int    `json:"timestamp"`
	Overtime        int    `json:"overtime"`
	Status          string `json:"status"`
	Clock           string `json:"clock"`
}

type DetailsWeb struct {
	Id               int    `json:"id"`
	LeagueShortName  string `json:"league_short_name"`
	Quarter          int    `json:"quarter"`
	IsActive         bool   `json:"is_active"`
	Clock            string `json:"clock"`
	Type             string `json:"type"`
	Timestamp        int    `json:"timestamp"`
	Status           string `json:"status"`
	BallLocation     string `json:"ball_location"`
	TeamPossessionId int    `json:"team_possession_id"`
	Down             int    `json:"down"`
	Distance         int    `json:"distance"`
	Location         string `json:"location"`
	LocationImageUrl string `json:"location_image_url"`
	Attendance       int    `json:"attendance"`
	GameName         string `json:"game_name"`
	Tbd              bool   `json:"tbd"`
	LocationImageMed string `json:"location_image_med"`
	LocationImageSml string `json:"location_image_sml"`
	Story            struct {
		Type     string `json:"type"`
		Headline string `json:"headline"`
		Content  string `json:"content"`
	} `json:"story"`
	Week        int    `json:"week"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Last9Events []struct {
		Time          string `json:"time"`
		Description   string `json:"description"`
		Quarter       int    `json:"quarter"`
		Down          int    `json:"down"`
		Distance      int    `json:"distance"`
		Yardline      string `json:"yardline"`
		EndYardline   string `json:"end_yardline"`
		Possession    string `json:"possession"`
		EndPossession string `json:"end_possession"`
		PlayType      string `json:"play_type"`
		PlayTypeId    int    `json:"play_type_id"`
		Id            int    `json:"id"`
		Players       []struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			ImageUrls struct {
				Lg string `json:"lg"`
				Md string `json:"md"`
				Sm string `json:"sm"`
				Xs string `json:"xs"`
			} `json:"image_urls"`
		} `json:"players"`
	} `json:"last_9_events"`
	Timeout       string `json:"timeout"`
	Last5Meetings []struct {
		Date              string `json:"date"`
		Location          string `json:"location"`
		GameId            int    `json:"game_id"`
		HomeTeam          int    `json:"home_team"`
		VisitingTeam      int    `json:"visiting_team"`
		VisitingTeamScore int    `json:"visiting_team_score"`
		HomeTeamScore     int    `json:"home_team_score"`
	} `json:"last_5_meetings"`
	State    string `json:"state"`
	GameType struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Detail string `json:"detail"`
	} `json:"game_type"`
	OpeningLine struct {
		FavId         int    `json:"fav_id"`
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
	} `json:"opening_line"`
	CurrentLine struct {
		FavId         int    `json:"fav_id"`
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
	} `json:"current_line"`
}
