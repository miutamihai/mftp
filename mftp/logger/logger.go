package logger

import (
	"context"
	"time"

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

func (logger *Logger) Log(level types.LogLevel, message string, attributes map[string]string) error {
	if !logger.IsInitialized() {
		panic("Tried to use uninitialized logger")
	}

	log := types.Log{
		Timestamp:  time.Now(),
		TraceId:    logger.currentContext.Value("trace_id").(string),
		Level:      level,
		Message:    message,
		Attributes: attributes,
	}

	logger.logs = append(logger.logs, log)

	if len(logger.logs) >= (*logger.driverInstance).GetBufferSize() {
		shouldUseColors := (*logger.driverInstance).SupportsANSIColors()

		err := (*logger.driverInstance).Write(logger.logs, makeLogEncoder(shouldUseColors))

		if err != nil {
			return err
		}

	}

	return nil
}

func (logger *Logger) IsInitialized() bool {
	contextInitialized := logger.currentContext != nil
	driverInitialized := logger.driverInstance != nil

	return contextInitialized && driverInitialized
}
