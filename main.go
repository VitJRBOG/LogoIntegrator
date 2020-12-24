package main

import (
	"github.com/VitJRBOG/test_go_logo_integrator/filemanager"
	"github.com/VitJRBOG/test_go_logo_integrator/imgs"
)

func main() {
	filemanager.CreateDirectories()
	fileNames := filemanager.GetImageFileNames("images/")
	for _, fileName := range fileNames {
		origImage := filemanager.OpenImage("images/" + fileName)
		logoImage := filemanager.OpenImage("logo/logo.png")
		imageWithLogo := imgs.AddLogoToImage(origImage, logoImage)
		filemanager.SaveImage(imageWithLogo, fileName)
	}
}
