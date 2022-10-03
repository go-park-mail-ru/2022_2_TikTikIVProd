package imagesRepository

import (
	imagesUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/usecase"
	"gorm.io/gorm"
)

type DataBaseImages struct {
	db *gorm.DB
}

func NewDataBaseImages(db *gorm.DB) *DataBaseImages {
	return &DataBaseImages{
		db: db,
	}
}

func (dbImages *DataBaseImages) SelectImagesInPost(postID int) (*[]imagesUsecase.Image, error) {
	return &[]imagesUsecase.Image{}, nil
}
