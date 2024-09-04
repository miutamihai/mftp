package logger

import (
	"fmt"
	"mihaimiuta/mftp/internal/mftp/ansi"
	"mihaimiuta/mftp/pkg/mftp/types"
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
