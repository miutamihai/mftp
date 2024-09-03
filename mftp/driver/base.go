package driver

import (
	"context"
	"time"
)

type Driver interface {
	Write() error
	Initialize(context.Context)
	IsInitialized() bool
	Log(level LogLevel, message string)
}

type LogLevel string

const (
	Debug LogLevel = "debug"
	Error LogLevel = "error"
)

type Log struct {
	Timestamp time.Time
	TraceId   string
	Level     LogLevel
	Message   string
}

const log_limit = 1
