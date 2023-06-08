package soccer

type GameDataWeb struct {
	Details      DetailsWeb `json:"details"`
	VisitingTeam TeamWeb    `json:"visiting_team"`
	HomeTeam     TeamWeb    `json:"home_team"`
	Periods      []struct {
		VisitingTeamScore int `json:"visiting_team_score"`
		HomeTeamScore     int `json:"home_team_score"`
		Goals             []struct {
			PlayerID       int    `json:"player_id"`
			TeamID         int    `json:"team_id"`
			FirstName      string `json:"first_name"`
			LastName       string `json:"last_name"`
			Minutes        int    `json:"minutes"`
			Seconds        int    `json:"seconds"`
			AdditionalMins int    `json:"additional_mins"`
			PenaltyKick    bool   `json:"penalty_kick"`
			OwnGoal        bool   `json:"own_goal"`
			ShootoutGoal   bool   `json:"shootout_goal"`
		} `json:"goals"`
		Cards []struct {
			Type           string `json:"type"`
			TeamID         int    `json:"team_id"`
			PlayerID       int    `json:"player_id"`
			FirstName      string `json:"first_name"`
			LastName       string `json:"last_name"`
			Minutes        int    `json:"minutes"`
			Seconds        int    `json:"seconds"`
			AdditionalMins int    `json:"additional_mins"`
		} `json:"cards"`
		Substitutions []struct {
			TeamID         int `json:"team_id"`
			Minutes        int `json:"minutes"`
			Seconds        int `json:"seconds"`
			AdditionalMins int `json:"additional_mins"`
			In             struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			} `json:"in"`
			Out struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			} `json:"out"`
		} `json:"substitutions"`
	} `json:"periods"`
}

