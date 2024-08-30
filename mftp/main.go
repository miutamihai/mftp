package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"mihai", "samy"}

	for _, name := range slices.Backward(names) {
		fmt.Printf("Name: %s\n", name)
	}
}
