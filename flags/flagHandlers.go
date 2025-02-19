package myFlags

import (
	"image_resize/imageHandler"
	"log"
)

func HandleArgs() {
	if MyFlagVal.ResizeH > 0 || MyFlagVal.ResizeW > 0 {
		if MyFlagVal.ResizeRel != 0 {
			log.Fatal("Error: Cannot have output width or height (-w, -h) specified together with relative output size (-f flag).\nChoose one or the other.")
		} else {
			if MyFlagVal.HandleAll == true {
				// TODO: if in dir same as outdir prompt user if he wants to overwrite all images
				// default behaviour should be to add some postfix (eg. _v1)
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
			// TODO: if in dir same as outdir prompt user if he wants to overwrite all images
			// default behaviour should be to add some postfix (eg. _v1)
			imagehandler.HandleAllImg(MyFlagVal.DirPath, MyFlagVal.ResizeW, MyFlagVal.ResizeH, MyFlagVal.OutDirPath, MyFlagVal.ResizeRel)
			return
		}

		imagehandler.HandleResizeRel(MyFlagVal.OutImgPath, MyFlagVal.InImgPath, MyFlagVal.ResizeRel)
		return
	}

	// No flags other than format means just reformat
	imagehandler.HandleFormat(MyFlagVal.InImgPath, MyFlagVal.OutImgPath)
}
