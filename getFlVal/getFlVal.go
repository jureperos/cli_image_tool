package fl_val

import (
	"flag"
	"fmt"
	cli_flags "image_resize/flags"
)

func GetValues() cli_flags.FlagValues {
	var MyFlagVal cli_flags.FlagValues

	inputImagePath := flag.String("in", "", "Path to the input file")
	outputImagePath := flag.String("out", "", `Path to the output file,
    with the *file name and extentions included (example: ./exampleImage.jpg`)

	resizeWidth := flag.Int("width", 0, "Optional width to resize width")
	resizeHeight := flag.Int("height", 0, "Optional width to resize height")

	relativeResize := flag.Float64("f", 0, `Float that resizes the new image
    relative to the original image`)

	format := flag.String("format", "jpg", "Optional image format (jpeg, png)")

	handleAll := flag.Bool("all", false, "Handle all images in a directory")

	flag.Parse()

	MyFlagVal = cli_flags.FlagValues{
		InImgPath:  *inputImagePath,
		OutImgPath: *outputImagePath,
		ResizeW:    *resizeWidth,
		ResizeH:    *resizeHeight,
		ResizeRel:  *relativeResize,
		Format:     *format,
		HandleAll:  *handleAll,
	}

	fmt.Printf(`
    Input path: %v
    Output path: %v
    Width: %v
    Height: %v
    Float: %v
    Format: %v
	Handle all: %v`,
		MyFlagVal.InImgPath, MyFlagVal.OutImgPath, MyFlagVal.ResizeW,
		MyFlagVal.ResizeH, MyFlagVal.ResizeRel, MyFlagVal.Format,
		MyFlagVal.HandleAll)

	return MyFlagVal
}
