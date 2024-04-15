package format

type colors struct {
	Green  string
	Yellow string
	Blue   string
	Red    string
	Purple string
	Orange string
	Cyan   string
	Gray   string
	White  string
	Reset  string
}

// Colors Material Darker Theme
var Colors = colors{
	Green:  "\033[38;2;195;232;141m",
	Yellow: "\033[38;2;255;203;107m",
	Blue:   "\033[38;2;130;170;255m",
	Red:    "\033[38;2;240;113;120m",
	Purple: "\033[38;2;199;146;234m",
	Orange: "\033[38;2;247;140;108m",
	Cyan:   "\033[38;2;137;221;255m",
	Gray:   "\033[38;2;97;97;97m",
	White:  "\033[38;2;238;255;255m",
	Reset:  "\033[0m",
}
