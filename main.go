package main

import (
	myFlags "image_resize/flags"
)

func main() {
    // TODO: Implement relative resizing operation
    // TODO: Implement flag for resizing all images in a directory

    // initialize flags
    myFlags.Init()

    // Handle reformatting flow
    myFlags.HandleArgs()

}
