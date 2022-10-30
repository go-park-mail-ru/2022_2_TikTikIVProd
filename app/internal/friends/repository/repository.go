package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type RepositoryI interface {
	AddFriend(friends models.Friends) error
	DeleteFriend(friends models.Friends) error
	CheckFriends(friends models.Friends) (bool, error)
}
