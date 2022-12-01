package imageUsecase

import (
	imageRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
)

type ImageUseCaseI interface {
	GetPostImages(postID uint64) ([]*models.Image, error)
	GetImageById(imageID uint64) (*models.Image, error)
	CreateImage(img *models.Image) error
}

type imageUsecase struct {
	imageRep imageRepository.RepositoryI
}

func NewImageUsecase(ir imageRepository.RepositoryI) ImageUseCaseI {
	return &imageUsecase{
		imageRep: ir,
	}
}

func (i *imageUsecase) GetPostImages(postID uint64) ([]*models.Image, error) {
	images, err := i.imageRep.GetPostImages(postID)

	if err != nil {
		return nil, err
	}

	return images, nil
}

func (i *imageUsecase) GetImageById(imageID uint64) (*models.Image, error) {
	image, err := i.imageRep.GetImage(imageID)

	if err != nil {
		return nil, errors.Wrap(err, "GetImage usecase error")
	}

	return image, nil
}

func (i *imageUsecase) CreateImage(img *models.Image) error {
	err := i.imageRep.CreateImage(img)

	if err != nil {
		return errors.Wrap(err, "imageUsecase.CreateImage error")
	}

	return nil
}
