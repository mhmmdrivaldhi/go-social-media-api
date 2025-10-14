package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/config"
	"github.com/mhmmdrivaldhi/go-social-media-api/controller"
	"github.com/mhmmdrivaldhi/go-social-media-api/repository"
	"github.com/mhmmdrivaldhi/go-social-media-api/service"
)

func AuthRouter(rg *gin.RouterGroup) {
	authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	// Public routes
	// rg := gin.Default()
	// routes := rg.Group("/api/v1")
	rg.POST("/register", authController.Register)
	rg.POST("/login", authController.Login)
}