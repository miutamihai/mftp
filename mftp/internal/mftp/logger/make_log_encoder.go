package logger

import (
	"fmt"
	"maps"
	"mihaimiuta/mftp/internal/mftp/ansi"
	"mihaimiuta/mftp/pkg/mftp/types"
	"strings"
)

func MakeLogEncoder(shouldUseColors bool) func(types.Log) string {
	return func(log types.Log) string {
		var builder strings.Builder

		fmt.Fprintf(&builder, "[Level=%s]", log.Level)
		fmt.Fprintf(&builder, "[Timestamp=%s]", log.Timestamp.String())
		fmt.Fprintf(&builder, "[TraceID=%s]", log.TraceId)

		if log.Attributes != nil {
			fmt.Fprint(&builder, "[Attributes={")

			for key, value := range maps.All(log.Attributes) {
				fmt.Fprintf(&builder, "[%s=%s],", key, value)
			}

			fmt.Fprint(&builder, "}]")
		}
		fmt.Fprintf(&builder, " %s", log.Message)

		message := builder.String()

		if shouldUseColors {
			return ansi.ColorMessage(message, logLevelToColor(log.Level))
		}

		return message
	}
}
