package main

import (
	"mihaimiuta/mftp/driver"
	"time"
)

func doWork(driver driver.Driver) {
	if !driver.IsInitialized() {
		driver.Initialize()
	}

	for _number := range 10_000 {
		driver.Log("something")
		driver.Log(_number)

		time.Sleep(2 * time.Second)
	}
}

func main() {
	driver := driver.StandardOutputDriver{}
	driver.Initialize()

	doWork(&driver)
}
