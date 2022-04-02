package haar

import (
	. "github.com/ken20001207/image-compressor/model"
	"github.com/ken20001207/image-compressor/utils"
)

func Horizontal(data [][]RGB, level int) {
	width, height := len(data[0])/level, len(data)/level
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
	width, height := len(data[0])/level, len(data)/level
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

func Compress(array [][]RGB, percent float64) {
	limit := 128 * percent
	for i, row := range array {
		for j, pixel := range row {
			if pixel.R-128 > -1*int(limit) && pixel.R-128 < int(limit) {
				array[i][j] = RGB{
					R: 128,
					G: 128,
					B: 128,
				}
			}
		}
	}
}
