package usecase

import (
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/models"
	image "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/proto"
	"github.com/pkg/errors"
)

type UseCaseI interface {
	GetPostImages(*image.GetPostImagesRequest) (*image.GetPostImagesResponse, error)
	GetImage(*image.ImageId) (*image.Image, error)
	CreateImage(*image.Image) (*image.Nothing, error)
}

type useCase struct {
	imageRepository imageRep.RepositoryI
}

func New(imageRepository imageRep.RepositoryI) UseCaseI {
	return &useCase{
		imageRepository: imageRepository,
	}
}

func (uc *useCase) GetPostImages(pbPostId *image.GetPostImagesRequest) (*image.GetPostImagesResponse, error) {
	images, err := uc.imageRepository.GetPostImages(pbPostId.PostId)
	if err != nil {
		return nil, errors.Wrap(err, "image repository postgres error")
	}

	pbImages := &image.GetPostImagesResponse{}

	for idx := range images {
		img := &image.Image {
			Id: images[idx].ID,
			ImgLink: images[idx].ImgLink,
		}
		pbImages.Images = append(pbImages.Images, img)
	}

	return pbImages, nil
}

func (uc *useCase) GetImage(pbImageId *image.ImageId) (*image.Image, error) {
	img, err := uc.imageRepository.GetImage(pbImageId.ImageId)
	if err != nil {
		return nil, errors.Wrap(err, "image repository postgres error")
	}
	return &image.Image {
		Id: img.ID,
		ImgLink: img.ImgLink,
	}, nil
}

func (uc *useCase) CreateImage(pbImage *image.Image) (*image.Nothing, error) {
	modelImage := models.Image {
		ID: pbImage.Id,
		ImgLink: pbImage.ImgLink,
	}
	err := uc.imageRepository.CreateImage(&modelImage)
	if err != nil {
		return nil, errors.Wrap(err, "image repository postgres error")
	}

	pbImage.Id = modelImage.ID

	return &image.Nothing{Dummy: true}, nil
}
