package driver

import (
	"fmt"
	"mihaimiuta/mftp/pkg/mftp/types"
)

type StandardOutputDriver struct {
}

func (driver *StandardOutputDriver) Write(logs []types.Log, encodeLog LogEncoder) error {
	for _, log := range logs {
		fmt.Print(encodeLog(log))
	}

	return nil
}

func (driver *StandardOutputDriver) GetBufferSize() int {
	return 1
}

func (driver *StandardOutputDriver) SupportsANSIColors() bool {
	return true
}
