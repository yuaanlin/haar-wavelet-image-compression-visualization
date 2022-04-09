package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ken20001207/image-compressor/haar"
	"github.com/ken20001207/image-compressor/model"
	"github.com/ken20001207/image-compressor/utils"
	"golang.org/x/image/bmp"
)

func DecompressController(c *gin.Context) {
	file, _, err := c.Request.FormFile("compressed")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No file",
		})
		c.Abort()
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(file); err != nil {
		c.JSON(400, gin.H{
			"message": "Error reading file",
		})
		c.Abort()
		return
	}
	by := buf.Bytes()

	var compressed model.Compressed
	compressed.FromBytes(by)
	array := compressed.RGB()

	level := int(compressed.Level)
	for s := 0; s < level*2; s++ {
		horizontal := s%2 == 1
		if horizontal {
			haar.ReverseHaarHorizontal(array, (level*2+1-s)/2)
		} else {
			haar.ReverseHaarVertical(array, (level*2-s)/2)
		}
	}

	utils.FixColorRange(array)

	// encode result into jpeg for display

	result := utils.Convert2DArrayToImage(array)
	buf = new(bytes.Buffer)
	err = bmp.Encode(buf, result)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.Header("Content-Type", "image/bmp")
	c.Header("Content-Disposition", "inline; filename=\"decompress.bmp\"")

	_, err = c.Writer.Write(buf.Bytes())
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
}