type DetailsWeb struct {
	LeagueShortName    string `json:"league_short_name"`
	ID                 int    `json:"id"`
	SrGameUuid         string `json:"sr_game_uuid"`
	Period             int    `json:"period"`
	IsActive           bool   `json:"is_active"`
	Clock              string `json:"clock"`          // "match_clock":"45:00"
	Time               int    `json:"time"`           // "match_time":45,
	StoppageClock      string `json:"stoppage_clock"` // "stoppage_time_clock":"45:00"
	StoppageTime       int    `json:"stoppage_time"`  // "stoppage_time":45,
	Type               string `json:"type"`
	Timestamp          int    `json:"timestamp"`
	Status             string `json:"status"`
	IsHalfTime         bool   `json:"ishalftime"`
	HasAggregateScores bool   `json:"hasAggregateScores"`
	Location           string `json:"location"`
	LocationImageURL   string `json:"location_image_url"`
	ExtraTime          bool   `json:"extra_time"`
	Overtime           int    `json:"overtime"`
	PenaltyKicks       bool   `json:"penalty_kicks"`
	Broadcast          []struct {
		Name                  string `json:"name"`
		ImageURL              string `json:"image_url"`
		BlackImageURL         string `json:"black_image_url"`
		HrImageURL            string `json:"hr_image_url"`
		URLSportsnetNow       string `json:"url_sportsnet_now"`
		URLBroadcasterLogoSvg string `json:"url_broadcaster_logo_svg"`
		ChannelID             string `json:"channel_id"`
		ChannelWatchliveID    string `json:"channel_watchlive_id"`
		ChannelNeulionID      int    `json:"channel_neulion_id,omitempty"`
		EndTime               int    `json:"end_time"`
		ChannelType           string `json:"channel_type,omitempty"`
		NeulionID             string `json:"neulion_id"`
	} `json:"broadcast"`
	LocationDetails struct {
		Venue   string `json:"venue"`
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"location_details"`
	Tbd bool `json:"tbd"`
}

type TeamWeb struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	City           string `json:"city"`
	ImageURL       string `json:"image_url"`
	Color          string `json:"color"`
	Score          int    `json:"score"`
	PenaltyScore   int    `json:"penalty_score"`
	AggregateScore *int   `json:"aggregate_score"` // Pointer so during unmarshalling it can be null when the json is missing the field, will handle the GO defaulting to zero values case
	ImageURL90     string `json:"image_url_90"`
	Division       struct {
	} `json:"division"`
	Conference struct {
	} `json:"conference"`
	Linescore struct {
		Score int `json:"score"`
		Shots int `json:"shots"`
		Halfs []struct {
			Half  int `json:"half"`
			Score int `json:"score"`
			Shots int `json:"shots"`
		} `json:"halfs"`
	} `json:"linescore"`
	GameStats struct {
		PossessionPercentage int `json:"possession_percentage"`
		Goals                int `json:"goals"`
		Assists              int `json:"assists"`
		Shots                int `json:"shots"`
		ShotsOnGoal          int `json:"shots_on_goal"`
		CornerKicks          int `json:"corner_kicks"`
		Fouls                int `json:"fouls"`
		Crosses              int `json:"crosses"`
		YellowCards          int `json:"yellow_cards"`
		RedCards             int `json:"red_cards"`
		Offsides             int `json:"offsides"`
		Saves                int `json:"saves"`
	} `json:"game_stats"`
	Goalkeepers []struct {
		ID            int    `json:"id"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		ImageURL      string `json:"image_url"`
		Position      string `json:"position"`
		ShortPosition string `json:"short_position"`
		Number        int    `json:"number"`
		GoalsAgainst  int    `json:"goals_against"`
		Saves         int    `json:"saves"`
		ShotsAgainst  int    `json:"shots_against"`
		ShotsOnGoal   int    `json:"shots_on_goal"`
	} `json:"goalkeepers"`
	Defenders []struct {
		ID            int    `json:"id"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		ImageURL      string `json:"image_url"`
		Position      string `json:"position"`
		ShortPosition string `json:"short_position"`
		Number        int    `json:"number"`
		Goals         int    `json:"goals"`
		Assists       int    `json:"assists"`
		Shots         int    `json:"shots"`
		ShotsOnGoal   int    `json:"shots_on_goal"`
		Crosses       int    `json:"crosses"`
		Tackles       int    `json:"tackles"`
		Blocks        int    `json:"blocks"`
		Touches       int    `json:"touches"`
		Passes        int    `json:"passes"`
		Fouls         int    `json:"fouls"`
		Offsides      int    `json:"offsides"`
		CornerKicks   int    `json:"corner_kicks"`
		YellowCards   int    `json:"yellow_cards"`
		RedCards      int    `json:"red_cards"`
		Starter       bool   `json:"starter"`
		OwnGoals      int    `json:"own_goals"`
	} `json:"defenders"`
	Midfielders []struct {
		ID            int    `json:"id"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		ImageURL      string `json:"image_url"`
		Position      string `json:"position"`
		ShortPosition string `json:"short_position"`
		Number        int    `json:"number"`
		Goals         int    `json:"goals"`
		Assists       int    `json:"assists"`
		Shots         int    `json:"shots"`
		ShotsOnGoal   int    `json:"shots_on_goal"`
		Crosses       int    `json:"crosses"`
		Tackles       int    `json:"tackles"`
		Blocks        int    `json:"blocks"`
		Touches       int    `json:"touches"`
		Passes        int    `json:"passes"`
		Fouls         int    `json:"fouls"`
		Offsides      int    `json:"offsides"`
		CornerKicks   int    `json:"corner_kicks"`
		YellowCards   int    `json:"yellow_cards"`
		RedCards      int    `json:"red_cards"`
		Starter       bool   `json:"starter"`
		OwnGoals      int    `json:"own_goals"`
	} `json:"midfielders"`
	Forwards []struct {
		ID            int    `json:"id"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		ImageURL      string `json:"image_url"`
		Position      string `json:"position"`
		ShortPosition string `json:"short_position"`
		Number        int    `json:"number"`
		Goals         int    `json:"goals"`
		Assists       int    `json:"assists"`
		Shots         int    `json:"shots"`
		ShotsOnGoal   int    `json:"shots_on_goal"`
		Crosses       int    `json:"crosses"`
		Tackles       int    `json:"tackles"`
		Blocks        int    `json:"blocks"`
		Touches       int    `json:"touches"`
		Passes        int    `json:"passes"`
		Fouls         int    `json:"fouls"`
		Offsides      int    `json:"offsides"`
		CornerKicks   int    `json:"corner_kicks"`
		YellowCards   int    `json:"yellow_cards"`
		RedCards      int    `json:"red_cards"`
		Starter       bool   `json:"starter"`
		OwnGoals      int    `json:"own_goals"`
	} `json:"forwards"`
	SeasonStats struct {
		Wins   int `json:"wins"`
		Ties   int `json:"ties"`
		Losses int `json:"losses"`
		Points int `json:"points"`
	} `json:"season_stats"`
	LeagueRank int `json:"league_rank"`
	TeamStats  struct {
		GoalsPerGame            string `json:"goals_per_game"`
		GoalsPerGameRank        int    `json:"goals_per_game_rank"`
		GoalsAgainstPerGame     string `json:"goals_against_per_game"`
		GoalsAgainstPerGameRank int    `json:"goals_against_per_game_rank"`
		GoalDifferential        int    `json:"goal_differential"`
		GoalDifferentialRank    int    `json:"goal_differential_rank"`
		Shutouts                int    `json:"shutouts"`
		ShutoutsRank            int    `json:"shutouts_rank"`
	} `json:"team_stats"`
}
