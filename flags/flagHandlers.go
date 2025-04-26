package cli_flags

import (
	"fmt"
	"image_resize/imageHandler"
)

func (fV FlagValues) HandleArgs() error {
	if fV.ResizeH > 0 || fV.ResizeW > 0 {
		return fV.resizeWH()
	}

	if fV.ResizeRel != 0 {
		return fV.resizeRel()
	}

	return fmt.Errorf("Error: No resize flag specified")
}

func (fV FlagValues) resizeWH() error {
	if fV.HandleAll == true {
		return imagehandler.HandleAllImg(fV.InImgPath, fV.ResizeW, fV.ResizeH, fV.OutImgPath, fV.ResizeRel)
	}

	return imagehandler.Resize(fV.ResizeW, fV.ResizeH, fV.InImgPath, fV.OutImgPath)
}

func (fV FlagValues) resizeRel() error {
	if fV.HandleAll == true {
		return imagehandler.HandleAllImg(fV.InImgPath, fV.ResizeW, fV.ResizeH, fV.OutImgPath, fV.ResizeRel)
	}

	return imagehandler.ResizeRel(fV.OutImgPath, fV.InImgPath, fV.ResizeRel)
}
