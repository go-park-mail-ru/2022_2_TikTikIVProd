package postgres

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"gorm.io/gorm"
	"log"
)

type RepositoryI interface {
	GetPostImages(postID int) ([]*models.Image, error)
}

type Image struct {
	ID      int
	ImgLink string
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) RepositoryI {
	return &imageRepository{
		db: db,
	}
}

func toPostgresImage(im *models.Image) *Image {
	return &Image{
		ID:      im.ID,
		ImgLink: im.ImgLink,
	}
}

func toModelImage(im *Image) *models.Image {
	return &models.Image{
		ID:      im.ID,
		ImgLink: im.ImgLink,
	}
}

func toModelImages(images []*Image) []*models.Image {
	out := make([]*models.Image, len(images))

	for i, b := range images {
		out[i] = toModelImage(b)
	}

	return out
}

func (dbImage *imageRepository) GetPostImages(imageId int) ([]*models.Image, error) {
	var images []*Image
	tx := dbImage.db.Table("images").Select("img_link").Joins("JOIN user_posts_images upi ON upi.img_id = images.id AND upi.user_post_id = ?", imageId).Scan(&images)

	log.Println("Fetch ")
	if tx.Error != nil {
		return nil, tx.Error
	}

	return toModelImages(images), nil
}
