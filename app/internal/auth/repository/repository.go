package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type RepositoryI interface {
	CreateCookie(c models.Cookie) (*models.Cookie, error)
	SelectCookie(value string) (*models.Cookie, error)
	DeleteCookie(value string) error
}
