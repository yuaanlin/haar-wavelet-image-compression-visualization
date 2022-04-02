package utils

import (
	. "github.com/ken20001207/image-compressor/model"
)

func Copy2DArray(dst *[][]RGB, src *[][]RGB) {
	for i := range *src {
		(*dst)[i] = make([]RGB, len((*src)[i]))
		copy((*dst)[i], (*src)[i])
	}
}
