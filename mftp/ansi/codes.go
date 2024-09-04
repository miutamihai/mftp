package ansi

type Color string

const (
	BrightCyan  Color = "[96m"
	BrightRed   Color = "[91m"
	BrightWhite Color = "[97m"
)

const (
	escape = "\033"
	reset  = "[0m"
)
