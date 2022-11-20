package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"

type RepositoryI interface {
	GetCommunity(id int) (*models.Community, error)
	UpdateCommunity(comm *models.Community) error
	CreateCommunity(comm *models.Community) error
	DeleteCommunity(id int) error
}
