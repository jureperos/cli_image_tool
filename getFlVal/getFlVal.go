package fl_val

import (
	"flag"
	"fmt"
	cli_flags "image_resize/flags"
)

func GetValues() cli_flags.FlagValues {
	var MyFlagVal cli_flags.FlagValues

	inputImagePath := flag.String("in", "", "Path to the input file or input directory if using -all flag")
	outputImagePath := flag.String("out", "", `Path to the output file,
    with the *file name and extentions included (example: ./exampleImage.jpg or output directory if using -all flag`)

	resizeWidth := flag.Int("width", 0, "Pixel resize width")
	resizeHeight := flag.Int("height", 0, "Pixel resize height")

	relativeResize := flag.Float64("rel", 0, `Float that resizes the new image
    relative to the original image`)

	handleAll := flag.Bool("all", false, "Handle all images in a directory")

	flag.Parse()

	MyFlagVal = cli_flags.FlagValues{
		InImgPath:  *inputImagePath,
		OutImgPath: *outputImagePath,
		ResizeW:    *resizeWidth,
		ResizeH:    *resizeHeight,
		ResizeRel:  *relativeResize,
		HandleAll:  *handleAll,
	}

	fmt.Printf(`
    Input path: %v
    Output path: %v
    Width: %v
    Height: %v
    Relative resize: %v
	Handle all: %v`,
		MyFlagVal.InImgPath, MyFlagVal.OutImgPath, MyFlagVal.ResizeW,
		MyFlagVal.ResizeH, MyFlagVal.ResizeRel, MyFlagVal.HandleAll)

	return MyFlagVal
}
