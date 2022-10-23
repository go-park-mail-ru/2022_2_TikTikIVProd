package postgres

import (
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Image struct {
	ID      int
	ImgLink string
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) imageRep.RepositoryI {
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

func (dbImage *imageRepository) GetPostImages(postID int) ([]*models.Image, error) {
	var images []*Image
	tx := dbImage.db.Table("images").Joins("JOIN user_posts_images upi ON upi.img_id = images.id AND upi.user_post_id = ?", postID).Scan(&images)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toModelImages(images), nil
}

func (dbImage *imageRepository) GetImage(imageID int) (*models.Image, error) {
	var img Image
	tx := dbImage.db.Table("images").First(&img, imageID)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "Get image repository error")
	}

	return toModelImage(&img), nil
}
