package main

import (
	"fmt"
	getFlVal "image_resize/getFlVal"
	"time"
)

func main() {
	// TODO: Remove timer when merging to main
	start := time.Now()

	flVal := getFlVal.GetValues()

	err := flVal.Validate()
	if err != nil {
		fmt.Printf("\nValidation error: %v\n", err)
	}

	err = flVal.HandleArgs()
	if err != nil {
		fmt.Printf("\nError handling images: %v\n", err)
	}

	end := time.Now()

	deltaTime := end.Sub(start)

	fmt.Printf("Elapsed time async: %s\n", deltaTime)
}
