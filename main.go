package main

import (
	myFlags "image_resize/flags"
	"log"
	"github.com/disintegration/imaging"
)

func main() {

    myFlagVal := myFlags.Init()
    log.Printf(`
    Input path: %v
    Output path: %v
    Width: %v
    Height: %v
    Float: %v
    Format: %v`,
    myFlagVal.InImgPath, myFlagVal.OutImgPath, myFlagVal.ResizeW, myFlagVal.ResizeH,
    myFlagVal.ResizeRel, myFlagVal.Format)

    inputImage, err := imaging.Open(myFlagVal.InImgPath)
    if err != nil {
        log.Fatalf("Failed to open %v", err)
    }

    if myFlagVal.ResizeW > 0 || myFlagVal.ResizeH > 0 {
        inputImage = imaging.Resize(inputImage, myFlagVal.ResizeW, myFlagVal.ResizeH,
        imaging.Lanczos)
    }

    err = imaging.Save(inputImage, myFlagVal.OutImgPath)
    if err != nil {
        log.Fatalf("Failed to save %v", err)
    }
}
