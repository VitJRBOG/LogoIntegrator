package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	createDirectories()
	origImage := openImage("images/img.png")
	logoImage := openImage("logo/logo.png")
	addLogoToImage(origImage, logoImage)
}

func createDirectories() {
	os.Mkdir("images", 0700)
	os.Mkdir("output", 0700)
}

func openImage(path string) image.Image {
	fileImage, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer fileImage.Close()
	image, err := png.Decode(fileImage)
	if err != nil {
		panic(err.Error())
	}

	return image
}

func addLogoToImage(origImage, logoImage image.Image) {

	newImage := image.NewRGBA(origImage.Bounds())

	draw.Draw(newImage, origImage.Bounds(), origImage, image.Point{0, 0}, draw.Src)

	var x, y int
	for {
		draw.Draw(newImage, origImage.Bounds(), logoImage, image.Point{x, y}, draw.Over)
		if -(y) >= newImage.Bounds().Max.Y {
			break
		} else {
			y -= 100
		}
	}

	saveImage(newImage)
}

func saveImage(newImage *image.RGBA) {
	newImageFile, err := os.Create("output/result.png")
	if err != nil {
		panic(err.Error())
	}

	png.Encode(newImageFile, newImage)
}
