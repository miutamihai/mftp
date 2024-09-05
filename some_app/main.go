package main

import (
	"context"
	"mihaimiuta/mftp/pkg/mftp"
	"mihaimiuta/mftp/pkg/mftp/driver"
	"strconv"
	"time"
)

func doWork(loggerInstance mftp.Logger) {
	if !loggerInstance.IsInitialized() {
		err := loggerInstance.Initialize(context.Background(), &driver.StandardOutputDriver{}, "./config.json")

		if err != nil {
			panic(err)
		}
	}

	for number := range 10_000 {
		err := loggerInstance.Info("doing some work", nil)

		if err != nil {
			panic(err)
		}

		if number%3 == 0 {
			err := loggerInstance.Error("something went wrong", map[string]string{
				"error": "number was multiple of 3",
			})

			if err != nil {
				panic(err)
			}
		} else {
			err := loggerInstance.Debug("something went right", nil)

			if err != nil {
				panic(err)
			}

			err2 := loggerInstance.Debug(strconv.Itoa(number), nil)

			if err2 != nil {
				panic(err2)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

func main() {
	loggerInstance := mftp.Logger{}

	go func() {
		err := loggerInstance.Initialize(context.Background(), &driver.StandardOutputDriver{}, "./config.json")

		if err != nil {
			panic(err)
		}

		doWork(loggerInstance)
	}()

	err := loggerInstance.Initialize(context.Background(), &driver.TextFileDriver{
		FilePath: "./log.txt",
	}, "./config.json")

	if err != nil {
		panic(err)
	}

	doWork(loggerInstance)
}
