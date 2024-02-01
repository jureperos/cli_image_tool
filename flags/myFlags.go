package myFlags

import (
    "flag"
)

type FlagValues struct {
    InImgPath string
    OutImgPath string
    ResizeW int
    ResizeH int
    ResizeRel float64
    Format string
}

func Init() FlagValues {

    inputImagePath := flag.String("in", "./", "Path to the input file")

    outputImagePath := flag.String("out", "./", `Path to the output file,
    with the file name and extentions included (example: ./exampleImage.jpg`)

    resizeWidth := flag.Int("width", 0, "Optional width to resize width")
    resizeHeight := flag.Int("height", 0, "Optional width to resize height")
    relativeResize := flag.Float64("f", 0, `Float that resizes the new image
    relative to the original image`)

    format := flag.String("format", "jpg", "Optional image format (jpeg, png)")

    flag.Parse()

    MyFlagVal := FlagValues{
       InImgPath: *inputImagePath,
       OutImgPath: *outputImagePath,
       ResizeW: *resizeWidth,
       ResizeH: *resizeHeight,
       ResizeRel: *relativeResize,
       Format: *format,
    }

    allFlagsP := &MyFlagVal

    HandleArgs(allFlagsP)
    HandleInImgPath(*inputImagePath)
    HandleResizeRel(*inputImagePath)


    return MyFlagVal
}
