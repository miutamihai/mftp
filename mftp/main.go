package main

import (
	"context"
	"mihaimiuta/mftp/driver"
	"mihaimiuta/mftp/logger"
	"mihaimiuta/mftp/types"
	"strconv"
	"time"
)

func doWork(loggerInstance logger.Logger) {
	if !loggerInstance.IsInitialized() {
		loggerInstance.Initialize(context.Background(), &driver.StandardOutputDriver{})
	}

	for number := range 10_000 {
		loggerInstance.Log(types.Info, "doing some work")

		if number%3 == 0 {
			loggerInstance.Log(types.Error, "something went wrong")
			loggerInstance.Log(types.Error, strconv.Itoa(number))
		} else {
			loggerInstance.Log(types.Debug, "something went right")
			loggerInstance.Log(types.Debug, strconv.Itoa(number))
		}

		time.Sleep(2 * time.Second)
	}
}

func main() {
	loggerInstance := logger.Logger{}

	loggerInstance.Initialize(context.Background(), &driver.StandardOutputDriver{})

	doWork(loggerInstance)
}
