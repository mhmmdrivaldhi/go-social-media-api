package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/config"
	"github.com/mhmmdrivaldhi/go-social-media-api/router"
)

func main() {
	config.LoadConfig()
	config.InitDB()

	route := gin.Default()
	api := route.Group("/api/v1")


	// api.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "PING!",
	// 	})
	// })
	router.AuthRouter(api)

	route.Run(fmt.Sprintf(":%s", config.AppConfig.APIPort))
}