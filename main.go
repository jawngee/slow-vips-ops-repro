package main

import (
	"github.com/davidbyttow/govips/v2/vips"
	"log"
	"os"
	"time"
)

func TrackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func blurImage() {
	defer TrackTime(time.Now(), "blurImage")

	img, err := vips.NewImageFromFile("./data/test.jpg")
	if err != nil {
		panic(err)
	}

	_ = img.GaussianBlur(18)

	jpgParams := vips.NewJpegExportParams()
	jpgParams.Quality = 100

	buffer, _, err := img.ExportJpeg(jpgParams)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data/blur.jpg", buffer, 0644)
	if err != nil {
		panic(err)
	}
}

func pixelateImage() {
	defer TrackTime(time.Now(), "pixelateImage")

	img, err := vips.NewImageFromFile("./data/test.jpg")
	if err != nil {
		panic(err)
	}

	_ = img.Resize(1.0/16.0, vips.KernelLanczos3)
	_ = img.Resize(16.0, vips.KernelNearest)

	jpgParams := vips.NewJpegExportParams()
	jpgParams.Quality = 100

	buffer, _, err := img.ExportJpeg(jpgParams)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data/pixelate.jpg", buffer, 0644)
	if err != nil {
		panic(err)
	}
}

func blurAndPixelateImage() {
	defer TrackTime(time.Now(), "blurAndPixelateImage")

	img, err := vips.NewImageFromFile("./data/test.jpg")
	if err != nil {
		panic(err)
	}

	_ = img.GaussianBlur(18)
	_ = img.Resize(1.0/16.0, vips.KernelLanczos3)
	_ = img.Resize(16.0, vips.KernelNearest)

	jpgParams := vips.NewJpegExportParams()
	jpgParams.Quality = 100

	buffer, _, err := img.ExportJpeg(jpgParams)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data/both.jpg", buffer, 0644)
	if err != nil {
		panic(err)
	}
}

func blurAndPixelateImageSeparately() {
	defer TrackTime(time.Now(), "blurAndPixelateImageSeparately")

	img, err := vips.NewImageFromFile("./data/test.jpg")
	if err != nil {
		panic(err)
	}

	_ = img.GaussianBlur(18)
	renderedBlur, _, err := img.ExportPng(nil)
	if err != nil {
		panic(err)
	}

	renderedBlurImg, err := vips.NewImageFromBuffer(renderedBlur)
	if err != nil {
		panic(err)
	}

	_ = renderedBlurImg.Resize(1.0/16.0, vips.KernelLanczos3)
	_ = renderedBlurImg.Resize(16.0, vips.KernelNearest)

	jpgParams := vips.NewJpegExportParams()
	jpgParams.Quality = 100

	buffer, _, err := renderedBlurImg.ExportJpeg(jpgParams)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data/both-separately.jpg", buffer, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	blurImage()
	pixelateImage()
	blurAndPixelateImageSeparately()
	blurAndPixelateImage()
}
