package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/model"
)

type RepositoryI interface {
	SelectUserByNickName(name string) (*model.User, error)
	SelectUserByEmail(email string) (*model.User, error)
	CreateUser(u model.User) (*model.User, error)
	CreateCookie(c model.Cookie) (*model.Cookie, error)
	SelectCookie(value string) (*model.Cookie, error)
}

