package microservice_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/repository/microservice"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	image "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/image"
	imageMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/image/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseCreateImage struct {
	ArgData *models.Image
	Error error
}

type TestCaseGetPostImages struct {
	ArgData uint64
	ExpectedRes []*models.Image
	Error error
}

type TestCaseGetImage struct {
	ArgData uint64
	ExpectedRes *models.Image
	Error error
}

func TestMicroserviceCreateImage(t *testing.T) {
	mockPbImage := image.Image {
		ImgLink: "link1",
	}

	img := models.Image {
		ImgLink: mockPbImage.ImgLink,
	}

	pbImageId := image.ImageId {
		ImageId: 1,
	}

	mockPbImageError := image.Image {
		ImgLink: "link2",
	}

	imgError := models.Image {
		ImgLink: mockPbImageError.ImgLink,
	}

	mockImageClient := imageMocks.NewImagesClient(t)

	ctx := context.Background()

	createErr := errors.New("error")

	mockImageClient.On("CreateImage", ctx, &mockPbImage).Return(&pbImageId, nil)
	mockImageClient.On("CreateImage", ctx, &mockPbImageError).Return(nil, createErr)

	repository := imageRep.New(mockImageClient)

	cases := map[string]TestCaseCreateImage {
		"success": {
			ArgData:   &img,
			Error: nil,
		},
		"error": {
			ArgData:   &imgError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.CreateImage(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockImageClient.AssertExpectations(t)
}

func TestMicroserviceGetPostImages(t *testing.T) {
	pbPostId := image.GetPostImagesRequest {
		PostId: 1,
	}

	var mockPbImages image.GetPostImagesResponse
	err := faker.FakeData(&mockPbImages)
	assert.NoError(t, err)

	images := make([]*models.Image, 0)

	for idx := range mockPbImages.Images {
		img := models.Image {
			ID: mockPbImages.Images[idx].Id,
			ImgLink: mockPbImages.Images[idx].ImgLink,
		}

		images = append(images, &img)
	}

	pbPostIdError := image.GetPostImagesRequest {
		PostId: 2,
	}
	
	mockImageClient := imageMocks.NewImagesClient(t)

	ctx := context.Background()

	getErr := errors.New("error")

	mockImageClient.On("GetPostImages", ctx, &pbPostId).Return(&mockPbImages, nil)
	mockImageClient.On("GetPostImages", ctx, &pbPostIdError).Return(nil, getErr)

	repository := imageRep.New(mockImageClient)

	cases := map[string]TestCaseGetPostImages {
		"success": {
			ArgData:   pbPostId.PostId,
			ExpectedRes: images,
			Error: nil,
		},
		"error": {
			ArgData:   pbPostIdError.PostId,
			ExpectedRes: nil,
			Error: getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedImages, err := repository.GetPostImages(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedImages)
			}
		})
	}
	mockImageClient.AssertExpectations(t)
}

func TestMicroserviceGetImage(t *testing.T) {
	pbImageId := image.ImageId {
		ImageId: 1,
	}

	mockPbImage := image.Image {
		Id: pbImageId.ImageId,
		ImgLink: "link1",
	}

	img := &models.Image {
		ID: mockPbImage.Id,
		ImgLink: mockPbImage.ImgLink,
	}

	pbImageIdError := image.ImageId {
		ImageId: 2,
	}
	
	mockImageClient := imageMocks.NewImagesClient(t)

	ctx := context.Background()

	getErr := errors.New("error")

	mockImageClient.On("GetImage", ctx, &pbImageId).Return(&mockPbImage, nil)
	mockImageClient.On("GetImage", ctx, &pbImageIdError).Return(nil, getErr)

	repository := imageRep.New(mockImageClient)

	cases := map[string]TestCaseGetImage {
		"success": {
			ArgData:   pbImageId.ImageId,
			ExpectedRes: img,
			Error: nil,
		},
		"error": {
			ArgData:   pbImageIdError.ImageId,
			ExpectedRes: nil,
			Error: getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedImage, err := repository.GetImage(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedImage)
			}
		})
	}
	mockImageClient.AssertExpectations(t)
}

