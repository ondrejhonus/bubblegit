package global

type Colours struct {
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	White  string
}

type Styles struct {
	Reset     string
	Bold      string
	Underline string
	Strike    string
	Italic    string
}

func Colour() Colours {
	return Colours{
		Red:    "\033[31m",
		Green:  "\033[32m",
		Yellow: "\033[33m",
		Blue:   "\033[34m",
		Purple: "\033[35m",
		Cyan:   "\033[36m",
		White:  "\033[37m",
	}
}

func Style() Styles {
	return Styles{
		Reset:     "\033[0m",
		Bold:      "\033[1m",
		Underline: "\033[4m",
		Strike:    "\033[9m",
		Italic:    "\033[3m",
	}
}
