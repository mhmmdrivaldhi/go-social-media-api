package entity

import "time"

type Post struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Tweet      string  `json:"tweet"`
	PictureUrl *string `json:"picture_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}