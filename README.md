# cli_image_tool
A simple cli application for resizing images

The cli tool runs the go imaging package under the hood (github.com/disintegration/imaging).

To use the tool compile the program or use go run . and add flags for desired functionality.

The tool works but is in the middle of development. New features and bug fixes are in development.

Flags:
  -in
    input path(e.g. -in ./myImage.jpg)
    
  -out
    outupt path, name and desired format (e.g. -out ./some_folder/resizedImage.png)

  -width
    desired width of image (e.g. -width 580)

  -height
      desired height of image (e.g. -height 400)

  -rel
    resize image relative to the original (e.g. -rel 0.5)
