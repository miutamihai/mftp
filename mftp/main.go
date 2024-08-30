package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
)

type LoggerInput interface {
	getBuffer() *bytes.Buffer
	getPrefix() string
}

func makeLogger(input LoggerInput) log.Logger {
	buffer := input.getBuffer()
	prefix := input.getPrefix()

	return *log.New(buffer, prefix, log.LstdFlags)
}

type Input struct {
	buffer *bytes.Buffer
}

func (input *Input) getBuffer() *bytes.Buffer {
	if input.buffer == nil {
		var buffer bytes.Buffer

		input.buffer = &buffer
	}

	return input.buffer
}

func (_ Input) getPrefix() string {
	return ""
}

func makeInput() *Input {
	input := Input{}

	return &input
}

func main() {
	input := makeInput()
	logger := makeLogger(input)

	names := []string{"mihai", "samy"}

	for _, name := range slices.Backward(names) {
		logger.Printf("Name: %s\n", name)
	}

	fmt.Println("Buffer")
	fmt.Printf("%s\n", input.buffer)
}
