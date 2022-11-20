package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"

type RepositoryI interface {
	GetPostById(id int) (*models.Post, error)
	GetUserPosts(userId int) ([]*models.Post, error)
	GetCommunityPosts(userId int) ([]*models.Post, error)
	UpdatePost(post *models.Post) error
	CreatePost(u *models.Post) error
	GetAllPosts() ([]*models.Post, error)
	DeletePostById(id int) error
}
