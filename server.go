package main

import (
	"io"
	"os"

	"github.com/0x414e4459/gin-poc/controller"
	"github.com/0x414e4459/gin-poc/middleware"
	"github.com/0x414e4459/gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/vedios", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	server.POST("/vedios", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})
	// for prod:
	// gin.SetMode(gin.ReleaseMode)
	server.Run(":5000")
}
