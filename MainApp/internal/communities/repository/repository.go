package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"

type RepositoryI interface {
	GetCommunity(id uint64) (*models.Community, error)
	UpdateCommunity(comm *models.Community) error
	CreateCommunity(comm *models.Community) error
	SearchCommunities(searchString string) ([]*models.Community, error)
	DeleteCommunity(id uint64) error
	JoinCommunity(id uint64, userId uint64) error
	LeaveCommunity(id uint64, userId uint64) error
	GetAllCommunities() ([]*models.Community, error)
	GetAllUserCommunities(userID uint64) ([]*models.Community, error)
	GetCountUserCommunity(id uint64) (uint64, error)
	CheckSubscriptionCommunity(id uint64, userID uint64) (bool, error)
}
