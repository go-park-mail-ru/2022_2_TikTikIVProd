package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"

type RepositoryI interface {
	GetPostById(id int) (*models.Post, error)
	CreatePost(u models.Post) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
}
