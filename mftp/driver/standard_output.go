package driver

import (
	"bytes"
	"fmt"
	"log"
)

type StandardOutputDriver struct {
	buffer         *bytes.Buffer
	loggerInstance *log.Logger
}

func (driver *StandardOutputDriver) Initialize() {
	driver.buffer = &bytes.Buffer{}
	driver.loggerInstance = log.New(driver.buffer, "", log.LstdFlags)
}

func (driver *StandardOutputDriver) Write() error {
	fmt.Printf("%s\n", driver.buffer)
	driver.buffer.Reset()

	return nil
}

func (driver *StandardOutputDriver) Log(values ...any) {
	driver.loggerInstance.Print(values...)

	driver.Write()
}

func (driver *StandardOutputDriver) IsInitialized() bool {
	return driver.buffer != nil && driver.loggerInstance != nil
}
