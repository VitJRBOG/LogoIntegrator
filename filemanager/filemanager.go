package filemanager

import (
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

// CreateDirectories создает каталоги images и output
func CreateDirectories() {
	os.Mkdir("images", 0700)
	os.Mkdir("output", 0700)
}

// GetImageFileNames возвращает список имен файлов с png-изображениями из каталога images
func GetImageFileNames(path string) []string {
	fileNames := findFileNames(path)
	imageFileNames := selectImageFileNames(fileNames)
	return imageFileNames
}

func findFileNames(path string) []string {
	var fileNames []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}

func selectImageFileNames(fileNames []string) []string {
	var imageFileNames []string
	for _, fileName := range fileNames {
		posFileExtension := strings.LastIndex(fileName, ".") + 1
		fileExtension := fileName[posFileExtension:]
		if fileExtension == "png" {
			imageFileNames = append(imageFileNames, fileName)
		}
	}
	return imageFileNames
}

// OpenImage возвращает файл с png-изображением
func OpenImage(path string) image.Image {
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

// SaveImage сохраняет новый файл с png-изображением
func SaveImage(newImage *image.RGBA, fileName string) {
	newImageFile, err := os.Create("output/" + fileName)
	if err != nil {
		panic(err.Error())
	}
	defer newImageFile.Close()

	png.Encode(newImageFile, newImage)
}
