package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"

type RepositoryI interface {
	GetPostFiles(postID uint64) ([]*models.File, error)
	GetFileById(fileID uint64) (*models.File, error)
	CreateFile(file *models.File) error
}
