package repository

import (
	"github.com/mhmmdrivaldhi/go-social-media-api/model/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db: db}
}

func (pr *postRepository) Create(post *entity.Post) error {
	err := pr.db.Create(&post).Error
	if err != nil {
		return err
	}

	return nil

}