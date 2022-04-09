package haar

import (
	. "github.com/ken20001207/image-compressor/model"
	"github.com/ken20001207/image-compressor/utils"
	"math"
)

func Horizontal(data [][]RGB, level int) {
	width, height := len(data[0])/int(math.Pow(2, float64(level-1))), len(data)/int(math.Pow(2, float64(level-1)))
	array2 := make([][]RGB, len(data))
	utils.Copy2DArray(&array2, &data)
	for i := 0; i < height; i++ {
		for j := 0; j < width/2; j++ {
			array2[i][j] = RGB{
				R: (data[i][2*j].R + data[i][2*j+1].R) / 2,
				G: (data[i][2*j].G + data[i][2*j+1].G) / 2,
				B: (data[i][2*j].B + data[i][2*j+1].B) / 2,
			}
		}
		for j := 0; j < width/2; j++ {
			array2[i][width/2+j] = RGB{
				R: (data[i][2*j+1].R-data[i][2*j].R)/2 + 128,
				G: (data[i][2*j+1].G-data[i][2*j].G)/2 + 128,
				B: (data[i][2*j+1].B-data[i][2*j].B)/2 + 128,
			}
		}
	}
	utils.Copy2DArray(&data, &array2)
}

func Vertical(data [][]RGB, level int) {
	width, height := len(data[0])/int(math.Pow(2, float64(level-1))), len(data)/int(math.Pow(2, float64(level-1)))
	array2 := make([][]RGB, len(data))
	utils.Copy2DArray(&array2, &data)
	for i := 0; i < width; i++ {
		for j := 0; j < height/2; j++ {
			array2[j][i] = RGB{
				R: (data[2*j][i].R + data[2*j+1][i].R) / 2,
				G: (data[2*j][i].G + data[2*j+1][i].G) / 2,
				B: (data[2*j][i].B + data[2*j+1][i].B) / 2,
			}
		}
		for j := 0; j < height/2; j++ {
			array2[height/2+j][i] = RGB{
				R: (data[2*j+1][i].R-data[2*j][i].R)/2 + 128,
				G: (data[2*j+1][i].G-data[2*j][i].G)/2 + 128,
				B: (data[2*j+1][i].B-data[2*j][i].B)/2 + 128,
			}
		}
	}
	utils.Copy2DArray(&data, &array2)
}

func Compress(array [][]RGB, percent float64, level int) {
	width, height := len(array[0])/int(math.Pow(2, float64(level))), len(array)/int(math.Pow(2, float64(level)))
	limit := 128 * percent
	for i, row := range array {
		for j, pixel := range row {
			if i < height && j < width {
				continue
			}
			if pixel.R-128 > -1*int(limit) && pixel.R-128 < int(limit) {
				array[i][j].R = 128
			}
			if pixel.G-128 > -1*int(limit) && pixel.G-128 < int(limit) {
				array[i][j].G = 128
			}
			if pixel.B-128 > -1*int(limit) && pixel.B-128 < int(limit) {
				array[i][j].B = 128
			}
		}
	}
}
