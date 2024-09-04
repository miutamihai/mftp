package mftp

import (
	"context"
	"mihaimiuta/mftp/internal/mftp/logger"
	"mihaimiuta/mftp/pkg/mftp/driver"
	"mihaimiuta/mftp/pkg/mftp/types"
	"time"

	"github.com/google/uuid"
)

type Logger struct {
	currentContext context.Context
	logs           []types.Log
	driverInstance *driver.Driver
}

func (loggerInstance *Logger) Initialize(parentContext context.Context, driverInstance driver.Driver) {
	loggerInstance.currentContext = context.WithValue(parentContext, "trace_id", uuid.New().String())
	loggerInstance.logs = []types.Log{}
	loggerInstance.driverInstance = &driverInstance
}

func (loggerInstance *Logger) WithContext(newParentContext context.Context) {
	if !loggerInstance.IsInitialized() {
		panic("Tried to use uninitialized logger")
	}

	loggerInstance.currentContext = context.WithValue(newParentContext, "trace_id", uuid.New().String())
}

func (loggerInstance *Logger) Info(message string, attributes map[string]string) error {
	return loggerInstance.Log(types.Info, message, attributes)
}

func (loggerInstance *Logger) Error(message string, attributes map[string]string) error {
	return loggerInstance.Log(types.Error, message, attributes)
}

func (loggerInstance *Logger) Debug(message string, attributes map[string]string) error {
	return loggerInstance.Log(types.Debug, message, attributes)
}

func (loggerInstance *Logger) Log(level types.LogLevel, message string, attributes map[string]string) error {
	if !loggerInstance.IsInitialized() {
		panic("Tried to use uninitialized logger")
	}

	log := types.Log{
		Timestamp:  time.Now(),
		TraceId:    loggerInstance.currentContext.Value("trace_id").(string),
		Level:      level,
		Message:    message,
		Attributes: attributes,
	}

	loggerInstance.logs = append(loggerInstance.logs, log)

	if len(loggerInstance.logs) >= (*loggerInstance.driverInstance).GetBufferSize() {
		shouldUseColors := (*loggerInstance.driverInstance).SupportsANSIColors()

		err := (*loggerInstance.driverInstance).Write(loggerInstance.logs, logger.MakeLogEncoder(shouldUseColors))

		if err != nil {
			return err
		}

	}

	return nil
}

func (loggerInstance *Logger) IsInitialized() bool {
	contextInitialized := loggerInstance.currentContext != nil
	driverInitialized := loggerInstance.driverInstance != nil

	return contextInitialized && driverInitialized
}
