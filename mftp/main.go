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
		err := loggerInstance.Log(types.Info, "doing some work")

		if err != nil {
			panic(err)
		}

		if number%3 == 0 {
			err := loggerInstance.Log(types.Error, "something went wrong")

			if err != nil {
				panic(err)
			}

			err2 := loggerInstance.Log(types.Error, strconv.Itoa(number))

			if err2 != nil {
				panic(err2)
			}
		} else {
			err := loggerInstance.Log(types.Debug, "something went right")

			if err != nil {
				panic(err)
			}

			err2 := loggerInstance.Log(types.Debug, strconv.Itoa(number))

			if err2 != nil {
				panic(err2)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

func main() {
	loggerInstance := logger.Logger{}

	go func() {
		loggerInstance.Initialize(context.Background(), &driver.StandardOutputDriver{})

		doWork(loggerInstance)
	}()

	loggerInstance.Initialize(context.Background(), &driver.TextFileDriver{
		FilePath: "./log.txt",
	})

	doWork(loggerInstance)
}
