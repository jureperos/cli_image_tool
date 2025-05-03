package cli_flags

type FlagValues struct {
	InImgPath, OutImgPath string
	ResizeW, ResizeH      int
	ResizeRel             float64
	HandleAll             bool
}
