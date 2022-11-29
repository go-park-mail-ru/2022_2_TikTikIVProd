package delivery

import (
	"context"
	imageUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/usecase"
	image "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/proto"
)

type ImageManager struct {
	image.UnimplementedImagesServer
	ImageUC imageUsecase.UseCaseI
}

func New(uc imageUsecase.UseCaseI) image.ImagesServer {
	return ImageManager{ImageUC: uc}
}

func (im ImageManager) GetPostImages(ctx context.Context, pbImages *image.GetPostImagesRequest) (*image.GetPostImagesResponse, error) {
	resp, err := im.ImageUC.GetPostImages(pbImages)
	return resp, err
}

func (im ImageManager) GetImage(ctx context.Context, pbId *image.ImageId) (*image.Image, error) {
	resp, err := im.ImageUC.GetImage(pbId)
	return resp, err
}

func (im ImageManager) CreateImage(ctx context.Context, pbImage *image.Image) (*image.ImageId, error) {
	_, err := im.ImageUC.CreateImage(pbImage)
	return &image.ImageId{ImageId: pbImage.Id}, err
}
