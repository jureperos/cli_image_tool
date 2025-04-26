# cli_image_tool

A straightforward command-line interface for image resizing.

This tool allows you to easily resize images from your terminal. Simply provide the input image path, desired output path and dimensions. It's currently under active development, with new features and improvements on the horizon.

## Getting Started

To use the tool, you'll need to either compile the Go source code or run it directly using `go run .`. You can then control the tool's behavior using the following flags:

## Flags

| Flag      | Description                                                                                                | Example                                                                    |
| :-------- | :--------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------- |
| `-in`     | Specifies the input path to the image file or directory.                                                  | `-in ./myImage.jpg` or `-in ./myImgDir/` (when used with the `-all` flag) |
| `-out`    | Defines the output path, desired filename, and format for the resized image(s).                            | `-out ./some_folder/resizedImage.png` or `-out ./myImgDir/resizedImgs/` (when used with `-all`) |
| `-width`  | Sets the desired width of the output image in pixels.                                                     | `-width 580`                                                              |
| `-height` | Sets the desired height of the output image in pixels.                                                    | `-height 400`                                                              |
| `-rel`    | Resizes the image proportionally based on a relative factor (e.g., `0.5` for 50% of the original size). | `-rel 0.5`                                                              |
| `-all`    | Processes all image files within the directory specified by the `-in` flag.                               | `-all`                                                                    |

## Resampling

This tool leverages the excellent [Go imaging package](https://github.com/disintegration/imaging) for image manipulation, specifically for its resampling filters. Currently, all resizing operations utilize the high-quality **Lanczos resampling method**.

## Important Notes

* **Formatting with `-all`:** Please be aware that reformatting images is not currently supported when using the `-all` flag. To convert the format of a single image, specify the desired extension in the `-out` flag (e.g., `-out image.png`).
* **Resource Limitations:** The tool does not yet implement resource management. When processing a large number of images with the `-all` flag, this could potentially lead to issues such as excessive memory usage, reaching file descriptor limits, or disk I/O bottlenecks. Exercise caution if processing very large directories.

## Future Enhancements

The following features are planned for future releases:

* Implementation of multiple resampling algorithms for greater flexibility.
* Development of custom resampling methods and removal of external dependencies.
* Adjustable resource rate limiting to improve stability when processing directories with the `-all` flag.
* Support for reformatting images when using the `-all` flag.
