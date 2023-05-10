package constants

var LEAGUE_BY_SPORT = map[string][]string{
	"hockey":     {"OLY_MHK", "OLY_WHK", "NHL", "OHL", "WHL", "CHL", "WCOH", "WHC", "WWHC", "AHL", "WJC", "QMJHL"},
	"basketball": {"NBA", "WNBA", "OLY_WBK", "OLY_MBK"},
	"baseball":   {"MLB", "WBC"},
	"football":   {"NFL", "CFL", "NCAAF", "NCAAB"},
	"soccer":     {"ENG_FA_CUP", "OLY_WSOC", "OLY_MSOC", "EPL", "BPL", "MLS", "BUND", "CHLG", "NATL", "WWC", "FRAN", "SERI", "LIGA", "EURO"},
	"juniors":    {"WJHC", "OHL", "WHL", "QMJHL", "CHL"},
	"autoracing": {"NASCAR", "NASCAR_2", "FORM1", "IRL"},
	"tennis":     {"ATP", "WTA"},
	"golf":       {"PGA", "LPGA"},
	"olympics":   {"OLY_MHK", "OLY_WHK", "OLY_WSOC", "OLY_MSOC", "OLY_WBK", "OLY_MBK"},
	"lacrosse":   {"NLL"},
}
