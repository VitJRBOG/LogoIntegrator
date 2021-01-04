package imgs

import (
	"image"
	"image/draw"
)

// AddLogoToImage добавляет изображение-логотип на изображение-оригинал
// и возвращает указатель на новое изображение
func AddLogoToImage(origImage, logoImage image.Image) *image.RGBA {
	newImage := makeNewImage(origImage)

	y := logoImage.Bounds().Max.Y
	for {
		x := logoImage.Bounds().Max.X
		for {
			draw.Draw(newImage, origImage.Bounds(), logoImage, image.Point{x, y}, draw.Over)

			x -= logoImage.Bounds().Max.X
			logoOutOfImage := checkNewOverlayPosition(x, newImage.Bounds().Max.X)
			if logoOutOfImage {
				break
			}
		}
		y -= logoImage.Bounds().Max.Y
		logoOutOfImage := checkNewOverlayPosition(y, newImage.Bounds().Max.Y)
		if logoOutOfImage {
			break
		}
	}
	return newImage
}

func makeNewImage(origImage image.Image) *image.RGBA {
	newImage := image.NewRGBA(origImage.Bounds())
	draw.Draw(newImage, origImage.Bounds(), origImage, image.Point{0, 0}, draw.Src)
	return newImage
}

func checkNewOverlayPosition(posLogo, posBorder int) bool {
	if -(posLogo) >= posBorder {
		return true
	}
	return false
}
