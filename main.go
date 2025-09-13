package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/config"
)

func main() {
	config.LoadConfig()
	config.InitDB()

	route := gin.Default()
	api := route.Group("/api")


	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PING!",
		})
	})

	route.Run(fmt.Sprintf(":%s", config.AppConfig.APIPort))
}