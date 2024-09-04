package driver

import (
	"mihaimiuta/mftp/types"
)

type LogEncoder func(types.Log) string

type Driver interface {
	Write(logs []types.Log, encodeLog LogEncoder) error
	SupportsANSIColors() bool
	GetBufferSize() int
}
