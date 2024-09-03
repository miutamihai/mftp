package main

import (
	"context"
	"mihaimiuta/mftp/driver"
	"strconv"
	"time"
)

func doWork(driverInstance driver.Driver) {
	if !driverInstance.IsInitialized() {
		driverInstance.Initialize(context.Background())
	}

	for number := range 10_000 {
		driverInstance.Log(driver.Debug, "something")
		driverInstance.Log(driver.Debug, strconv.Itoa(number))

		time.Sleep(2 * time.Second)
	}
}

func main() {
	driver := driver.StandardOutputDriver{}
	driver.Initialize(context.Background())

	doWork(&driver)
}
