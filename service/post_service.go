package service

import (
	"github.com/mhmmdrivaldhi/go-social-media-api/model/dto"
	"github.com/mhmmdrivaldhi/go-social-media-api/model/entity"
	"github.com/mhmmdrivaldhi/go-social-media-api/repository"
	"github.com/mhmmdrivaldhi/go-social-media-api/utils"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *postService {
	return &postService{postRepo: repo}
}

func (ps *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UserID: req.UserID,
		Tweet: req.Tweet,
	}

	if post.PictureUrl != nil {
		post.PictureUrl = &req.PictureUrl.Filename
	}

	err := ps.postRepo.Create(&post)
	if err != nil {
		return &utils.InternalServerError{Message: err.Error()}
	}

	return nil
}