package imagesRepository

import (
	imagesModel "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/model"
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

func (dbImages *DataBaseImages) SelectImagesInPost(postID int) (*[]imagesModel.Image, error) {
	var images []imagesModel.Image
	return &images, nil
}
