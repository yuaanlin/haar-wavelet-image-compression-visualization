package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ken20001207/image-compressor/haar"
	"github.com/ken20001207/image-compressor/model"
	"github.com/ken20001207/image-compressor/services"
	"github.com/ken20001207/image-compressor/utils"
	"strconv"
)

func CompressedController(c *gin.Context) {
	// level means the number of haar transform
	level := 2
	if c.Query("level") != "" {
		l, _ := c.GetQuery("level")
		ls, err := strconv.Atoi(l)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "level must be a number",
			})
			return
		}
		level = ls
	}

	ratio := 0
	if c.Query("ratio") != "" {
		s, _ := c.GetQuery("ratio")
		ss, err := strconv.Atoi(s)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "ratio must be a number",
			})
			return
		}
		ratio = ss
	}

	image, err := services.GetImage(c.Request.Context(), c.Query("uid"))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Image not found",
		})
		c.Abort()
		return
	}

	array, err := utils.ConvertImageTo2dArray(*image)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	// execute haar transform to the specified step

	s := 0
	for s <= level*2+1 {

		// transform
		if s != 0 && s < 2*level+1 {
			horizontal := s%2 == 1
			if horizontal {
				haar.Horizontal(array, (s-1)/2+1)
			} else {
				haar.Vertical(array, (s-1)/2+1)
			}
		}

		// compress
		if s == 2*level+1 {
			haar.Compress(array, float64(ratio)/100, level)
		}

		s++
	}

	utils.FixColorRange(array)

	compressed := model.Compressed{}
	compressed.FromRGB(array, level)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "inline; filename=\""+c.Query("uid")+".compressed\"")
	_, err = c.Writer.Write(compressed.Bytes())
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
}
