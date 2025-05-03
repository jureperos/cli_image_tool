package cli_flags

import (
	"fmt"
	"image_resize/utils"
)

func (fV FlagValues) Validate() error {
	if fV.InImgPath == "./" && !fV.HandleAll {
		return fmt.Errorf("No input image provided!")
	}

	err := pathCheck(fV.InImgPath, fV.OutImgPath, fV.HandleAll)
	if err != nil {
		return err
	}

	if fV.ResizeH < 0 || fV.ResizeW < 0 {
		return fmt.Errorf("Error: Sizes cant' be negative")
	}

	if (fV.ResizeH > 0 || fV.ResizeW > 0) && fV.ResizeRel != 0 {
		return fmt.Errorf("Error: Cannot have output width or height (-w, -h) specified together with relative output size (-rel flag).\nChoose one or the other.")
	}

	if fV.InImgPath == "" || fV.OutImgPath == "" {
		return fmt.Errorf("Error: Specify input and output path")
	}

	arePathsSame, _ := utils.ArePathsSame(fV.InImgPath, fV.OutImgPath)
	if arePathsSame {
		if utils.YesNoPrompt("Having the same input and output paths will overwrite images! Proceed?") {
			return nil
		} else {
			return fmt.Errorf("Change output path to prevent overwrite!")
		}
	}

	return nil
}

func pathCheck(inPath, outPath string, handleAll bool) error {
	err := valPath(inPath)
	if err != nil {
		return fmt.Errorf("Input error\n %v", err)
	}

	if handleAll {
		err = valPath(outPath)
		if err != nil {
			return fmt.Errorf("Output error\n %v", err)
		}
	}

	return nil
}

func valPath(path string) error {
	pathExists, err := utils.Exists(path)
	if err != nil {
		return fmt.Errorf("Error checking: %v", err)
	}

	if !pathExists {
		return fmt.Errorf("Path does not exist")
	}

	return nil
}
