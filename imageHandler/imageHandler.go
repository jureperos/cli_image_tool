package imagehandler

import (
	"bufio"
	"image"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

// TODO:Handle resizing to different formats than input format
func openImg(inImgPath string) *image.Image {
	input, err := imaging.Open(inImgPath)
	if err != nil {
		log.Fatalf("Failed to open %v", err)
	}

	return &input
}

func isValidFormat(filename string) bool {
	_, err := imaging.FormatFromFilename(filename)
	if err != nil {
		return false
	}

	return true
}

func getImageOutPath(inImg string, outDir string) string {
	return outDir + inImg
}

func HandleAllImg(inDir string, ImgWidth int, ImgHeight int, outDir string) {
	files, _ := os.ReadDir(inDir)

	for _, file := range files {
		filename := file.Name()

		filename = strings.TrimPrefix(filename, "- ")

		if isValidFormat(filename) {
			outPath := getImageOutPath(filename, outDir)
			go HandleResize(ImgWidth, ImgHeight, filename, outPath)
		}
	}
}

func HandleResize(ImgWidth int, ImgHeight int, inputPath string, imgOutPath string) {
	inImage := openImg(inputPath)
	log.Println("opica")
	resizedImg := imaging.Resize(*inImage, ImgWidth, ImgHeight, imaging.Lanczos)

	err := imaging.Save(resizedImg, imgOutPath)
	if err != nil {
		log.Fatal("Error saving resized image", err)
	}
}

func HandleResizeRel(outputPath string, inputPath string, relResize float64) {
	imageFile, err := os.Open(inputPath)
	if err != nil {
		log.Println("Error opening image file", err)
	}
	defer imageFile.Close()

	config, _, err := image.DecodeConfig(imageFile)
	if err != nil {
		log.Println("Error decoding image configuration:", err)
		return
	}

	relSizeF := float64(config.Width) * relResize
	relSizeI := int(relSizeF)

	HandleResize(relSizeI, 0, inputPath, outputPath)

	log.Printf("width: %v, \n height: %v", config.Width, config.Height)
}

func HandleFormat(inPath string, outPath string) {
	imgFile, err := os.Open(inPath)
	defer imgFile.Close()

	if err != nil {
		log.Panic("Error opening image file", err)
	}

	dImg, err := imaging.Decode(imgFile)
	if err != nil {
		log.Panic("Error decoding image", err)
	}

	file, err := os.Create(outPath)
	if err != nil {
		log.Panic("Could not create writer", err)
	}

	w := bufio.NewWriter(file)

	format, err := imaging.FormatFromFilename(outPath)
	if err != nil {
		log.Fatal("Error: could not reat format from output string", err)
	}

	err = imaging.Encode(w, dImg, format)
	if err != nil {
		log.Fatal("Error could not encode image", err)
	}
}
