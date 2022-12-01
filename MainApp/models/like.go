package models

type LikePost struct {
	UserID uint64 `json:"user_id" gorm:"column:user_id"`
	PostID uint64 `json:"post_id" gorm:"column:user_post_id"`
}

func (LikePost) TableName() string {
	return "like_post"
}