package imagehandler

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
)

// TODO: Handle resizing to different formats than input format
// TODO: Implement resource managment for processing big number of files

func HandleAllImg(inDir string, ImgWidth int, ImgHeight int, outDir string, relSize float64) error {
	files, err := os.ReadDir(inDir)
	if err != nil {
		return fmt.Errorf("Could not get current directory info %v", err)
	}

	var wg sync.WaitGroup

	type imgFile struct {
		imgName string
		err     error
	}
	errCh := make(chan imgFile)

	for _, file := range files {
		filename := file.Name()
		filename = strings.TrimPrefix(filename, "- ")

		if isValidFormat(filename) {
			outPath := filepath.Join(outDir, filename)
			filename = filepath.Join(inDir, filename)

			wg.Add(1)

			go func(filename string, outPath string) {
				defer wg.Done()

				var imgFile imgFile
				imgFile.imgName = filename

				if relSize != 0 {
					imgFile.err = ResizeRel(outPath, filename, relSize)
				} else {
					imgFile.err = Resize(ImgWidth, ImgHeight, filename, outPath)
				}

				errCh <- imgFile
			}(filename, outPath)
		} else {
			valFormats := `"jpg" (or "jpeg"), "png", "gif", "tif" (or "tiff") and "bmp" are supported.`
			fmt.Printf("\nInvalid file format: %s\n Valid formats: %s\n\n", filename, valFormats)
		}
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for image := range errCh {
		if image.err != nil {
			fmt.Printf("Error handling %s: %s\n", image.imgName, image.err)
		}

		fmt.Printf("Image \"%s\" complete\n", image.imgName)
	}

	return nil
}

func Resize(ImgWidth int, ImgHeight int, inputPath string, imgOutPath string) error {
	inImage, err := openImg(inputPath)
	if err != nil {
		return err
	}
	// TODO: Add more resample filters
	resizedImg := imaging.Resize(*inImage, ImgWidth, ImgHeight, imaging.Lanczos)

	err = imaging.Save(resizedImg, imgOutPath)
	if err != nil {
		return fmt.Errorf("Save image %s: %w", imgOutPath, err)
	}

	return nil
}

func ResizeRel(outputPath string, inputPath string, relResize float64) error {
	imageFile, err := openImg(inputPath)
	if err != nil {
		return err
	}

	bounds := (*imageFile).Bounds()
	width := bounds.Dx()

	relSizeF := float64(width) * relResize
	relSizeI := int(relSizeF)

	err = Resize(relSizeI, 0, inputPath, outputPath)
	if err != nil {
		return err
	}

	return nil
}

func HandleFormat(inPath string, outPath string) error {
	imgFile, err := os.Open(inPath)
	defer imgFile.Close()

	if err != nil {
		return fmt.Errorf("Error opening image file: %v", err)
	}

	dImg, err := imaging.Decode(imgFile)
	if err != nil {
		return fmt.Errorf("Error decoding image: %v; inPath: %v", err, inPath)
	}

	file, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("Could not create writer: %v", err)
	}

	w := bufio.NewWriter(file)

	format, err := imaging.FormatFromFilename(outPath)
	if err != nil {
		return fmt.Errorf("Could not reat format from output string: %v", err)
	}

	err = imaging.Encode(w, dImg, format)
	if err != nil {
		return fmt.Errorf("Could not encode image: %v", err)
	}

	return nil
}

func openImg(inImgPath string) (*image.Image, error) {
	input, err := imaging.Open(inImgPath)
	if err != nil {
		return nil, fmt.Errorf("Open image %s: %w", inImgPath, err)
	}

	return &input, nil
}

func isValidFormat(filename string) bool {
	_, err := imaging.FormatFromFilename(filename)
	if err != nil {
		return false
	}

	return true
}
