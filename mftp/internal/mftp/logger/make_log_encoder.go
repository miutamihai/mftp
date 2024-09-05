package logger

import (
	"fmt"
	"maps"
	"mihaimiuta/mftp/internal/mftp/ansi"
	"mihaimiuta/mftp/pkg/mftp/types"
	"strings"
)

const defaultFormat = "LEVEL,TIMESTAMP,TRACE_ID,ATTRIBUTES"

func encodeWithFormat(format string, log types.Log) string {
	var builder strings.Builder

	for _, word := range strings.Split(format, ",") {
		switch strings.TrimSpace(word) {
		case "LEVEL":
			fmt.Fprintf(&builder, "[Level=%s]", log.Level)
		case "TIMESTAMP":
			fmt.Fprintf(&builder, "[Timestamp=%s]", log.Timestamp.String())
		case "TRACE_ID":
			fmt.Fprintf(&builder, "[TraceID=%s]", log.TraceId)
		case "ATTRIBUTES":
			if log.Attributes != nil {
				fmt.Fprint(&builder, "[Attributes={")

				for key, value := range maps.All(log.Attributes) {
					fmt.Fprintf(&builder, "[%s=%s]", key, value)
				}

				fmt.Fprint(&builder, "}]")
			}
		default:
			panic(fmt.Sprintf("unexpected format word: %#v", word))
		}
	}

	fmt.Fprintf(&builder, " %s\n", log.Message)

	return builder.String()
}

func MakeLogEncoder(shouldUseColors bool, format string) func(types.Log) string {
	return func(log types.Log) string {
		formatToUse := defaultFormat

		if format != "" {
			formatToUse = format
		}

		message := encodeWithFormat(formatToUse, log)

		if shouldUseColors {
			return ansi.ColorMessage(message, logLevelToColor(log.Level))
		}

		return message
	}
}
