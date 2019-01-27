package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	os.Mkdir("images", 0700)
	os.Mkdir("output", 0700)

	file, err := os.Open("images/img.jpg")
	fmt.Print("file, err := " + err.Error())

	//logo, err := os.Open("images/logo.png")
	//fmt.Print(err.Error())

	img, _, err := image.Decode(file)

	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	//

	new_file, err := os.Create("output/img.jpg")
	fmt.Print("new_file, err := " + err.Error())

	new_file.Close()

	png.Encode(new_file, img)

}
