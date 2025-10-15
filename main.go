package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/config"
	"github.com/mhmmdrivaldhi/go-social-media-api/router"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	config.InitDB()

	route := gin.Default()
	api := route.Group("/api/v1")


	// api.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "PING!",
	// 	})
	// })
	router.AuthRouter(api)
	router.PostRouter(api)


	route.Run(fmt.Sprintf(":%s", cfg.APIPort))
}