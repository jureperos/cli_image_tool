package myFlags

import (
	"bufio"
	"fmt"
	"image_resize/imageHandler"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (fV FlagValues) HandleArgs() {
	if fV.ResizeH > 0 || fV.ResizeW > 0 {
		fV.resize()
		return
	} else if fV.ResizeH < 0 || fV.ResizeW < 0 {
		log.Fatal("Error: Sizes cant' be negative")
	}

	if fV.ResizeRel != 0 {
		fV.resizeRel()
		return
	}

	// No flags other than format means just reformat
	imagehandler.HandleFormat(fV.InImgPath, fV.OutImgPath)
}

func (fV FlagValues) resize() {
	if fV.ResizeRel != 0 {
		log.Fatal("Error: Cannot have output width or height (-w, -h) specified together with relative output size (-f flag).\nChoose one or the other.")
	} else {
		if fV.HandleAll == true {
			fV.handleAll(fV.DirPath, fV.OutDirPath)

			imagehandler.HandleAllImg(fV.DirPath, fV.ResizeW, fV.ResizeH, fV.OutDirPath, fV.ResizeRel)
			return
		}
		imagehandler.HandleResize(fV.ResizeW, fV.ResizeH, fV.InImgPath, fV.OutImgPath)
		return
	}
}

func (fV FlagValues) handleAll(inPath string, outPath string) {
	arePathsSame, _ := arePathsSame(inPath, outPath)
	if arePathsSame {
		if yesNoPrompt("Do you want to overwrite all images?") {
			imagehandler.HandleAllImg(fV.DirPath, fV.ResizeW, fV.ResizeH, fV.OutDirPath, fV.ResizeRel)
			return
		} else {
			log.Println("Change output path to not overwrite it")
			return
		}
	}
}

func (fV FlagValues) resizeRel() {
	if fV.HandleAll == true {
		fV.handleAll(fV.InImgPath, fV.OutImgPath)

		imagehandler.HandleAllImg(fV.DirPath, fV.ResizeW, fV.ResizeH, fV.OutDirPath, fV.ResizeRel)
		return
	}

	imagehandler.HandleResizeRel(fV.OutImgPath, fV.InImgPath, fV.ResizeRel)
	return
}

func arePathsSame(path1, path2 string) (bool, error) {
	// Resolve symlinks and clean paths
	abs1, err := filepath.EvalSymlinks(path1)
	if err != nil {
		return false, fmt.Errorf("error resolving symlink for %q: %v", path1, err)
	}

	abs2, err := filepath.EvalSymlinks(path2)
	if err != nil {
		return false, fmt.Errorf("error resolving symlink for %q: %v", path2, err)
	}

	// Clean paths to normalize them
	clean1 := filepath.Clean(abs1)
	clean2 := filepath.Clean(abs2)

	return clean1 == clean2, nil
}

func yesNoPrompt(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Printf("%s [y/n]: ", question)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v\n", err)
			return false
		}
		response = strings.TrimSpace(response)

		switch strings.ToLower(response) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			log.Println("Please enter 'y' or 'n'")
		}
	}
}
