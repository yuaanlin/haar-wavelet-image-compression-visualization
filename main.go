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
	r.POST("/upload", controllers.UploadImage)
	r.GET("/download", controllers.DownloadImage)

	panic(r.Run())
}
