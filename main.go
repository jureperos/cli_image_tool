package main

import (
	"fmt"
	myFlags "image_resize/flags"
	"time"
)

func main() {
	start := time.Now()

	// initialize flags
	myFlags.Init()

	// Validate
	myFlags.Validate()

	// Handle reformatting flow
	myFlags.HandleArgs()

	end := time.Now()

	deltaTime := end.Sub(start)

	fmt.Printf("Elapsed time async: %s\n", deltaTime)
}
