package types

import "time"

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Error LogLevel = "error"
)

type Log struct {
	Timestamp time.Time
	TraceId   string
	Level     LogLevel
	Message   string
}
