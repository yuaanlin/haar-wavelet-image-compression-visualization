package utils

import "github.com/ken20001207/image-compressor/model"

func FixColorRange(array [][]model.RGB) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j].R < 0 {
				array[i][j].R = 0
			}
			if array[i][j].R > 255 {
				array[i][j].R = 255
			}
			if array[i][j].G < 0 {
				array[i][j].G = 0
			}
			if array[i][j].G > 255 {
				array[i][j].G = 255
			}
			if array[i][j].B < 0 {
				array[i][j].B = 0
			}
			if array[i][j].B > 255 {
				array[i][j].B = 255
			}
		}
	}
}
