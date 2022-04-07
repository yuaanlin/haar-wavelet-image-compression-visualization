package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ken20001207/image-compressor/haar"
	"github.com/ken20001207/image-compressor/services"
	"github.com/ken20001207/image-compressor/utils"
	"image/jpeg"
	"strconv"
)

func VisualizationController(c *gin.Context) {

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

	// step means the number of transform step
	step := 0
	if c.Query("step") != "" {
		s, _ := c.GetQuery("step")
		ss, err := strconv.Atoi(s)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "step must be a number",
			})
			return
		}
		step = ss
	}

	if step < 0 {
		step = 0
	}

	if step > level*4+1 {
		step = level*4 + 1
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
	for s <= step {

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
			haar.Compress(array, float64(ratio)/100)
		}

		// reverse
		if s > 2*level+1 {
			horizontal := s%2 == 1
			if horizontal {
				haar.ReverseHaarHorizontal(array, level-(s-level-1)/2+2)
			} else {
				haar.ReverseHaarVertical(array, level-(s-level)/2+2)
			}
		}

		s++
	}

	utils.FixColorRange(array)

	// encode result into jpeg for display

	result := utils.Convert2DArrayToImage(array)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, result, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.Header("Content-Type", "image/jpeg")

	_, err = c.Writer.Write(buf.Bytes())
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
}
