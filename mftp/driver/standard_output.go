package driver

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type StandardOutputDriver struct {
	currentContext context.Context
	logs           []Log
}

func (driver *StandardOutputDriver) Initialize(parentContext context.Context) {
	driver.currentContext = context.WithValue(parentContext, "trace_id", uuid.New().String())
}

func (driver *StandardOutputDriver) Write() error {
	for _, log := range driver.logs {
		var colorCode int

		switch log.Level {
		case Debug:
			colorCode = 96
		case Error:
			colorCode = 91
		default:
			panic(fmt.Sprintf("unexpected driver.LogLevel: %#v", log.Level))
		}

		message := fmt.Sprintf("[Level=%s][Timestamp=%s][TraceID=%s]: %s\n", log.Level, log.Timestamp.String(), log.TraceId, log.Message)

		fmt.Printf("\033[%sm%s\x1b[0m", strconv.Itoa(colorCode), message)
	}

	driver.logs = []Log{}

	return nil
}

func (driver *StandardOutputDriver) Log(level LogLevel, message string) {
	driver.logs = append(driver.logs, Log{
		Timestamp: time.Now(),
		TraceId:   driver.currentContext.Value("trace_id").(string),
		Level:     level,
		Message:   message,
	})

	if len(driver.logs) >= log_limit {
		driver.Write()
	}
}

func (driver *StandardOutputDriver) IsInitialized() bool {
	return driver.currentContext != nil
}
