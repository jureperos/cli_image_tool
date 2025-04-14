package main

import (
	"fmt"
	getFlVal "image_resize/getFlVal"
	"time"
)

func main() {
	start := time.Now()

	flVal := getFlVal.GetValues()

	flVal.Validate()

	// Handle reformatting flow
	flVal.HandleArgs()

	end := time.Now()

	deltaTime := end.Sub(start)

	fmt.Printf("Elapsed time async: %s\n", deltaTime)
}
