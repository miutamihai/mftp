package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
)

func main() {
	var buffer bytes.Buffer

	logger := log.New(&buffer, "mftp: ", log.LstdFlags)

	names := []string{"mihai", "samy"}

	for _, name := range slices.Backward(names) {
		logger.Printf("Name: %s\n", name)
	}

	fmt.Println("Buffer")
	fmt.Printf("%s\n", &buffer)
}
