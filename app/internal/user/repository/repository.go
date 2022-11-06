package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type RepositoryI interface {
	SelectUserByNickName(name string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	CreateUser(u *models.User) error
	SelectUserById(id int) (*models.User, error)
	UpdateUser(user models.User) error
}
