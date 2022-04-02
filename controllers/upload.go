package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ken20001207/image-compressor/services"
	"golang.org/x/image/bmp"
)

func UploadImage(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
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

	r := bytes.NewReader(buf.Bytes())
	img, err := bmp.Decode(r)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	uid, err := services.SaveImage(c.Request.Context(), &img)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"id": uid})
	c.Abort()
}
