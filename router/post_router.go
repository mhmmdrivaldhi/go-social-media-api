package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/config"
	"github.com/mhmmdrivaldhi/go-social-media-api/controller"
	"github.com/mhmmdrivaldhi/go-social-media-api/middleware"
	"github.com/mhmmdrivaldhi/go-social-media-api/repository"
	"github.com/mhmmdrivaldhi/go-social-media-api/service"
	"github.com/mhmmdrivaldhi/go-social-media-api/utils"
)

func PostRouter(rg *gin.RouterGroup) {
	cfg, _ := config.NewConfig()
	jwtService := utils.NewJwtToken(cfg.JWTConfig)


	postRepo := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	rg.Use(middleware.JWTMiddleware(jwtService))

	rg.POST("/tweets", postController.Create)


}