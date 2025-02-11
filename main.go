package main

import (
	myFlags "image_resize/flags"
)

func main() {
	// TODO: Implement flag for resizing all images in a directory

	// initialize flags
	myFlags.Init()

	// Validate
	myFlags.Validate()

	// Handle reformatting flow
	myFlags.HandleArgs()
}
