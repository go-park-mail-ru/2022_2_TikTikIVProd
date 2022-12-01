package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/models"
)

type RepositoryI interface {
	GetPostImages(postID uint64) ([]*models.Image, error)
	GetImage(imageID uint64) (*models.Image, error)
	CreateImage(image *models.Image) error
}