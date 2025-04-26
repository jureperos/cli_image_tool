package cli_flags

import (
	"fmt"
	"image_resize/imageHandler"
	"image_resize/utils"
)

func (fV FlagValues) HandleArgs() error {
	if fV.ResizeH > 0 || fV.ResizeW > 0 {
		return fV.resizeWH()
	}

	if fV.ResizeRel != 0 {
		return fV.resizeRel()
	}

	// No flags other than format means just reformat
	return imagehandler.Format(fV.InImgPath, fV.OutImgPath)
}

func (fV FlagValues) resizeWH() error {
	if fV.HandleAll == true {
		return fV.handleAll(fV.InImgPath, fV.OutImgPath)
	}

	return imagehandler.Resize(fV.ResizeW, fV.ResizeH, fV.InImgPath, fV.OutImgPath)
}

func (fV FlagValues) handleAll(inPath string, outPath string) error {
	arePathsSame, _ := utils.ArePathsSame(inPath, outPath)
	if arePathsSame {
		if utils.YesNoPrompt("Do you want to overwrite all images?") {
			return imagehandler.HandleAllImg(fV.InImgPath, fV.ResizeW, fV.ResizeH, fV.OutImgPath, fV.ResizeRel)
		} else {
			fmt.Println("Change output path to not overwrite it")
			return nil
		}
	}

	return imagehandler.HandleAllImg(fV.OutImgPath, fV.ResizeW, fV.ResizeH, fV.OutImgPath, fV.ResizeRel)
}

func (fV FlagValues) resizeRel() error {
	if fV.HandleAll == true {
		return fV.handleAll(fV.InImgPath, fV.OutImgPath)
	}

	return imagehandler.ResizeRel(fV.OutImgPath, fV.InImgPath, fV.ResizeRel)
}
