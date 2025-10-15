package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/config"
	"github.com/mhmmdrivaldhi/go-social-media-api/controller"
	"github.com/mhmmdrivaldhi/go-social-media-api/repository"
	"github.com/mhmmdrivaldhi/go-social-media-api/service"
)

func PostRouter(rg *gin.RouterGroup) {
	postRepo := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	rg.POST("/posts", postController.Create)

}