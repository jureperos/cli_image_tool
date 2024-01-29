package main

import (
	"flag"
	"log"
	"os"
	"github.com/disintegration/imaging"
)

func main() {
    inputImagePath := flag.String("input", "./", "Path to the input file")
    outputImagePath := flag.String("output", "./", `Path to the output file,
    with the file name and extentions included (example: ./exampleImage.jpg`)
    resizeWidth := flag.Int("width", 0, "Optional width to resize width")
    resizeHeight := flag.Int("height", 0, "Optional width to resize height")
    format := flag.String("format", "jpg", "Optional image format (jpeg, png)")

    flag.Parse()

    if *inputImagePath == "" {
        log.Fatal("Error: No input path provided")
    }

    if _, err := os.Stat(*inputImagePath); err != nil {
        log.Fatal("Error: Input image path not found")
    }

    log.Printf(`Input path: %v
    Output path: %v
    Width: %v
    Height: %v
    Format: %v`,
    *inputImagePath, *outputImagePath, *resizeWidth, *resizeHeight, *format)

    inputImage, err := imaging.Open(*inputImagePath)
    if err != nil {
        log.Fatalf("Failed to open %v", err)
    }

    if *resizeWidth > 0 || *resizeHeight > 0 {
        inputImage = imaging.Resize(inputImage, *resizeWidth, *resizeHeight,
        imaging.Lanczos)
    }

    err = imaging.Save(inputImage, *outputImagePath)
    if err != nil {
        log.Fatalf("Failed to save %v", err)
    }
}
