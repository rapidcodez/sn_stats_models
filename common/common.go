package common

type GameBroadcast struct {
	Name                  string `json:"name"`
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
