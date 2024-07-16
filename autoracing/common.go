package autoracing

type Race struct {
	ID                string   `json:"id"`
	Timestamp         int      `json:"timestamp"`
	Date              string   `json:"date"`
	DefendingChampion Driver   `json:"defending_champion"`
	Laps              int      `json:"laps"`
	LapsCompleted     int      `json:"laps_completed"`
	Location          string   `json:"location"`
	Miles             string   `json:"miles"`
	Name              string   `json:"name"`
	Nickname          string   `json:"nickname"`
	Purse             string   `json:"purse"`
	Results           []Result `json:"results"`
	Status            string   `json:"status"`
	Track             string   `json:"track"`
	UnderCaution      string   `json:"under_caution"`
	Winner            Driver   `json:"winner"`
}

type Result struct {
	Driver
	AverageSpeed     string `json:"average_speed"`
	Bonus            int    `json:"bonus"`
	CarMake          string `json:"car_make"`
	CarNumber        int    `json:"car_number"`
	DriverStatus     string `json:"driver_status"`
	LapsCompleted    int    `json:"laps_completed"`
	LapsLed          int    `json:"laps_led"`
	MoneyWon         string `json:"money_won"`
	OwnerPoints      int    `json:"owner_points"`
	PointsPenalized  int    `json:"points_penalized"`
	Rank             int    `json:"rank"`
	StartingPosition int    `json:"starting_position"`
}

type Driver struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	FlagURL   string `json:"flag_url"`
	LastName  string `json:"last_name"`
}
