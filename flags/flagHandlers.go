package myFlags

import (
    "log"
    "os"
    "image"
)

func HandleArgs(allFlagsP *FlagValues) {
    allFlags := *allFlagsP

    if allFlags.ResizeH > 0 || allFlags.ResizeW > 0 {
        if allFlags.ResizeRel != 0  {
            log.Fatal(`Error: Cannot have output width or height
            specified (-w, -h) together with relative output size (-f flag)`)

            HandleResizeRel(allFlags.InImgPath)
        } else {
            HandleResize(allFlags.ResizeW, allFlags.ResizeH)
        }
    }

    if allFlags.ResizeRel != 0  {
        HandleResizeRel(allFlags.InImgPath)
    }
}

func HandleInImgPath(inputImagePath string) {

    if inputImagePath == "" {
        log.Fatal("Error: No input path provided")
    }

    if _, err := os.Stat(inputImagePath); err != nil {
        log.Fatal("Error: Input image path not found")
    }
}
func HandleOutImgPath() {
}

func HandleResize(ImgWidth int, ImgHeight int) {
    // Call imaging module to resize
}

func HandleResizeRel(inputImagePath string) {
    imageFile, err := os.Open(inputImagePath)
    defer imageFile.Close()

    config, _, err := image.DecodeConfig(imageFile)
    if err != nil {
        log.Println("Error decoding image configuration:", err)
        return
    }
    log.Printf("width: %v, \n height: %v", config.Width, config.Height)
}
func HandleFormat() {}
