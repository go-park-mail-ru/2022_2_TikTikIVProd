package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	imageMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/repository/mocks"
	imageUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/models"
	image "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetPostImages struct {
	ArgData *image.GetPostImagesRequest
	ExpectedRes *image.GetPostImagesResponse
	Error error
}

type TestCaseGetImage struct {
	ArgData *image.ImageId
	ExpectedRes *image.Image
	Error error
}

type TestCaseCreateImage struct {
	ArgData *image.Image
	Error error
}

func TestUsecaseCreateImage(t *testing.T) {
	var mockPbImageSuccess image.Image
	err := faker.FakeData(&mockPbImageSuccess)
	assert.NoError(t, err)

	modelImageSuccess := models.Image{
		ID: mockPbImageSuccess.Id,
		ImgLink: mockPbImageSuccess.ImgLink,
	}

	var mockPbImageError image.Image
	err = faker.FakeData(&mockPbImageError)
	assert.NoError(t, err)

	modelImageError := models.Image{
		ID: mockPbImageError.Id,
		ImgLink: mockPbImageError.ImgLink,
	}

	createErr := errors.New("error")

	mockImageRepo := imageMocks.NewRepositoryI(t)

	mockImageRepo.On("CreateImage", &modelImageSuccess).Return(nil)
	mockImageRepo.On("CreateImage", &modelImageError).Return(createErr)

	useCase := imageUsecase.New(mockImageRepo)

	cases := map[string]TestCaseCreateImage {
		"success": {
			ArgData:   &mockPbImageSuccess,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbImageError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.CreateImage(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockImageRepo.AssertExpectations(t)
}

func TestUsecaseGetImage(t *testing.T) {
	mockPbImageIdSuccess := image.ImageId {
		ImageId: 1,
	}

	mockPbImageSuccess := image.Image {
		Id: mockPbImageIdSuccess.ImageId,
		ImgLink: "link",
	}

	modelImageSuccess := models.Image{
		ID: mockPbImageSuccess.Id,
		ImgLink: mockPbImageSuccess.ImgLink,
	}

	mockPbImageIdError := image.ImageId {
		ImageId: 2,
	}

	getErr := errors.New("error")

	mockImageRepo := imageMocks.NewRepositoryI(t)

	mockImageRepo.On("GetImage", mockPbImageIdSuccess.ImageId).Return(&modelImageSuccess, nil)
	mockImageRepo.On("GetImage", mockPbImageIdError.ImageId).Return(nil, getErr)

	useCase := imageUsecase.New(mockImageRepo)

	cases := map[string]TestCaseGetImage {
		"success": {
			ArgData:   &mockPbImageIdSuccess,
			ExpectedRes: &mockPbImageSuccess,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbImageIdError,
			Error: getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			img, err := useCase.GetImage(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, img)
			}
		})
	}
	mockImageRepo.AssertExpectations(t)
}

func TestUsecaseGetPostImages(t *testing.T) {
	mockPbPostIdSuccess := image.GetPostImagesRequest {
		PostId: 1,
	}

	modelImages := make([]*models.Image, 0)
	err := faker.FakeData(&modelImages)
	assert.NoError(t, err)

	mockPbImages := &image.GetPostImagesResponse{}

	for idx := range modelImages {
		img := &image.Image {
			Id: modelImages[idx].ID,
			ImgLink: modelImages[idx].ImgLink,
		}
		mockPbImages.Images = append(mockPbImages.Images, img)
	}

	mockPbPostIdError := image.GetPostImagesRequest {
		PostId: 2,
	}

	getErr := errors.New("error")

	mockImageRepo := imageMocks.NewRepositoryI(t)

	mockImageRepo.On("GetPostImages", mockPbPostIdSuccess.PostId).Return(modelImages, nil)
	mockImageRepo.On("GetPostImages", mockPbPostIdError.PostId).Return(nil, getErr)

	useCase := imageUsecase.New(mockImageRepo)

	cases := map[string]TestCaseGetPostImages {
		"success": {
			ArgData:   &mockPbPostIdSuccess,
			ExpectedRes: mockPbImages,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbPostIdError,
			Error: getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			imgs, err := useCase.GetPostImages(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, imgs)
			}
		})
	}
	mockImageRepo.AssertExpectations(t)
}

