package microservice

import (
	"context"

	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	image "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/image"
	"github.com/pkg/errors"
)

type microService struct {
	client image.ImagesClient
}

func New(client image.ImagesClient) imageRep.RepositoryI {
	return &microService{
		client: client,
	}
}

func (imageMS *microService) GetPostImages(postID uint64) ([]*models.Image, error) {
	ctx := context.Background()

	pbGetPostImagesRequest := image.GetPostImagesRequest{
		PostId: postID,
	}

	pbImages, err := imageMS.client.GetPostImages(ctx, &pbGetPostImagesRequest)
	if err != nil {
		return nil, errors.Wrap(err, "image microservice error")
	}

	images := make([]*models.Image, 0)

	for idx := range pbImages.Images {
		img := &models.Image {
			ID: pbImages.Images[idx].Id,
			ImgLink: pbImages.Images[idx].ImgLink,
		}
		images = append(images, img)
	}

	return images, nil
}

func (imageMS *microService) GetImage(imageID uint64) (*models.Image, error) {
	ctx := context.Background()

	pbGetImageRequest := image.ImageId {
		ImageId: imageID,
	}

	pbImage, err := imageMS.client.GetImage(ctx, &pbGetImageRequest)
	if err != nil {
		return nil, errors.Wrap(err, "image microservice error")
	}

	img := &models.Image {
		ID: pbImage.Id,
		ImgLink: pbImage.ImgLink,
	}

	return img, nil
}

func (imageMS *microService) CreateImage(img *models.Image) (error) {
	ctx := context.Background()

	pbImage := image.Image {
		ImgLink: img.ImgLink,
	}

	imgId, err := imageMS.client.CreateImage(ctx, &pbImage)
	if err != nil {
		return errors.Wrap(err, "image microservice error")
	}

	img.ID = imgId.ImageId

	return nil
}

