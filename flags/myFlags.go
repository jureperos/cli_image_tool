package myFlags

import (
	"flag"
	"fmt"
	"log"
)

type FlagValues struct {
	InImgPath, OutImgPath string
	ResizeW, ResizeH      int
	ResizeRel             float64
	Format                string
	HandleAll             bool
}

var MyFlagVal FlagValues

func Init() {
	inputImagePath := flag.String("in", "./", "Path to the input file")
	outputImagePath := flag.String("out", "./", `Path to the output file,
    with the *file name and extentions included (example: ./exampleImage.jpg`)

	resizeWidth := flag.Int("width", 0, "Optional width to resize width")
	resizeHeight := flag.Int("height", 0, "Optional width to resize height")

	relativeResize := flag.Float64("f", 0, `Float that resizes the new image
    relative to the original image`)

	format := flag.String("format", "jpg", "Optional image format (jpeg, png)")

	handleAll := flag.Bool("all", false, "Handle all images in a directory")

	flag.Parse()

	MyFlagVal = FlagValues{ // TODO: maybe make a singleton with once.Do()?
		InImgPath:  *inputImagePath,
		OutImgPath: *outputImagePath,
		ResizeW:    *resizeWidth,
		ResizeH:    *resizeHeight,
		ResizeRel:  *relativeResize,
		Format:     *format,
		HandleAll:  *handleAll,
	}

	log.Printf(`
    Input path: %v
    Output path: %v
    Width: %v
    Height: %v
    Float: %v
    Format: %v`,
		MyFlagVal.InImgPath, MyFlagVal.OutImgPath, MyFlagVal.ResizeW,
		MyFlagVal.ResizeH, MyFlagVal.ResizeRel, MyFlagVal.Format)

	if MyFlagVal.InImgPath == "./" {
		fmt.Println(MyFlagVal.InImgPath)
		log.Fatal("No input image provided")
	}

	if MyFlagVal.OutImgPath == "./" {
		log.Fatal("No output name provided")
	}
}
