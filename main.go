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

	y := logoImage.Bounds().Max.Y
	for {
		x := logoImage.Bounds().Max.X
		for {
			draw.Draw(newImage, origImage.Bounds(), logoImage, image.Point{x, y}, draw.Over)

			x -= logoImage.Bounds().Max.X

			if -(x) >= newImage.Bounds().Max.X {
				break
			}
		}

		y -= logoImage.Bounds().Max.Y

		if -(y) >= newImage.Bounds().Max.Y {
			break
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
