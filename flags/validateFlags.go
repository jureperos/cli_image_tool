package cli_flags

import (
	"fmt"
	"image_resize/utils"
)

func (fV FlagValues) Validate() error {
	if fV.InImgPath == "./" && !fV.HandleAll {
		return fmt.Errorf("No input image provided!")
	}

	err := valPath(fV.InImgPath, fV.OutImgPath)
	if err != nil {
		return err
	}

	if fV.ResizeH < 0 || fV.ResizeW < 0 {
		return fmt.Errorf("Error: Sizes cant' be negative")
	}

	if fV.ResizeH > 0 || fV.ResizeW > 0 && fV.ResizeRel != 0 {
		return fmt.Errorf("Error: Cannot have output width or height (-w, -h) specified together with relative output size (-f flag).\nChoose one or the other.")
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

func valPath(inPath string, outPath string) error {
	inPathExists, err := utils.Exists(inPath)
	if err != nil {
		return fmt.Errorf("Error checking input path: %v", err)
	}

	if !inPathExists {
		return fmt.Errorf("Input path does not exist")
	}

	outPathExists, err := utils.Exists(outPath)
	if err != nil {
		return fmt.Errorf("Error checking output path: %v", err)
	}

	if !outPathExists {
		return fmt.Errorf("Output path does not exist")
	}

	return nil
}
