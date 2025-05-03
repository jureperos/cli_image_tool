package main

import (
	"fmt"
	getFlVal "image_resize/getFlVal"
	"os"
)

func main() {
	flVal := getFlVal.GetValues()

	err := flVal.Validate()
	if err != nil {
		fmt.Printf("\nValidation error: %v\n", err)
		os.Exit(1)
	}

	err = flVal.HandleArgs()
	if err != nil {
		fmt.Printf("\nError handling images: %v\n", err)
		os.Exit(1)
	}
}
