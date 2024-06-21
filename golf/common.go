package sportradar

import "time"

type TournamentSummary struct {
	CourseTimezone string                     `json:"course_timezone"`
	Coverage       string                     `json:"coverage"`
	Currency       string                     `json:"currency"`
	EndDate        string                     `json:"end_date"`
	EventType      string                     `json:"event_type"`
	Field          []TournamentSummaryField   `json:"field"`
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	ParentID       string                     `json:"parent_id"`
	Points         int                        `json:"points"`
	Purse          int                        `json:"purse"`
	Rounds         []TournamentSummaryRounds  `json:"rounds"`
	Seasons        []TournamentSummarySeasons `json:"seasons"`
	StartDate      string                     `json:"start_date"`
	Status         string                     `json:"status"`
	Venue          TournamentSummaryVenue     `json:"venue"`
	WinningShare   int                        `json:"winning_share"`
}
type TournamentSummaryField struct {
	AbbrName  string `json:"abbr_name"`
	Amateur   bool   `json:"amateur,omitempty"`
	Country   string `json:"country"`
	FirstName string `json:"first_name"`
	ID        string `json:"id"`
	LastName  string `json:"last_name"`
}
type TournamentSummaryBroadcasts struct {
	EndAt    time.Time `json:"end_at"`
	Radio    string    `json:"radio,omitempty"`
	StartAt  time.Time `json:"start_at"`
	Network  string    `json:"network,omitempty"`
	Internet string    `json:"internet,omitempty"`
}
type TournamentSummaryWind struct {
	Direction string `json:"direction"`
	Speed     int    `json:"speed"`
}
type TournamentSummaryWeather struct {
	Condition string                `json:"condition"`
	Temp      int                   `json:"temp"`
	Wind      TournamentSummaryWind `json:"wind"`
}
type TournamentSummaryRounds struct {
	Broadcasts []TournamentSummaryBroadcasts `json:"broadcasts"`
	ID         string                        `json:"id"`
	Number     int                           `json:"number"`
	Status     string                        `json:"status"`
	Weather    TournamentSummaryWeather      `json:"weather"`
}
type TournamentSummaryTour struct {
	Alias string `json:"alias"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}
type TournamentSummarySeasons struct {
	ID   string                `json:"id"`
	Tour TournamentSummaryTour `json:"tour"`
	Year int                   `json:"year"`
}
type TournamentSummaryHoles struct {
	Number  int `json:"number"`
	Par     int `json:"par"`
	Yardage int `json:"yardage"`
}
type TournamentSummaryCourses struct {
	Holes   []TournamentSummaryHoles `json:"holes"`
	ID      string                   `json:"id"`
	Name    string                   `json:"name"`
	Par     int                      `json:"par"`
	Yardage int                      `json:"yardage"`
}
type TournamentSummaryVenue struct {
	City      string                     `json:"city"`
	Country   string                     `json:"country"`
	Courses   []TournamentSummaryCourses `json:"courses"`
	ID        string                     `json:"id"`
	Latitude  string                     `json:"latitude"`
	Longitude string                     `json:"longitude"`
	Name      string                     `json:"name"`
	State     string                     `json:"state"`
	Zipcode   string                     `json:"zipcode"`
}
