package utils

import (
	"github.com/ken20001207/image-compressor/model"
	"image"
	"image/color"
)

func Convert2DArrayToImage(array [][]model.RGB) image.Image {
	width := len(array[0])
	height := len(array)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(array[y][x].R),
				G: uint8(array[y][x].G),
				B: uint8(array[y][x].B),
				A: 255,
			})
		}
	}

	return img
}

func ConvertImageTo2dArray(img image.Image) ([][]model.RGB, error) {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	array := make([][]model.RGB, height)

	for y := 0; y < height; y++ {
		array[y] = make([]model.RGB, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			array[y][x] = model.RGB{
				R: int(r >> 8),
				G: int(g >> 8),
				B: int(b >> 8),
			}
		}
	}

	return array, nil
}
