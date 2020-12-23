package main

import (
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	createDirectories()
	filesNames := findFilesImages("images/")
	for _, fileName := range filesNames {
		origImage := openImage("images/" + fileName)
		logoImage := openImage("logo/logo.png")
		newImage := addLogoToImage(origImage, logoImage)
		saveImage(newImage, fileName)
	}
}

func createDirectories() {
	os.Mkdir("images", 0700)
	os.Mkdir("output", 0700)
}

func findFilesImages(path string) []string {
	var filesNames []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		posFileExtension := strings.LastIndex(file.Name(), ".") + 1
		fileExtension := file.Name()[posFileExtension:]
		if fileExtension == "png" {
			filesNames = append(filesNames, file.Name())
		}
	}
	return filesNames
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

func addLogoToImage(origImage, logoImage image.Image) *image.RGBA {

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

	return newImage
}

func saveImage(newImage *image.RGBA, fileName string) {
	newImageFile, err := os.Create("output/" + fileName)
	if err != nil {
		panic(err.Error())
	}
	defer newImageFile.Close()

	png.Encode(newImageFile, newImage)
}
