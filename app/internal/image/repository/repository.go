package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"

type RepositoryI interface {
	GetPostImages(postID int) ([]*models.Image, error)
	GetImage(imageID int) (*models.Image, error)
}
