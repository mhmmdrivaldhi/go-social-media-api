package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/model/dto"
	"github.com/mhmmdrivaldhi/go-social-media-api/service"
	"github.com/mhmmdrivaldhi/go-social-media-api/utils"
)

type authController struct{
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *authController {
	return &authController{authService: service}
}

func (ac *authController) Register(ctx *gin.Context) {
	var register *dto.RegisterRequest

	err := ctx.ShouldBind(&register)
	if err != nil {
		utils.HandleError(ctx, &utils.BadRequestError{Message: err.Error()})
		return
	}

	err = ac.authService.Register(register)
	if err != nil {
		utils.HandleError(ctx, &utils.InternalServerError{Message: err.Error()})
		return
	}

	response := utils.Response(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
		Data:       nil,
	})

	ctx.JSON(http.StatusCreated, response)
}

func (ac *authController) Login(ctx *gin.Context) {
	var login *dto.LoginRequest

	err := ctx.ShouldBind(&login)
	if err != nil {
		utils.HandleError(ctx, &utils.BadRequestError{Message: err.Error()})
		return
	}

	
}