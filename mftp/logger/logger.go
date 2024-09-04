package logger

import (
	"context"
	"fmt"
	"time"

	"mihaimiuta/mftp/ansi"
	"mihaimiuta/mftp/driver"
	"mihaimiuta/mftp/types"

	"github.com/google/uuid"
)

type Logger struct {
	currentContext context.Context
	logs           []types.Log
	driverInstance *driver.Driver
}

func (logger *Logger) Initialize(parentContext context.Context, driverInstance driver.Driver) {
	logger.currentContext = context.WithValue(parentContext, "trace_id", uuid.New().String())
	logger.logs = []types.Log{}
	logger.driverInstance = &driverInstance
}

func (logger *Logger) WithContext(newParentContext context.Context) {
	if !logger.IsInitialized() {
		panic("Tried to use uninitialized logger")
	}

	logger.currentContext = context.WithValue(newParentContext, "trace_id", uuid.New().String())
}

func (logger *Logger) Log(level types.LogLevel, message string) {
	if !logger.IsInitialized() {
		panic("Tried to use uninitialized logger")
	}

	logger.logs = append(logger.logs, types.Log{
		Timestamp: time.Now(),
		TraceId:   logger.currentContext.Value("trace_id").(string),
		Level:     level,
		Message:   message,
	})

	if len(logger.logs) >= (*logger.driverInstance).GetBufferSize() {
		shouldUseColors := (*logger.driverInstance).SupportsANSIColors()

		(*logger.driverInstance).Write(logger.logs, func(log types.Log) string {
			message := fmt.Sprintf("[Level=%s][Timestamp=%s][TraceID=%s]: %s\n", log.Level, log.Timestamp.String(), log.TraceId, log.Message)

			if shouldUseColors {
				return ansi.ColorMessage(message, logLevelToColor(log.Level))
			}

			return message
		})
	}
}

func (logger *Logger) IsInitialized() bool {
	contextInitialized := logger.currentContext != nil
	driverInitialized := logger.driverInstance != nil

	return contextInitialized && driverInitialized
}
