package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"

type RepositoryI interface {
	GetPostById(id uint64) (*models.Post, error)
	GetUserPosts(userId uint64) ([]*models.Post, error)
	GetCommunityPosts(userId uint64) ([]*models.Post, error)
	UpdatePost(post *models.Post) error
	CreatePost(u *models.Post) error
	GetAllPosts() ([]*models.Post, error)
	DeletePostById(id uint64) error
	LikePost(id uint64, userId uint64) error
	UnLikePost(id uint64, userId uint64) error
	GetCountLikesPost(id uint64) (uint64, error)
	CheckLikePost(id uint64, userID uint64) (bool, error)
}
