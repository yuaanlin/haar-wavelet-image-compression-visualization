package haar

import (
	. "github.com/ken20001207/image-compressor/model"
	"github.com/ken20001207/image-compressor/utils"
)

func ReverseHaarHorizontal(data [][]RGB, level int) {
	width, height := len(data[0])/level, len(data)/level
	array2 := make([][]RGB, len(data))
	utils.Copy2DArray(&array2, &data)
	for i := 0; i < height; i++ {
		for j := 0; j < width/2; j++ {
			array2[i][2*j] = RGB{
				R: data[i][j].R - data[i][width/2+j].R + 128,
				G: data[i][j].G - data[i][width/2+j].G + 128,
				B: data[i][j].B - data[i][width/2+j].B + 128,
			}
			array2[i][2*j+1] = RGB{
				R: data[i][j].R + data[i][width/2+j].R - 128,
				G: data[i][j].G + data[i][width/2+j].G - 128,
				B: data[i][j].B + data[i][width/2+j].B - 128,
			}
		}
	}
	utils.Copy2DArray(&data, &array2)
}

func ReverseHaarVertical(data [][]RGB, level int) {
	width, height := len(data[0])/level, len(data)/level
	array2 := make([][]RGB, len(data))
	utils.Copy2DArray(&array2, &data)
	for i := 0; i < width; i++ {
		for j := 0; j < height/2; j++ {
			array2[2*j][i] = RGB{
				R: data[j][i].R - data[height/2+j][i].R + 128,
				G: data[j][i].G - data[height/2+j][i].G + 128,
				B: data[j][i].B - data[height/2+j][i].B + 128,
			}
			array2[2*j+1][i] = RGB{
				R: data[j][i].R + data[height/2+j][i].R - 128,
				G: data[j][i].G + data[height/2+j][i].G - 128,
				B: data[j][i].B + data[height/2+j][i].B - 128,
			}
		}
	}
	utils.Copy2DArray(&data, &array2)
}
