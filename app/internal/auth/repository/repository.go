package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type RepositoryI interface {
	CreateCookie(cookie *models.Cookie) error
	GetCookie(value string) (string, error)
	DeleteCookie(value string) error
}
