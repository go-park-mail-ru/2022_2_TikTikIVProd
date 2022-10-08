package imageUsecase

import (
	imageRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository/postgres"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type ImageUseCaseI interface {
	GetPostImages(postID int) ([]*models.Image, error)
}

type imageUsecase struct {
	imageRep imageRepository.RepositoryI
}

func NewImageUsecase(ir imageRepository.RepositoryI) ImageUseCaseI {
	return &imageUsecase{
		imageRep: ir,
	}
}

func (i *imageUsecase) GetPostImages(postID int) ([]*models.Image, error) {
	images, err := i.imageRep.GetPostImages(postID) //TODO ошибки

	if err != nil {
		return nil, err
	}

	return images, nil
}
