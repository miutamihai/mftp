package mftp

import (
	"context"
	"errors"
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
	config         types.Config
}

type InitializationInput struct {
	ParentContext  context.Context
	DriverInstance driver.Driver
	ConfigPath     string
}

func (loggerInstance *Logger) Initialize(input InitializationInput) error {
	config, err := readConfig(input.ConfigPath)

	if err != nil {
		return err
	}

	loggerInstance.currentContext = context.WithValue(input.ParentContext, "trace_id", uuid.New().String())
	loggerInstance.logs = []types.Log{}
	loggerInstance.driverInstance = &input.DriverInstance
	loggerInstance.config = config

	return nil
}

func (loggerInstance *Logger) WithContext(newParentContext context.Context) error {
	if !loggerInstance.IsInitialized() {
		return errors.New("Tried to use uninitialized logger")
	}

	loggerInstance.currentContext = context.WithValue(newParentContext, "trace_id", uuid.New().String())

	return nil
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
		return errors.New("Tried to use uninitialized logger")
	}

	log := types.Log{
		Timestamp:  time.Now(),
		TraceId:    loggerInstance.currentContext.Value("trace_id").(string),
		Level:      level,
		Message:    message,
		Attributes: attributes,
	}

	loggerInstance.logs = append(loggerInstance.logs, log)

	bufferSize := (*loggerInstance.driverInstance).GetBufferSize()

	if loggerInstance.config.Driver != nil {
		bufferSize = loggerInstance.config.Driver.OverrideBufferSize
	}

	if len(loggerInstance.logs) >= bufferSize {
		shouldUseColors := (*loggerInstance.driverInstance).SupportsANSIColors()

		err := (*loggerInstance.driverInstance).Write(loggerInstance.logs, logger.MakeLogEncoder(shouldUseColors, loggerInstance.config.LogFormat))

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
