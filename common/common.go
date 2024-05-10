package common

type GameBroadcast struct {
	Name                  string `json:"name"`
	TMSID                 string `json:"tmsid"`
	EventID               string `json:"event_id"`
	SourceFeed            string `json:"source_feed"`
	ImageURL              string `json:"image_url"`
	BlackImageURL         string `json:"black_image_url"`
	HrImageURL            string `json:"hr_image_url"`
	URLSportsnetNow       string `json:"url_sportsnet_now"`
	URLBroadcasterLogoSvg string `json:"url_broadcaster_logo_svg"`
	ChannelID             string `json:"channel_id"`
	ChannelWatchliveID    string `json:"channel_watchlive_id"`
	ChannelNeulionID      int    `json:"channel_neulion_id"`
	EndTime               int    `json:"end_time"`
	ChannelType           string `json:"channel_type"`
	NeulionID             string `json:"neulion_id"`
	URLSportsnetNowEvent  string `json:"url_sportsnet_now_event"`
	SportsnetNowDeeplink  string `json:"sportsnet_now_deeplink"`
	IsRegional            bool   `json:"is_regional"`
	FreeBroadcast         bool   `json:"free_broadcast"`
}

type RecentGames struct {
	Date              string `json:"date"`
	Location          string `json:"location"`
	GameID            string `json:"game_id,omitempty"`
	ID                string `json:"id"`
	SrGameId          string `json:"sr_game_id"`
	VisitingTeamScore int    `json:"visiting_team_score"`
	HomeTeamScore     int    `json:"home_team_score"`
	VisitingTeam      string `json:"visiting_team,omitempty"`
	VisitingTeamID    string `json:"visiting_team_id,omitempty"`
	HomeTeam          string `json:"home_team,omitempty"`
	HomeTeamID        string `json:"home_team_id,omitempty"`
}

type Story struct {
	Type     string `json:"type"`
	Headline string `json:"headline"`
	Content  string `json:"content"`
}

type Injuries struct {
	PlayerID             string `json:"player_id"`
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

type ImageUrl struct {
	Lg string `json:"lg"`
	Md string `json:"md"`
	Sm string `json:"sm"`
	Xs string `json:"xs"`
}
