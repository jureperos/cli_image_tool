package myFlags

import (
	"bufio"
	"image"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func openImage(inImg string) *image.Image {
	input, err := imaging.Open(inImg)
	if err != nil {
		log.Fatalf("Failed to open %v", err)
	}

	return &input
}

// TODO: Separate image handling from arg handling
func HandleArgs() {
	inImage := openImage(MyFlagVal.InImgPath)

	if MyFlagVal.ResizeH > 0 || MyFlagVal.ResizeW > 0 {
		if MyFlagVal.ResizeRel != 0 {
			log.Fatal("Error: Cannot have output width or height (-w, -h) specified together with relative output size (-f flag).\nChoose one or the other.")
		} else {
			handleResize(MyFlagVal.ResizeW, MyFlagVal.ResizeH, inImage, MyFlagVal.OutImgPath)
			return
		}
	} else if MyFlagVal.ResizeH < 0 || MyFlagVal.ResizeW < 0 {
		log.Fatal("Error: Sizes cant' be negative")
	}

	if MyFlagVal.ResizeRel != 0 {
		HandleResizeRel(MyFlagVal.OutImgPath, inImage)
		return
	}

	// No flags other than format means just reformat
	HandleFormat()
}

func handleResize(ImgWidth int, ImgHeight int, imgIn *image.Image, imgOutPath string) {
	resizedImg := imaging.Resize(*imgIn, ImgWidth, ImgHeight, imaging.Lanczos)

	err := imaging.Save(resizedImg, imgOutPath)
	if err != nil {
		log.Fatal("Error saving resized image", err)
	}
}

func HandleResizeRel(outputPath string, inImg *image.Image) {
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

	handleResize(relSizeI, 0, inImg, outputPath)

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
