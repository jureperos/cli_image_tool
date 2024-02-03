package myFlags

import (
	"bufio"
	"image"
	"log"
	"os"
	"github.com/disintegration/imaging"
)

var inImage image.Image
var err error
func HandleArgs() {

    // Call imaging module to resize
    inImage, err = imaging.Open(MyFlagVal.InImgPath)
    if err != nil {
        log.Fatalf("Failed to open %v", err)
}


    if MyFlagVal.ResizeH > 0 || MyFlagVal.ResizeW > 0 {
        if MyFlagVal.ResizeRel != 0  {
            log.Fatal(`Error: Cannot have output width or height (-w, -h)
            specified together with relative output size (-f flag).
            Choose one or the other.`)
        } else {
            HandleResize(MyFlagVal.ResizeW, MyFlagVal.ResizeH)
            return
        }
    }

    if MyFlagVal.ResizeRel != 0  {
        HandleResizeRel()
        return
    }

    // No flags other than format means just reformat
    HandleFormat()
}

func HandleResize(ImgWidth int, ImgHeight int) {

    resizedImg := imaging.Resize(inImage, ImgWidth, ImgHeight, imaging.Lanczos)

    err := imaging.Save(resizedImg, MyFlagVal.OutImgPath)
    if err != nil {
        log.Fatal("Error saving resized image", err)
    }

}

func HandleResizeRel() {
    imageFile, err := os.Open(MyFlagVal.InImgPath)
    if err != nil {
        log.Println("Error opening image file", err)
    }
    defer imageFile.Close()

    config, _, err := image.DecodeConfig(imageFile)
    if err != nil {
        log.Println("Error decoding image configuration:", err)
        return
    }

    relSizeF := float64(config.Width) * MyFlagVal.ResizeRel
    relSizeI := int(relSizeF)

    HandleResize(relSizeI, 0)

    log.Printf("width: %v, \n height: %v", config.Width, config.Height)
}

func HandleFormat() {
    imgFile, err := os.Open(MyFlagVal.InImgPath)
    defer imgFile.Close()

    if err != nil {
        log.Panic("Error opening image file", err)
    }

    dImg, err := imaging.Decode(imgFile)
    if err != nil {
        log.Panic("Error decoding image", err)
    }

    file, err := os.Create(MyFlagVal.OutImgPath)
    if err != nil {
        log.Panic("Could not create writer", err)
    }

    w := bufio.NewWriter(file)

    format, err := imaging.FormatFromFilename(MyFlagVal.OutImgPath)
    if err != nil {
        log.Fatal("Error: could not reat format from output string", err)
    }

    err = imaging.Encode(w, dImg, format)
    if err != nil {
        log.Fatal("Error could not encode image", err)
    }

}
