package repository

import "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"

type RepositoryI interface {
	GetAllStickers() ([]*models.Sticker, error)
	GetStickerByID(id uint64) (*models.Sticker, error)
}
