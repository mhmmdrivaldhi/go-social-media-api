package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mhmmdrivaldhi/go-social-media-api/model/dto"
	"github.com/mhmmdrivaldhi/go-social-media-api/service"
	"github.com/mhmmdrivaldhi/go-social-media-api/utils"
)

type postController struct {
	postService service.PostService
}

func NewPostController(service service.PostService) *postController {
	return &postController{
		postService: service,
	}
}

func (pc *postController) Create(ctx *gin.Context) {
	var post dto.PostRequest

	err := ctx.ShouldBind(&post)
	if err != nil {
		utils.HandleError(ctx, &utils.BadRequestError{Message: err.Error()})
		return
	}

	if post.PictureUrl != nil {
		err := os.MkdirAll("/public/picture", 0755)
		if err != nil {
			utils.HandleError(ctx, &utils.InternalServerError{Message: err.Error()})
			return
		}

		// Rename Picture
		ext := filepath.Ext(post.PictureUrl.Filename)
		newFileName := uuid.New().String() + ext

		// Save image to directory
		dst := filepath.Join("/public/picture", filepath.Base(newFileName))
		ctx.SaveUploadedFile(post.PictureUrl, dst)

		post.PictureUrl.Filename = fmt.Sprintf("%s/public/picture/%s", ctx.Request.Host, newFileName)
	}

	userID := 1
	post.UserID = userID
	
	err = pc.postService.Create(&post)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	resp := utils.Response(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "Post Tweet created successfully",
		Data:       nil,
	})

	ctx.JSON(http.StatusCreated, resp)
}

