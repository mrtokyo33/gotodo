package utils

const (
	Reset = "\033[0m"

	Red    = "\033[31m"
	Yellow = "\033[33m"

	Blue      = "\033[34m"
	LightBlue = "\033[94m"
	Cyan      = "\033[36m"
	DarkBlue  = "\033[38;5;25m"

	Bold = "\033[1m"
)

func Error(msg string) string {
	return Red + msg + Reset
}

func Warning(msg string) string {
	return Yellow + msg + Reset
}

func Primary(msg string) string {
	return LightBlue + msg + Reset
}

func Highlight(msg string) string {
	return Cyan + msg + Reset
}

func Title(msg string) string {
	return Bold + Blue + msg + Reset
}
