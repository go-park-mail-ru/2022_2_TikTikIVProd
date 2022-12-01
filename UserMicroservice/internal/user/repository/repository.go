package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/models"
)

type RepositoryI interface {
	SelectUserByNickName(name string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	CreateUser(u *models.User) error
	SelectUserById(id uint64) (*models.User, error)
	UpdateUser(user models.User) error
	SelectAllUsers() ([]models.User, error)
	SearchUsers(name string) ([]models.User, error)
	
	AddFriend(friends models.Friends) error
	DeleteFriend(friends models.Friends) error
	CheckFriends(friends models.Friends) (bool, error)
	SelectFriends(id uint64) ([]models.User, error)
}
