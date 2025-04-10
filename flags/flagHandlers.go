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

func HandleArgs() {
	// TODO:Refactor
	if MyFlagVal.ResizeH > 0 || MyFlagVal.ResizeW > 0 {
		if MyFlagVal.ResizeRel != 0 {
			log.Fatal("Error: Cannot have output width or height (-w, -h) specified together with relative output size (-f flag).\nChoose one or the other.")
		} else {
			if MyFlagVal.HandleAll == true {
				arePathsSame, _ := arePathsSame(MyFlagVal.DirPath, MyFlagVal.OutDirPath)
				if arePathsSame {
					if yesNoPrompt("Do you want to overwrite all images?") {
						imagehandler.HandleAllImg(MyFlagVal.DirPath, MyFlagVal.ResizeW, MyFlagVal.ResizeH, MyFlagVal.OutDirPath, MyFlagVal.ResizeRel)
						return
					} else {
						log.Println("Change output path to not overwrite it")
						return
					}
				}

				imagehandler.HandleAllImg(MyFlagVal.DirPath, MyFlagVal.ResizeW, MyFlagVal.ResizeH, MyFlagVal.OutDirPath, MyFlagVal.ResizeRel)
				return
			}
			imagehandler.HandleResize(MyFlagVal.ResizeW, MyFlagVal.ResizeH, MyFlagVal.InImgPath, MyFlagVal.OutImgPath)
			return
		}
	} else if MyFlagVal.ResizeH < 0 || MyFlagVal.ResizeW < 0 {
		log.Fatal("Error: Sizes cant' be negative")
	}

	if MyFlagVal.ResizeRel != 0 {
		if MyFlagVal.HandleAll == true {
			arePathsSame, _ := arePathsSame(MyFlagVal.InImgPath, MyFlagVal.OutImgPath)
			if arePathsSame {
				if yesNoPrompt("Do you want to overwrite all images?") {
					imagehandler.HandleAllImg(MyFlagVal.DirPath, MyFlagVal.ResizeW, MyFlagVal.ResizeH, MyFlagVal.OutDirPath, MyFlagVal.ResizeRel)
					return
				} else {
					log.Println("Change output path to not overwrite it")
					return
				}
			}
			imagehandler.HandleAllImg(MyFlagVal.DirPath, MyFlagVal.ResizeW, MyFlagVal.ResizeH, MyFlagVal.OutDirPath, MyFlagVal.ResizeRel)
			return
		}

		imagehandler.HandleResizeRel(MyFlagVal.OutImgPath, MyFlagVal.InImgPath, MyFlagVal.ResizeRel)
		return
	}

	// No flags other than format means just reformat
	imagehandler.HandleFormat(MyFlagVal.InImgPath, MyFlagVal.OutImgPath)
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
