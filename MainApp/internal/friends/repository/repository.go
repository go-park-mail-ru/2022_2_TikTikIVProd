package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
)

type RepositoryI interface {
	AddFriend(friends models.Friends) error
	DeleteFriend(friends models.Friends) error
	CheckFriends(friends models.Friends) (bool, error)
	SelectFriends(id uint64) ([]models.User, error)
}
