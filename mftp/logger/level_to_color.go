package logger

import (
	"fmt"
	"mihaimiuta/mftp/ansi"
	"mihaimiuta/mftp/types"
)

func logLevelToColor(level types.LogLevel) ansi.Color {
	switch level {
	case types.Debug:
		return ansi.BrightCyan
	case types.Error:
		return ansi.BrightRed
	case types.Info:
		return ansi.BrightWhite

	default:
		panic(fmt.Sprintf("unexpected types.LogLevel: %#v", level))
	}
}
