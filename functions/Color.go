package handler


var clr = struct {
	reset   string
	red     string
	green   string
	yellow  string
	blue    string
	magenta string
	cyan    string
	gray    string
	white   string
	orange  string 
}{
	reset:   "\033[0m",
	red:     "\033[31m",
	green:   "\033[32m",
	yellow:  "\033[33m",
	blue:    "\033[34m", 
	magenta: "\033[35m",
	cyan:    "\033[36m",
	gray:    "\033[37m",
	white:   "\033[97m",
	orange:	 "\033[38;5;208m",
}

func Color(str []string, color string) []string {
	cv := ""
	switch color {
	case "red":
		cv = clr.red
	case "green":
		cv = clr.green
	case "yellow":
		cv = clr.yellow
	case "blue":
		cv = clr.blue
	case "magenta":
		cv = clr.magenta
	case "cyan":
		cv = clr.cyan
	case "gray":
		cv = clr.gray
	case "white":
		cv = clr.white
	case "orange":
		cv = clr.orange
	default:
		cv = clr.reset
	}

	for i := 0; i < len(str); i++ {
		str[i] = cv + str[i] + clr.reset
	}

	return str
}
