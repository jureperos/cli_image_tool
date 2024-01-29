# cli_image_tool

A simple CLI application for resizing images.

The CLI tool runs the Go imaging package under the hood ([github.com/disintegration/imaging](https://github.com/disintegration/imaging)).

To use the tool, compile the program or use `go run .` and add flags for the desired functionality. The tool is a work in progress, with new features and bug fixes in development.

## Flags:

- **-in**
  - Input path (e.g., `-in ./myImage.jpg`)

- **-out**
  - Output path, name, and desired format (e.g., `-out ./some_folder/resizedImage.png`)

- **-width**
  - Desired width of the image (e.g., `-width 580`)

- **-height**
  - Desired height of the image (e.g., `-height 400`)

- **-rel**
  - Resize image relative to the original (e.g., `-rel 0.5`)

