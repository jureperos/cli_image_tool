package cli_flags

import (
	"fmt"
)

// TODO: finish validating all cases
func (fV FlagValues) Validate() error {
	if fV.InImgPath == "./" && !fV.HandleAll {
		return fmt.Errorf("No input image provided!")
	}

	if fV.ResizeH < 0 || fV.ResizeW < 0 {
		return fmt.Errorf("Error: Sizes cant' be negative")
	}

	if fV.ResizeH > 0 || fV.ResizeW > 0 && fV.ResizeRel != 0 {
		return fmt.Errorf("Error: Cannot have output width or height (-w, -h) specified together with relative output size (-f flag).\nChoose one or the other.")
	}

	return nil
}
