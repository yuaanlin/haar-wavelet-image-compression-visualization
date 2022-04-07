package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ken20001207/image-compressor/controllers"
	"github.com/ken20001207/image-compressor/redis"
)

func main() {

	redis.NewRedisClient()
	r := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	r.GET("/", controllers.HealthCheck)

	// Upload a .bmp image and get a id
	r.POST("/upload", controllers.UploadImage)

	// Get result of specific step of compression process
	r.GET("/visualization", controllers.VisualizationController)

	// Download .compressed file from given id
	r.GET("/compressed", controllers.CompressedController)

	// Decompress .compressed file into .bmp file
	r.POST("/decompress", controllers.DecompressController)

	panic(r.Run())
}
