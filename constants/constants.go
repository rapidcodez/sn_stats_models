package constants

var LEAGUE_BY_SPORT = map[string][]string{
	"hockey":     {"OLY_MHK", "OLY_WHK", "NHL", "OHL", "WHL", "CHL", "WCOH", "WHC", "WWHC", "AHL", "WJC", "QMJHL"},
	"basketball": {"NBA", "WNBA", "OLY_WBK", "OLY_MBK", "NCAAB"},
	"baseball":   {"MLB", "WBC", "NCAAB"},
	"football":   {"NFL", "CFL", "NCAAF"},
	"soccer":     {"ENG_FA_CUP", "OLY_WSOC", "OLY_MSOC", "EPL", "BPL", "MLS", "BUND", "CHLG", "NATL", "WWC", "FRAN", "SERI", "LIGA", "EURO"},
	"juniors":    {"WJHC", "OHL", "WHL", "QMJHL", "CHL"},
	"autoracing": {"NASCAR", "NASCAR_2", "FORM1", "IRL"},
	"tennis":     {"ATP", "WTA"},
	"golf":       {"PGA", "LPGA"},
	"olympics":   {"OLY_MHK", "OLY_WHK", "OLY_WSOC", "OLY_MSOC", "OLY_WBK", "OLY_MBK"},
	"lacrosse":   {"NLL"},
}

var SEASON_TYPES = map[string]int{
	"PRE":            1,
	"REG":            2,
	"PST":            3,
	"playoffs":       3,
	"regular season": 2,
	"pre-season":     1,
	"pre season":     1,
	"qualification":  1,
	"SIM":            10,
	"PIT":            5,
}

var SEASON_TYPES_R = map[int]string{
	1:  "PRE",
	2:  "REG",
	3:  "PST",
	10: "SIM",
	5:  "PIT",
}
