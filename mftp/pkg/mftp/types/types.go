package types

import "time"

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Error LogLevel = "error"
)

type Log struct {
	Timestamp  time.Time
	TraceId    string
	Level      LogLevel
	Message    string
	Attributes map[string]string
}

type Config struct {
	LogFormat string `json:"log_format"`
	Driver    *struct {
		OverrideBufferSize int `json:"override_buffer_size"`
	} `json:"driver"`
}
