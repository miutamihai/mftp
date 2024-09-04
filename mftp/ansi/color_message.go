package ansi

import (
	"fmt"
)

func ColorMessage(message string, color Color) string {
	return fmt.Sprintf("%s%s%s%s%s", escape, color, message, escape, reset)
}
