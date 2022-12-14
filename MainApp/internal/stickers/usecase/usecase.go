package imageUsecase

import (
	stickerRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/stickers/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
)

type StickerUseCaseI interface {
	GetAllStickers() ([]*models.Sticker, error)
	GetStickerByID(id uint64) (*models.Sticker, error)
}

type stickerUsecase struct {
	stickerRep stickerRepository.RepositoryI
}

func NewStickerUsecase(sr stickerRepository.RepositoryI) StickerUseCaseI {
	return &stickerUsecase{
		stickerRep: sr,
	}
}

func (su *stickerUsecase) GetStickerByID(id uint64) (*models.Sticker, error) {
	sticker, err := su.stickerRep.GetStickerByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "GetStickerByID usecase error")
	}

	return sticker, nil
}

func (su *stickerUsecase) GetAllStickers() ([]*models.Sticker, error) {
	stickers, err := su.stickerRep.GetAllStickers()
	if err != nil {
		return nil, errors.Wrap(err, "GetAllStickers usecase error")
	}

	return stickers, nil
}

