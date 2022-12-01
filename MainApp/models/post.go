package models

import "time"

type Post struct {
	ID            uint64    `json:"id"`
	UserID        uint64    `json:"user_id"`
	CountLikes    uint64    `json:"count_likes"`
	IsLiked       bool      `json:"is_liked"`
	CommunityID   uint64    `json:"community_id,omitempty"`
	AvatarID      uint64    `json:"avatar_id" readonly:"true"`
	UserFirstName string    `json:"user_first_name" readonly:"true"`
	UserLastName  string    `json:"user_last_name" readonly:"true"`
	Message       string    `json:"message" validate:"required"`
	CreateDate    time.Time `json:"create_date" readonly:"true"`
	Images        []Image   `json:"images"`
}
